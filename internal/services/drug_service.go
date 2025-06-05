package services

import (
	"context"
	"encoding/json"

	"github.com/AryaJayadi/MedTrace_api_org4/internal/models"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type DrugService struct {
	contract *client.Contract
}

func NewDrugService(contract *client.Contract) *DrugService {
	return &DrugService{contract: contract}
}

func (s *DrugService) GetHistoryDrug(ctx context.Context, drugID string) models.BaseListResponse[models.HistoryDrug] {
	resultBytes, err := s.contract.EvaluateTransaction("GetHistoryDrug", drugID)
	if err != nil {
		return models.ErrorListResponse[models.HistoryDrug](500, "Failed to evaluate GetHistoryDrug transaction: %v", err)
	}

	var records []models.HistoryDrug
	err = json.Unmarshal(resultBytes, &records)
	if err != nil {
		return models.ErrorListResponse[models.HistoryDrug](500, "Failed to unmarshal history drug data for GetHistoryDrug: %v", err)
	}

	recordsPtrs := make([]*models.HistoryDrug, len(records))
	for i := range records {
		recordsPtrs[i] = &records[i]
	}

	return models.SuccessListResponse(recordsPtrs)
}
