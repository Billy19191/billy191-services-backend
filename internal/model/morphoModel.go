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
	MarketPositions  []MarketPositionEntity  `json:"marketPositions"`
}

type VaultV2PositionEntity struct {
	Assets json.Number `json:"assets"`
	Shares json.Number `json:"shares"`
	Vault  VaultEntity `json:"vault"`
}

type VaultEntity struct {
	NetApy              float64     `json:"netApy"`
	Name                string      `json:"name"`
	Owner               OwnerEntity `json:"owner"`
	Liquidity           float64     `json:"liquidity"`
	TotalAssets         float64     `json:"totalAssets"`
	DeallocateLiquidity float64     `json:"forceDeallocatableLiquidity"`
}

type OwnerEntity struct {
	Address string `json:"address"`
}

type MarketPositionEntity struct {
	HealthFactor float64           `json:"healthFactor"`
	State        MarketStateEntity `json:"state"`
	Market       MarketEntity      `json:"market"`
}

type MarketStateEntity struct {
	BorrowPnlUsd    float64 `json:"borrowPnlUsd"`
	BorrowAssetsUsd float64 `json:"borrowAssetsUsd"`
	CollateralUsd   float64 `json:"collateralUsd"`
	Collateral      float64 `json:"collateral"`
}

type MarketEntity struct {
	State           MarketApyEntity         `json:"state"`
	CollateralAsset MarketAssetSymbolEntity `json:"collateralAsset"`
	LoanAsset       MarketAssetSymbolEntity `json:"loanAsset"`
}

type MarketApyEntity struct {
	AvgBorrowApy float64 `json:"avgBorrowApy"`
	NetBorrowApy float64 `json:"netBorrowApy"`
}

type MarketAssetSymbolEntity struct {
	Symbol string `json:"symbol"`
}

type MorphoResponseModel struct {
	Data MorphoDataModel `json:"data"`
}

type MorphoDataModel struct {
	Vault  []VaultModel  `json:"vault"`
	Borrow []BorrowModel `json:"borrow"`
}

type VaultModel struct {
	VaultName     string  `json:"vaultName"`
	TotalAssetUsd float64 `json:"totalAssetUsd"`
	Liquidity     float64 `json:"liquidity"`
	MyAssetUsd    float64 `json:"myAssetUsd"`
	NetApy        float64 `json:"netApy"`
	SharedInVault float64 `json:"sharedInVault"`
}

type BorrowModel struct {
	Name                  string  `json:"name"`
	HealthFactor          float64 `json:"healthFactor"`
	BorrowPnlUsd          float64 `json:"borrowPnlUsd"`
	BorrowAssetsUsd       float64 `json:"borrowAssetsUsd"`
	CollateralAssetAmount float64 `json:"collateralAssetAmount"`
	CollateralUsd         float64 `json:"collateralUsd"`
	AvgBorrowApy          float64 `json:"avgBorrowApy"`
	NetBorrowApy          float64 `json:"netBorrowApy"`
	CollateralAsset       string  `json:"collateralAsset"`
	LoanAsset             string  `json:"loanAsset"`
}
