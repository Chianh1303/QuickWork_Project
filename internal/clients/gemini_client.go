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
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=%s", apiKey)

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

	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("Gemini request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read Gemini response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("Gemini returned HTTP %d: %s", resp.StatusCode, string(body))
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
		return "", fmt.Errorf("invalid response from gemini")
	}

	rawText := strings.TrimSpace(geminiResp.Candidates[0].Content.Parts[0].Text)
	rawText = strings.TrimPrefix(rawText, "```json")
	rawText = strings.TrimPrefix(rawText, "```")
	rawText = strings.TrimSuffix(rawText, "```")
	rawText = strings.TrimSpace(rawText)

	return rawText, nil
}

func (c *geminiClient) GenerateText(apiKey string, prompt string) (string, error) {
	return c.EvaluateCV(apiKey, prompt, nil)
}
