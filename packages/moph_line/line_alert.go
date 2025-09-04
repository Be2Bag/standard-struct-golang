package moph_line

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
)

func (m *Client) SendLine(ctx context.Context, token string, cid string) (string, int, string, string, error) {

	endpoint := "/api/send-message/template/send-now"
	fullURL := m.url + endpoint

	dataContent := []map[string]any{
		{
			"type":       "flex",
			"header":     "ทดสอบ header",
			"sub_header": "ทดสอบ sub header",
			"text":       "ทดสอบ text",
		},
	}
	reqBody := map[string]any{
		"cid":      cid,
		"template": "alerting",
		"data":     dataContent,
	}
	reqBodyJSON, err := sonic.Marshal(reqBody)
	if err != nil {
		return endpoint, http.StatusInternalServerError, "", "", fmt.Errorf("marshal request body: %w", err)
	}
	reqBodyStr := string(reqBodyJSON)

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
	}

	resp, err := m.http.Post(ctx, fullURL, headers, bytes.NewBuffer(reqBodyJSON), 10)
	if err != nil {
		status := 0
		body := ""
		if resp != nil {
			status = resp.Code
			body = string(resp.Body)
		}
		if ctx.Err() == context.DeadlineExceeded {
			return endpoint, status, reqBodyStr, body, fmt.Errorf("request timed out after %d seconds", 10)
		}
		return endpoint, status, reqBodyStr, body, fmt.Errorf("post request: %w", err)
	}

	respBodyStr := string(resp.Body)

	if resp.Code == http.StatusOK {
		return "success", resp.Code, reqBodyStr, respBodyStr, nil
	}

	var response struct {
		Message     string `json:"message"`
		MessageCode int    `json:"message_code"`
	}
	if err := sonic.Unmarshal(resp.Body, &response); err != nil {
		m.log.WithFields(logrus.Fields{"raw_body": respBodyStr}).Error("unmarshal response failed")
		return endpoint, resp.Code, reqBodyStr, respBodyStr, fmt.Errorf("unmarshal response: %w", err)
	}
	if response.MessageCode == 200 && strings.EqualFold(response.Message, "Success") {
		return "success", resp.Code, reqBodyStr, respBodyStr, nil
	}

	return endpoint, resp.Code, reqBodyStr, respBodyStr, fmt.Errorf("notification failed: status=%d message=%s message_code=%d", resp.Code, response.Message, response.MessageCode)
}
