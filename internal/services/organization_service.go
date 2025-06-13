package services

import (
	"context"
	"encoding/json"

	"github.com/AryaJayadi/MedTrace_api_org4/internal/models"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type OrganizationService struct {
	contract *client.Contract
}

func NewOrganizationService(contract *client.Contract) *OrganizationService {
	return &OrganizationService{contract: contract}
}

// GetOrganizations retrieves all organizations from the ledger using the provided contract.
func (s *OrganizationService) GetOrganizations(ctx context.Context) models.BaseListResponse[models.Organization] {
	resp, err := s.contract.EvaluateTransaction("GetAllOrganizations")
	if err != nil {
		return models.ErrorListResponse[models.Organization](500, "Failed to evaluate transaction to Fabric: %v", err)
	}

	var organizations []models.Organization
	if err := json.Unmarshal(resp, &organizations); err != nil {
		return models.ErrorListResponse[models.Organization](500, "Failed to unmarshal Fabric response: %v", err)
	}

	ptrList := make([]*models.Organization, len(organizations))
	for i := range organizations {
		ptrList[i] = &organizations[i]
	}

	return models.SuccessListResponse(ptrList)
}
