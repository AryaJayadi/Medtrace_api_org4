package services

import (
	"context"
	"encoding/json"

	"github.com/AryaJayadi/MedTrace_api_org4/internal/models"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type DrugService struct{}

func NewDrugService() *DrugService {
	return &DrugService{}
}

func (s *DrugService) GetHistoryDrug(contract *client.Contract, ctx context.Context, drugID string) models.BaseListResponse[models.HistoryDrug] {
	resultBytes, err := contract.EvaluateTransaction("GetHistoryDrug", drugID)
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
