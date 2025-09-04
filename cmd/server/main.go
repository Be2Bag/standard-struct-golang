// @title Standard Struct Golang API
// @version 1.0
// @description API for Standard Struct Golang API
// @BasePath        /

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"standard-struct-golang/app"
	"standard-struct-golang/config"
	module "standard-struct-golang/modules"

	"github.com/alecthomas/kingpin"
)

var Version string

func main() {
	//สร้าง Instance ของ kingpin เพื่อเอาไว้ใช้ read arg ตอน run
	kp := kingpin.New(filepath.Base(os.Args[0]), fmt.Sprintf("%s", Version))
	//นำค่าจาก --tag version มาเก็บไว้ ตัวแปล Version
	kp.Version(Version)
	//กำหนด arg version
	versionCmd := kp.Command("version", "Show application version.")
	//กำหนด arg start
	startCmd := kp.Command("start", "Start application.")

	//กำหนด arg config-file เพื่อให้สามารถใส่ path
	cfgFilePath := kp.Flag("config-file", "Set load config file (default: config.yml)").Default("config.yml").String()
	switch kingpin.MustParse(kp.Parse(os.Args[1:])) {
	case versionCmd.FullCommand():
		//ถ้าตอน run ใช้ tag version จะ print version ออกมา
		fmt.Println(Version)
	case startCmd.FullCommand():
		//ถ้าตอน run ใช้ tag start จะ run
		config := config.LoadConfig(*cfgFilePath, Version) //load config ของแอพ
		application := app.NewApp(config)                  // ประกาศ instance ของ app
		application.InitialFiberSever()                    // initial fiberApp
		appLog := application.NewLogger().WithField("package", "main")
		//ทำการติดตั้ง module ลงไปที่แอพ
		if err := module.CreateModule(application); err != nil {
			appLog.Errorln("[x] Start create module failed -:", err)
			os.Exit(1)
		}
		errOnStart := application.StartHTTP() // start fiber server
		//ดัก error ในตอนที่ server error
		if errOnStart != nil {
			appLog.Fatal("[x] Start http server error :", errOnStart)
			os.Exit(2)
		}
	}

}
