package model

import "encoding/json"

type MorphoRequest struct {
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables"`
	OperationName string         `json:"operationName"`
}

type MorphoResponse struct {
	Data MorphoData `json:"data"`
}

type MorphoData struct {
	UserByAddress UserByAddress `json:"userByAddress"`
}

type UserByAddress struct {
	VaultV2Positions []VaultV2Position `json:"vaultV2Positions"`
}

type VaultV2Position struct {
	Assets json.Number `json:"assets"`
	Shares json.Number `json:"shares"`
	Vault  Vault       `json:"vault"`
}

type Vault struct {
	AvgNetApy   float64 `json:"avgNetApy"`
	Name        string  `json:"name"`
	Owner       Owner   `json:"owner"`
	Liquidity   float64 `json:"liquidity"`
	TotalAssets float64 `json:"totalAssets"`
}

type Owner struct {
	Address string `json:"address"`
}
