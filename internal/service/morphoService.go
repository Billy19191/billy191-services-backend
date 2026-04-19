package service

import (
	"fmt"
	"math"

	"github.com/Billy19191/billy191-services-backend/internal/client"
	"github.com/Billy19191/billy191-services-backend/internal/model"
	"gorm.io/gorm"
)

type MorphoService struct {
	db           *gorm.DB
	morphoClient *client.MorphoClient
}

func NewMorphoService(db *gorm.DB, morphoClient *client.MorphoClient) *MorphoService {
	return &MorphoService{
		db:           db,
		morphoClient: morphoClient,
	}
}

func (s *MorphoService) GetVaultPositionByWallet(walletAddress string, chainID int) (*model.MorphoResponseModel, error) {
	result, err := s.morphoClient.GetVaultPositionByWallet(walletAddress, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault position: %w", err)
	}

	var vaults []model.VaultModel
	for _, position := range result.Data.UserByAddress.VaultV2Positions {
		assets, err := position.Assets.Float64()
		if err != nil {
			return nil, fmt.Errorf("failed to convert position assets to float64: %w", err)
		}
		if assets == 0 {
			continue
		}

		shareInVault := assets / position.Vault.TotalAssets

		vaults = append(vaults, model.VaultModel{
			VaultName:     position.Vault.Name,
			TotalAssetUsd: position.Vault.TotalAssets / math.Pow(10, 6),
			Liquidity:     position.Vault.Liquidity / math.Pow(10, 6),
			MyAssetUsd:    assets / math.Pow(10, 6),
			NetApy:        position.Vault.NetApy * 100,
			SharedInVault: shareInVault * 100,
		})
	}

	return &model.MorphoResponseModel{
		Data: model.MorphoDataModel{
			Vault: vaults,
		},
	}, nil
}
