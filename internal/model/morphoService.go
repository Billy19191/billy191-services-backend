package model

import "encoding/json"

type MorphoRequestEntity struct {
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables"`
	OperationName string         `json:"operationName"`
}

type MorphoResponseEntity struct {
	Data MorphoDataEntity `json:"data"`
}

type MorphoDataEntity struct {
	UserByAddress UserByAddressEntity `json:"userByAddress"`
}

type UserByAddressEntity struct {
	VaultV2Positions []VaultV2PositionEntity `json:"vaultV2Positions"`
}

type VaultV2PositionEntity struct {
	Assets json.Number `json:"assets"`
	Shares json.Number `json:"shares"`
	Vault  VaultEntity `json:"vault"`
}

type VaultEntity struct {
	AvgNetApy   float64     `json:"avgNetApy"`
	Name        string      `json:"name"`
	Owner       OwnerEntity `json:"owner"`
	Liquidity   float64     `json:"liquidity"`
	TotalAssets float64     `json:"totalAssets"`
}

type OwnerEntity struct {
	Address string `json:"address"`
}

type MorphoResponseModel struct {
	Data MorphoDataModel `json:"data"`
}

type MorphoDataModel struct {
	Vault []VaultModel `json:"vault"`
}

type VaultModel struct {
	VaultName     string  `json:"vaultName"`
	TotalAssetUsd float64 `json:"totalAssetUsd"`
	Liquidity     float64 `json:"liquidity"`
	MyAssetUsd    float64 `json:"myAssetUsd"`
	AvgApy        float64 `json:"avgApy"`
	SharedInVault float64 `json:"sharedInVault"`
}
