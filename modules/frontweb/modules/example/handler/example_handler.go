package handler

import (
	"errors"
	"standard-struct-golang/models"
	dto "standard-struct-golang/modules/frontweb/modules/dto"
	exampledto "standard-struct-golang/modules/frontweb/modules/example/dto"
	example_repo "standard-struct-golang/modules/frontweb/modules/example/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *ExampleHandler) AddExampleRouter(r fiber.Router) {
	versionOne := r.Group("/v1")
	exampleGroup := versionOne.Group("/example")

	exampleGroup.Get("/:example_id", h.GetExample)
	exampleGroup.Post("/", h.CreateExample)
	exampleGroup.Post("/send-notify/:cid", h.SendMophLineNotify)
}

// @Summary Create a new example
// @Description คำอธิบาย API Create a new example
// @Tags Example (ตัวอย่าง)
// @Accept json
// @Produce json
// @Param request body exampledto.ExampleRequest true "Example request"
// @Success 200 {object} dto.BaseResponse{} "Example created successfully"
// @Router /v1/example/ [post]
func (h *ExampleHandler) CreateExample(ctx *fiber.Ctx) error {
	var reqBody exampledto.ExampleRequest

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			MessageTH:  "คำขอไม่ถูกต้อง",
			MessageEN:  err.Error(),
			Status:     "error",
			Data:       nil,
		})
	}

	example := models.Example{
		ExampleID: uuid.New().String(),
		Detail:    reqBody.Detail,
	}

	if err := h.svc.CreateExample(ctx.Context(), example); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.BaseResponse{
			StatusCode: fiber.StatusInternalServerError,
			MessageTH:  "ข้อผิดพลาดภายในเซิร์ฟเวอร์",
			MessageEN:  err.Error(),
			Status:     "error",
			Data:       nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.BaseResponse{
		StatusCode: fiber.StatusOK,
		MessageTH:  "สร้างตัวอย่างสำเร็จ",
		MessageEN:  "Example created successfully",
		Status:     "success",
		Data:       nil,
	})
}

// @Summary Get an example by ID
// @Description คำอธิบาย API Get an example by ID
// @Tags Example (ตัวอย่าง)
// @Accept json
// @Produce json
// @Param example_id path string true "Example ID"
// @Success 200 {object} dto.BaseResponse{} "Example retrieved successfully"
// @Router /v1/example/{example_id} [get]
func (h *ExampleHandler) GetExample(ctx *fiber.Ctx) error {

	exampleID := ctx.Params("example_id")

	example, err := h.svc.GetExampleByID(ctx.Context(), exampleID)
	if err != nil {
		if errors.Is(err, example_repo.ErrExampleNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(dto.BaseResponse{
				StatusCode: fiber.StatusNotFound,
				MessageTH:  "ไม่พบข้อมูลตัวอย่าง",
				MessageEN:  "example not found",
				Status:     "error",
				Data:       nil,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.BaseResponse{ // other errors
			StatusCode: fiber.StatusInternalServerError,
			MessageTH:  "ข้อผิดพลาดภายในเซิร์ฟเวอร์",
			MessageEN:  err.Error(),
			Status:     "error",
			Data:       nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.BaseResponse{
		StatusCode: fiber.StatusOK,
		MessageTH:  "ดึงตัวอย่างสำเร็จ",
		MessageEN:  "Get example successfully",
		Status:     "success",
		Data:       example,
	})
}

// @Summary Send Moph Line Notification
// @Description คำอธิบาย API Send Moph Line Notification
// @Tags Example (ตัวอย่าง)
// @Accept json
// @Produce json
// @Param cid path string true "CID"
// @Success 200 {object} dto.BaseResponse{} "Notification sent successfully"
// @Router /v1/example/send-notify/{cid} [post]
func (h *ExampleHandler) SendMophLineNotify(ctx *fiber.Ctx) error {
	cid := ctx.Params("cid")

	if err := h.svc.SendMophLineNotify(ctx.Context(), cid); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.BaseResponse{
			StatusCode: fiber.StatusInternalServerError,
			MessageTH:  "ข้อผิดพลาดภายในเซิร์ฟเวอร์",
			MessageEN:  err.Error(),
			Status:     "error",
			Data:       nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.BaseResponse{
		StatusCode: fiber.StatusOK,
		MessageTH:  "ส่งการแจ้งเตือนสำเร็จ",
		MessageEN:  "Send notification successfully",
		Status:     "success",
		Data:       nil,
	})
}
