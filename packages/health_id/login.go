package health_id

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *HealthId) GetURLLogin(ctx context.Context) (string, error) {
	return fmt.Sprintf("%s/oauth/redirect?client_id=%s&redirect_uri=%s&response_type=code", h.url, h.id, h.redirectUrl), nil
}

func (h *HealthId) GetURLLocal(ctx context.Context) (string, error) {
	return fmt.Sprintf("%s/oauth/redirect?client_id=%s&redirect_uri=%s&response_type=code", h.url, h.id, h.redirectLocalhost), nil
}

func (h *HealthId) GetHealthIdTokenByCode(ctx context.Context, code string, uri string) (*ResponseHealthIdToken, string, int, string, string, error) {
	// ctx, cancel := context.WithTimeout(ctx, h.Timeout)
	// defer cancel()

	if h.url == "" {
		return nil, "", http.StatusNotFound, "", "", errors.New("url is nil")
	}

	baseURL := h.url
	endpoint := "/api/v1/token"
	uriPath := baseURL + endpoint

	body, _ := sonic.Marshal(fiber.Map{
		"grant_type":    "authorization_code",
		"code":          code,
		"client_id":     h.id,
		"client_secret": h.secret,
		"redirect_uri":  uri,
	})

	header := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
	}

	h.log.WithFields(logrus.Fields{
		"url":          endpoint,
		"redirect_uri": uri,
		"client_id":    h.id,
		"code_length":  len(code),
	}).Debug("Requesting Health ID token")

	resp, err := h.http.Post(ctx, uriPath, header, bytes.NewBuffer(body), h.Timeout)
	reqBodyLog := string(body)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			h.log.WithField("url", endpoint).Error("Request time out after 30 seconds")
			return nil, endpoint, resp.Code, reqBodyLog, string(resp.Body), fmt.Errorf("request time out after 30 seconds")
		}
		h.log.WithFields(logrus.Fields{
			"url":   endpoint,
			"error": err.Error(),
		}).Error("Failed to make POST request")
		return nil, endpoint, resp.Code, reqBodyLog, string(resp.Body), fmt.Errorf("failed to make POST request: %w", err)
	}

	// Log the raw response for debugging
	h.log.WithFields(logrus.Fields{
		"status_code": resp.Code,
		"raw_body":    string(resp.Body),
		"url":         endpoint,
		"headers":     resp.Header,
	}).Debug("Raw Health ID token response")

	var response HealthIdResponse
	if err := sonic.Unmarshal(resp.Body, &response); err != nil {
		// Log the error with the raw response body
		h.log.WithFields(logrus.Fields{
			"error":       err.Error(),
			"status_code": resp.Code,
			"raw_body":    string(resp.Body),
			"url":         endpoint,
		}).Error("Failed to unmarshal Health ID token response")

		// Include part of the raw response in the error for immediate debugging
		return nil, endpoint, resp.Code, reqBodyLog, string(resp.Body), fmt.Errorf("unmarshal response error: %w", err)
	}

	if resp.Code == http.StatusOK && response.Status == "success" {
		response.Data.RedirectUri = uri
		h.log.WithField("url", endpoint).Info("Successfully obtained Health ID token")
		return &response.Data, endpoint, resp.Code, reqBodyLog, string(resp.Body), nil
	}

	h.log.WithFields(logrus.Fields{
		"status_code": resp.Code,
		"message":     response.Message,
		"status":      response.Status,
		"url":         endpoint,
	}).Error("Health ID API returned error")

	return nil, endpoint, resp.Code, reqBodyLog, string(resp.Body), fmt.Errorf("%s (status code %d)", response.Message, resp.Code)
}
