package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"runtime/debug"
	"standard-struct-golang/packages/util"
	"time"

	"standard-struct-golang/docs"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

func (app *App) InitialFiberSever() {

	app.log.Infof("[*] Initialize fiber server")
	//ตั้งค่า fiber
	fiberCfg := fiber.Config{
		ReadTimeout:           app.Config.ServerConfig.ReadTimeout,
		WriteTimeout:          app.Config.ServerConfig.WriteTimeout,
		IdleTimeout:           app.Config.ServerConfig.IdleTimeout,
		DisableStartupMessage: true,
		Prefork:               false,
		ServerHeader:          app.Config.ServerConfig.ServerHeader,
		ProxyHeader:           app.Config.ServerConfig.ProxyHeader,
		ErrorHandler:          serverErrorHandler,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	}
	//สร้าง instance fiber ใหม่
	fiberApp := fiber.New(fiberCfg)
	//หากใน Env มีการ Enable core จะทำให้ Fiber App ใช้ Core
	if app.Config.ServerConfig.EnableCORS {
		app.log.Infoln("[*] Used fiber cors middleware")
		fiberApp.Use(cors.New())
	}
	app.log.Infoln("[*] Used fiber request id middleware")
	//ทำการกำหนดให้ Fiber ใช้ middleware ที่จะสร้าง id ให้กับ request
	fiberApp.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return ksuid.New().String()
		},
	}))
	app.log.Infoln("[*] Used custom fiber logger middleware")
	//ทำการกำหนดรูปแบบของ Log เมื่อเกิด request เป็นแบบกำหนดเอง
	fiberApp.Use(logger.New(logger.Config{
		Format:     "${locals:requestid} - ${ip} - ${method} ${path} ${status} - ${latency}\n",
		TimeZone:   "Asia/Bangkok",
		TimeFormat: time.ANSIC,
		Next: func(c *fiber.Ctx) bool {
			// ถ้าเป็ย path /api/-/health จะไม่แสดง log
			switch c.Path() {
			case "/api/-/health":
				return true
			case "/swagger":
				return true
			default:
				return false
			}
		},
	}))
	app.log.Infoln("[*] Use recovery on crash middleware")
	//ให้ fibe ใช้ middleware recovery on crash เพื่อให้เมื่อเกิด crash แล้ว sever ไม่ down
	fiberApp.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: fiberStackTraceHandler,
	}))
	app.log.Infoln("[*] Set X-Server headers")
	//นำ ip host มา set X-Server header
	fiberApp.Use(func(ctx *fiber.Ctx) error {
		host, _ := os.Hostname()
		//กำหนด regex ของ IP
		re := regexp.MustCompile(`\d{1,3}(?:\.\d{1,3}){3}`)
		xServer := re.FindString(host)
		if xServer == "" {
			xServer = host
		}
		ctx.Set("X-Server-By", re.FindString(xServer))
		type requestidKey string
		keyname := requestidKey("requestid")
		//เพิ่ม requestid User context
		userCtx := context.WithValue(ctx.UserContext(), keyname, util.GetHttpRequestId(ctx.Context()))
		ctx.SetUserContext(userCtx)
		return ctx.Next()
	})

	app.log.Infoln("[*] Set Swagger documentation")
	app.setupSwagger(fiberApp)

	app.log.Infoln("[*] Set api router")
	app.Router = fiberApp
	//เพิ่ม Endpoint health check
	app.log.Infoln("[*] Initialize healthcheck endpoint")
	app.Router.Get("/api/-/health", func(fiberContext *fiber.Ctx) error {
		return fiberContext.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
	})
}

func fiberStackTraceHandler(_ *fiber.Ctx, e interface{}) {
	logrus.Errorln(fmt.Sprintf("[PANIC] %v\n%s\n", e, debug.Stack()))
}

func serverErrorHandler(fiberContext *fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError
	fiberError, ok := err.(*fiber.Error)
	if ok {
		statusCode = fiberError.Code
	}
	if statusCode >= fiber.StatusInternalServerError {
		logrus.Errorln("[PANIC]", fmt.Sprintf("[%s] %s %s : %s", fiberContext.IP(),
			fiberContext.Route().Method, fiberContext.Route().Path, err))
	}
	fiberContext.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return fiberContext.Status(statusCode).JSON(fiber.Map{"error": err.Error()})
}

func (app *App) StartHTTP() error {

	if app.Config.AppConfig.IsDebug() {
		fr := app.Router.GetRoutes()
		for _, r := range fr {
			app.log.Debugln(r.Name, r.Method, r.Path)
		}
	}

	app.log.Infoln("[*] Starting http server...")
	serverShutdown := make(chan os.Signal, 1)
	signal.Notify(serverShutdown, os.Interrupt)
	go func() {
		<-serverShutdown
		app.log.Infof("[*] Server terminating...")
		errRouterShutdown := app.Router.Shutdown()
		if errRouterShutdown != nil {
			app.log.Errorf("[x] Server shutdown failed: %v\n", errRouterShutdown)
		}
	}()

	listenAt := fmt.Sprintf("%s:%s", app.Config.ServerConfig.ListenIP,
		app.Config.ServerConfig.Port)
	app.log.Infof("[*] Starting server at %s", listenAt)
	errOnListen := app.Router.Listen(listenAt)
	if errOnListen != nil && !errors.Is(errOnListen, http.ErrServerClosed) {
		app.log.Errorf("[x] Start server error: %s", errOnListen.Error())
		return errOnListen
	}
	return nil
}

func (app *App) setupSwagger(fiberApp *fiber.App) {

	docs.SwaggerInfo.Host = "localhost" + ":" + app.Config.ServerConfig.Port
	docs.SwaggerInfo.BasePath = app.Config.AppConfig.PrefixPath
	fiberApp.Get("/swagger/*", swagger.HandlerDefault)
	app.log.Infoln("[*] Swagger UI available at /swagger/")
}
