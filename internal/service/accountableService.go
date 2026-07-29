package service

import (
	"fmt"

	"github.com/Billy19191/billy191-services-backend/internal/client"
	"github.com/Billy19191/billy191-services-backend/internal/model"
	"gorm.io/gorm"
)

type AccountableService struct {
	db                *gorm.DB
	accountableClient *client.AccountableClient
}

func NewAccountableService(db *gorm.DB, accountableClient *client.AccountableClient) *AccountableService {
	return &AccountableService{db: db, accountableClient: accountableClient}
}

func (s *AccountableService) GetPositionAccountableData(walletAddress string) (*model.AccountableResponseEntity, error) {
	result, err := s.accountableClient.GetPositionAccountableData(walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get accountable position data: %w", err)
	}

	return result, nil
}

func (s *AccountableService) GetLoanOverview(userAddress string) (*model.AccountableLoanOverviewEntity, error) {
	result, err := s.accountableClient.GetLoanOverview(userAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get loan overview: %w", err)
	}

	return result, nil
}
