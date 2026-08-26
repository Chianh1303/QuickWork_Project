package clients

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type GeminiClient interface {
	EvaluateCV(apiKey string, prompt string, pdfBytes []byte) (string, error)
	GenerateText(apiKey string, prompt string) (string, error)
}

type geminiClient struct {
	httpClient *http.Client
}

func NewGeminiClient() GeminiClient {
	return &geminiClient{
		httpClient: &http.Client{
			Timeout: 90 * time.Second,
		},
	}
}

func (c *geminiClient) EvaluateCV(apiKey string, prompt string, pdfBytes []byte) (string, error) {
	modelsList := []string{
		"gemini-2.5-flash",
		"gemini-2.0-flash",
		"gemini-1.5-flash",
	}

	var lastErr error
	for _, model := range modelsList {
		url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, apiKey)

		partsList := []map[string]interface{}{}
		if len(pdfBytes) > 0 {
			base64PDF := base64.StdEncoding.EncodeToString(pdfBytes)
			partsList = append(partsList, map[string]interface{}{
				"inline_data": map[string]interface{}{
					"mime_type": "application/pdf",
					"data":      base64PDF,
				},
			})
		}
		partsList = append(partsList, map[string]interface{}{"text": prompt})

		requestBody, err := json.Marshal(map[string]interface{}{
			"contents": []map[string]interface{}{{"parts": partsList}},
		})
		if err != nil {
			return "", fmt.Errorf("failed to build Gemini request: %w", err)
		}

		for attempt := 1; attempt <= 3; attempt++ {
			resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(requestBody))
			if err != nil {
				lastErr = fmt.Errorf("Gemini request failed: %w", err)
				time.Sleep(time.Duration(attempt) * 1 * time.Second)
				continue
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				lastErr = fmt.Errorf("failed to read Gemini response: %w", err)
				time.Sleep(time.Duration(attempt) * 1 * time.Second)
				continue
			}

			if resp.StatusCode == 503 || resp.StatusCode == 429 {
				lastErr = fmt.Errorf("Gemini returned HTTP %d: %s", resp.StatusCode, string(body))
				time.Sleep(time.Duration(attempt) * 1 * time.Second)
				continue
			}

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				lastErr = fmt.Errorf("Gemini returned HTTP %d: %s", resp.StatusCode, string(body))
				break // Not temporary 503, try fallback model
			}

			var geminiResp struct {
				Candidates []struct {
					Content struct {
						Parts []struct {
							Text string `json:"text"`
						} `json:"parts"`
					} `json:"content"`
				} `json:"candidates"`
			}

			if err := json.Unmarshal(body, &geminiResp); err != nil || len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
				lastErr = fmt.Errorf("invalid response from gemini")
				break
			}

			rawText := strings.TrimSpace(geminiResp.Candidates[0].Content.Parts[0].Text)
			rawText = strings.TrimPrefix(rawText, "```json")
			rawText = strings.TrimPrefix(rawText, "```")
			rawText = strings.TrimSuffix(rawText, "```")
			rawText = strings.TrimSpace(rawText)

			return rawText, nil
		}
	}

	return "", lastErr
}

func (c *geminiClient) GenerateText(apiKey string, prompt string) (string, error) {
	return c.EvaluateCV(apiKey, prompt, nil)
}
