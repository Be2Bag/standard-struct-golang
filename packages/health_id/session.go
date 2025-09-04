package health_id

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

// ใช้สำหรับการขอ access token ที่ใช้สำหรับการใช้งานในระดับ app
func (h *HealthId) GetSession(ctx context.Context) (*BaseHealthIdResponse, error) {

	if h.mophUrl == "" {
		return nil, fmt.Errorf("url is nil")
	}

	uri := h.mophUrl + "/api/v1/sessions"
	body, _ := sonic.Marshal(fiber.Map{
		"client_id":  h.mophClientID,
		"secret_key": h.mophSecret,
	})

	header := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
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
			return &response, nil
		}
		return nil, fmt.Errorf("unsuccessful result: %s", response.Message)
	} else if resp.Header["Content-Type"][0] != "application/json" {
		resp.Code = 400
		return nil, fmt.Errorf("login health id by code error : invalid request body")
	}

	return nil, fmt.Errorf("login health id by code error : unexpected error")
}
