package handler

import (
	"fmt"
	"net/http"
	"standard-struct-golang/appconst"
	authdto "standard-struct-golang/modules/frontweb/modules/auth/dto"
	"time"

	"github.com/gofiber/fiber/v2"

	dto "standard-struct-golang/modules/frontweb/modules/dto"
	validator "standard-struct-golang/packages/util"
)

func (h *AuthHandler) AddAuthRouter(r fiber.Router) {
	versionOne := r.Group("/v1")
	authGroup := versionOne.Group("/auth")

	authGroup.Post("/LoginWithProviderId", h.LoginWithProviderId)

}

// @Summary Login with Health ID
// @Description Login with Health ID
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body authdto.RequestLoginWithHealthId true "Login with Health ID"
// @Success 200 {object} dto.BaseResponse
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /v1/auth/LoginWithProviderId [post]
func (h AuthHandler) LoginWithProviderId(ctx *fiber.Ctx) error {
	var req authdto.RequestLoginWithHealthId
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			MessageEN:  "Request body is invalid: " + err.Error(),
			MessageTH:  fmt.Sprintf("Request body ไม่ถูกต้อง: %s", err.Error()),
			Status:     fiber.ErrBadRequest.Message,
			Data:       nil,
		})
	}
	if err := validator.ValidateStruct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			MessageEN:  "Some field is invalid.",
			MessageTH:  "พบบาง field ผิดพลาด",
			Status:     fiber.ErrBadRequest.Message,
			Data:       err,
		})
	}
	auth, status, err := h.svc.LoginHealthId(ctx.Context(), req.Code, req.RedirectUri)
	if err != nil {
		if status >= 500 {
			return ctx.Status(fiber.StatusServiceUnavailable).JSON(dto.BaseResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				MessageEN:  "Login with Health ID failed: " + err.Error(),
				MessageTH:  "ไม่สามารถเข้าสู่ระบบได้: การเข้าสู่ระบบ Health ID ล้มเหลว",
				Status:     fiber.ErrServiceUnavailable.Message,
				Data:       nil,
			})
		}
		return ctx.Status(status).JSON(dto.BaseResponse{
			StatusCode: status,
			MessageEN:  "Login with Health ID failed: " + err.Error(),
			MessageTH:  "ไม่สามารถเข้าสู่ระบบได้: การเข้าสู่ระบบ Health ID ล้มเหลว",
			Status:     http.StatusText(status),
			Data:       nil,
		})
	}
	provider, status, err := h.svc.GetProviderToken(ctx.Context(), auth.AccessToken)
	if err != nil {
		if status >= 500 {
			return ctx.Status(fiber.StatusServiceUnavailable).JSON(dto.BaseResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				MessageEN:  "Login with Health ID failed: " + err.Error(),
				MessageTH:  "ไม่สามารถเข้าสู่ระบบได้: ไม่สามารถรับ token ผู้ให้บริการ",
				Status:     fiber.ErrServiceUnavailable.Message,
				Data:       nil,
			})
		}
		return ctx.Status(status).JSON(dto.BaseResponse{
			StatusCode: status,
			MessageEN:  "Failed to get provider token: " + err.Error(),
			MessageTH:  "ไม่สามารถเข้าสู่ระบบได้: ไม่สามารถรับ token ผู้ให้บริการ",
			Status:     http.StatusText(status),
			Data:       nil,
		})
	}

	data, status, err := h.svc.GetProviderData(ctx.Context(), provider.AccessToken)
	if err != nil {
		if status >= 500 {
			return ctx.Status(fiber.StatusServiceUnavailable).JSON(dto.BaseResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				MessageEN:  "Login with Health ID failed: " + err.Error(),
				MessageTH:  "ไม่สามารถเข้าสู่ระบบได้: ไม่สามารถรับข้อมูลผู้ให้บริการ",
				Status:     fiber.ErrServiceUnavailable.Message,
				Data:       nil,
			})
		}
		return ctx.Status(status).JSON(dto.BaseResponse{
			StatusCode: status,
			MessageEN:  "Failed to get provider data: " + err.Error(),
			MessageTH:  "ไม่สามารถเข้าสู่ระบบได้: ไม่สามารถรับข้อมูลผู้ให้บริการ",
			Status:     http.StatusText(status),
			Data:       nil,
		})
	}
	sessionToken, errCreateSession := h.svc.CreateSession(ctx.Context(), data, provider.AccessToken)
	if errCreateSession != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.BaseResponse{
			StatusCode: fiber.StatusInternalServerError,
			MessageEN:  "Failed to create session token : " + errCreateSession.Error(),
			MessageTH:  "ไม่สามารถเข้าสู่ระบบได้: ไม่สามารถสร้าง token session ได้",
			Status:     fiber.ErrInternalServerError.Message,
			Data:       nil,
		})
	}

	validator.CreateCookie(ctx, appconst.AuthRegister, sessionToken, time.Now().Add(15*time.Minute))

	return ctx.Status(fiber.StatusOK).JSON(dto.BaseResponse{
		StatusCode: fiber.StatusOK,
		MessageEN:  "Login success",
		MessageTH:  "เข้าสู่ระบบสําเร็จ",
		Status:     "OK",
		Data:       nil,
	})
}
