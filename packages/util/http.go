package util

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// get request id จาก context
func GetHttpRequestId(ctx context.Context) string {
	requestId, ok := ctx.Value("requestid").(string)
	if ok {
		return requestId
	}
	return ""
}

type HttpSkipper struct {
	Rule map[string]struct{}
}

func NewHttpSkipper() *HttpSkipper {
	return &HttpSkipper{Rule: map[string]struct{}{}}
}

func (s *HttpSkipper) Add(m string, p string) {
	s.Rule[fmt.Sprintf("%s|%s", m, p)] = struct{}{}
}

func (s *HttpSkipper) Has(m string, p string) bool {
	if _, ok := s.Rule[fmt.Sprintf("%s|%s", m, p)]; ok {
		return ok
	}
	return false
}

func CreateCookie(ctx *fiber.Ctx, name string, value string, times time.Time) {
	ctx.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		Expires:  times,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})
}
func DeleteCookie(ctx *fiber.Ctx, name string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     name,
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})
}
