package health_id

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

// ใช้สำหรับขอ authen-code
func (h *HealthId) GetAuthenCode(ctx context.Context, input AuthenCodeInput) (*BaseHealthIdResponse, error) {

	if h.mophUrl == "" {
		return nil, fmt.Errorf("url is nil")
	}

	uri := h.mophUrl + "/api/v1/health-id/authen-code"
	body, _ := sonic.Marshal(input)

	header := map[string]string{
		fiber.HeaderContentType:   fiber.MIMEApplicationJSON,
		fiber.HeaderAuthorization: "Bearer " + input.Authorization,
	}

	resp, err := h.http.Post(ctx, uri, header, bytes.NewBuffer(body), h.Timeout)
	if err != nil {
		h.log.WithError(err).Error("HTTP Post request failed")
		return nil, fmt.Errorf("failed to make POST request: %w", err)
	}

	if resp.Header["Content-Type"][0] == "application/json" {
		var response BaseHealthIdResponse
		if err := sonic.Unmarshal(resp.Body, &response); err != nil {
			return nil, fmt.Errorf("unmarshal response error: %w", err)
		}
		if response.Status == "success" {
			jsonResData, errOnMarshal := sonic.Marshal(response.Data)
			if errOnMarshal != nil {
				return nil, errOnMarshal
			}
			var responseData AuthenCodeResponseData
			errOnUnmarshal := sonic.Unmarshal(jsonResData, &responseData)
			if errOnUnmarshal != nil {
				return nil, errOnUnmarshal
			}
			response.Data = responseData
			return &response, nil
		}
		return nil, fmt.Errorf("%s", response.Message)
	} else {
		return nil, fmt.Errorf("login health id by code error : invalid request body")
	}
}
