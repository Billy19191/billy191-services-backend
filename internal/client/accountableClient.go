package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/Billy19191/billy191-services-backend/internal/model"
)

type AccountableClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAccountableClient() *AccountableClient {
	return &AccountableClient{
		baseURL:    "https://yield.accountable.capital/api",
		httpClient: &http.Client{},
	}
}

func (c *AccountableClient) GetPositionAccountableData(walletAddress string) (*model.AccountableResponseEntity, error) {
	var (
		wg         sync.WaitGroup
		earnBody   []byte
		earnErr    error
		loanResult *model.AccountableLoanOverviewEntity
		loanErr    error
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		requestURL := fmt.Sprintf("%s/metrics/%s/earn", strings.TrimRight(c.baseURL, "/"), walletAddress)

		req, err := http.NewRequest(http.MethodGet, requestURL, bytes.NewReader(nil))
		if err != nil {
			earnErr = fmt.Errorf("failed to create request: %w", err)
			return
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			earnErr = fmt.Errorf("failed to send request: %w", err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			earnErr = fmt.Errorf("failed to read response: %w", err)
			return
		}

		if res.StatusCode != http.StatusOK {
			earnErr = fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, string(body))
			return
		}

		earnBody = body
	}()

	go func() {
		defer wg.Done()
		loanResult, loanErr = c.GetLoanOverview(walletAddress)
	}()

	wg.Wait()

	if earnErr != nil {
		return nil, earnErr
	}
	if loanErr != nil {
		return nil, fmt.Errorf("failed to get loan overview: %w", loanErr)
	}

	var accountableResponse model.AccountableResponseEntity
	if err := json.Unmarshal(earnBody, &accountableResponse); err == nil {
		if len(accountableResponse.VaultAllocations) > 0 && loanResult.TotalDeposits > 0 {
			accountableResponse.VaultAllocations[0].TotalInvestedInVaultUsd = loanResult.LoopedTvlUsd
			return &accountableResponse, nil
		} else {
			return nil, fmt.Errorf("no vault allocations found")
		}
	}

	var encodedBody string
	if err := json.Unmarshal(earnBody, &encodedBody); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if err := json.Unmarshal([]byte(encodedBody), &accountableResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal nested response: %w", err)
	}

	return &accountableResponse, nil
}

func (c *AccountableClient) GetLoanOverview(userAddress string) (*model.AccountableLoanOverviewEntity, error) {
	requestURL := fmt.Sprintf("%s/users/%s/personal-data-overview", strings.TrimRight(c.baseURL, "/"), userAddress)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
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

	var loanOverview model.AccountableLoanOverviewEntity
	if err := json.Unmarshal(resBody, &loanOverview); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &loanOverview, nil
}
