package service

import (
	"fmt"

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

func (s *MorphoService) GetVaultPositionByWallet(walletAddress string, chainID int) (*model.MorphoResponse, error) {
	result, err := s.morphoClient.GetVaultPositionByWallet(walletAddress, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault position: %w", err)
	}

	filteredZeroAssetsPosition := make([]model.VaultV2Position, 0)
	for _, position := range result.Data.UserByAddress.VaultV2Positions {
		assets, _ := position.Assets.Float64()
		if assets != 0 {
			filteredZeroAssetsPosition = append(filteredZeroAssetsPosition, position)
		}
	}
	result.Data.UserByAddress.VaultV2Positions = filteredZeroAssetsPosition

	return result, nil
}
