package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Billy19191/billy191-services-backend/internal/model"
)

type AccountableClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAccountableClient() *AccountableClient {
	return &AccountableClient{
		baseURL:    "https://yield.accountable.capital/api/metrics",
		httpClient: &http.Client{},
	}
}

func (c *AccountableClient) GetPositionAccountableData(walletAddress string) (*model.AccountableResponseEntity, error) {
	requestURL := fmt.Sprintf("%s/%s/earn", strings.TrimRight(c.baseURL, "/"), walletAddress)

	req, err := http.NewRequest(http.MethodGet, requestURL, bytes.NewReader(nil))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, string(resBody))
	}

	var accountableResponse model.AccountableResponseEntity
	if err := json.Unmarshal(resBody, &accountableResponse); err == nil {
		return &accountableResponse, nil
	}

	var encodedBody string
	if err := json.Unmarshal(resBody, &encodedBody); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if err := json.Unmarshal([]byte(encodedBody), &accountableResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal nested response: %w", err)
	}

	return &accountableResponse, nil
}
