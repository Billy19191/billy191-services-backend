package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Billy19191/billy191-services-backend/internal/model"
)

type MorphoClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewMorphoClient() *MorphoClient {
	return &MorphoClient{
		baseURL:    "https://api.morpho.org/graphql",
		httpClient: &http.Client{},
	}
}

func (c *MorphoClient) GetVaultPositionByWallet(address string, chainId int) (*model.MorphoResponse, error) {
	const vaultPositionQuery = `query UserByAddress($address: String!, $chainId: Int) {
		userByAddress(address: $address, chainId: $chainId) {
			vaultV2Positions {
			assets
			vault {
				avgNetApy
				name
				owner {
				address
				}
				liquidity
				totalAssets
			}
			shares
			}
		}
	}`
	reqBody := model.MorphoRequest{
		Query: vaultPositionQuery,
		Variables: map[string]any{
			"address": address,
			"chainId": chainId,
		},
		OperationName: "UserByAddress",
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.baseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

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

	var morphoResponse model.MorphoResponse

	errMapModel := json.Unmarshal(resBody, &morphoResponse)

	if errMapModel != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", errMapModel)
	}

	return &morphoResponse, nil
}
