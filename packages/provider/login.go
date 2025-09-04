package provider

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (p *Provider) GetProviderToken(ctx context.Context, accessToken string) (*ResponseProviderToken, string, int, string, string, error) {

	baseURL := p.url
	endpoint := "/api/v1/services/token"
	uriPath := baseURL + endpoint

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body, _ := sonic.Marshal(fiber.Map{
		"client_id":  p.id,
		"secret_key": p.secret,
		"token_by":   "Health ID",
		"token":      accessToken,
	})

	resp, err := p.http.Post(ctx, uriPath, headers, bytes.NewBuffer(body), p.Timeout)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, endpoint, resp.Code, string(body), string(resp.Body), fmt.Errorf("request time out after 30 seconds")
		}
		return nil, endpoint, resp.Code, string(body), string(resp.Body), err
	}

	p.log.WithFields(logrus.Fields{
		"status_code": resp.Code,
		"raw_body":    string(resp.Body),
		"url":         endpoint,
	}).Debug("Raw provider token response")

	var response ProviderResponse
	if err := sonic.Unmarshal(resp.Body, &response); err != nil {
		p.log.WithFields(logrus.Fields{
			"error":       err.Error(),
			"status_code": resp.Code,
			"raw_body":    string(resp.Body),
			"url":         endpoint,
		}).Error("Failed to unmarshal provider token response")

		return nil, endpoint, resp.Code, string(body), string(resp.Body), fmt.Errorf("unmarshal response error: %w", err)
	}

	if resp.Code == http.StatusOK && response.Data.Result == "Success" {
		return &response.Data, endpoint, resp.Code, string(body), "", nil
	}

	p.log.WithFields(logrus.Fields{
		"status_code": resp.Code,
		"message":     response.Message,
		"url":         endpoint,
	}).Error("Provider token API returned error")

	return nil, endpoint, resp.Code, string(body), string(resp.Body), fmt.Errorf("get provider token error: %s (status code %d)", response.Message, resp.Code)
}

func (p *Provider) GetProviderData(ctx context.Context, accessToken string) (*ProviderData, string, int, string, string, error) {

	baseURL := p.url
	endpoint := "/api/v1/services/profile"
	uriPath := baseURL + endpoint
	reqBodytoken := "access_token=" + accessToken
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		"client-id":     p.id,
		"secret-key":    p.secret,
		"content-type":  "application/json",
	}
	resp, err := p.http.Get(ctx, uriPath, header, nil, p.Timeout)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, endpoint, resp.Code, reqBodytoken, string(resp.Body), fmt.Errorf("request time out after 30 seconds")
		}
		return nil, endpoint, resp.Code, reqBodytoken, string(resp.Body), err
	}

	p.log.WithFields(logrus.Fields{
		"status_code": resp.Code,
		"raw_body":    string(resp.Body),
		"url":         endpoint,
	}).Debug("Raw provider data response")

	var response ResponseProviderData
	if err := sonic.Unmarshal(resp.Body, &response); err != nil {

		p.log.WithFields(logrus.Fields{
			"error":       err.Error(),
			"status_code": resp.Code,
			"raw_body":    string(resp.Body),
			"url":         endpoint,
		}).Error("Failed to unmarshal provider data response")

		return nil, endpoint, resp.Code, reqBodytoken, string(resp.Body), fmt.Errorf("unmarshal response error: %w", err)
	}

	if resp.Code == http.StatusOK || response.Status == 200 {
		return &response.Data, endpoint, resp.Code, reqBodytoken, string(resp.Body), nil
	}

	p.log.WithFields(logrus.Fields{
		"status_code": resp.Code,
		"message":     response.Message,
		"url":         endpoint,
	}).Error("Provider data API returned error")

	return nil, endpoint, resp.Code, reqBodytoken, string(resp.Body), fmt.Errorf("get provider data error: %s (status code %d)", response.Message, resp.Code)
}
