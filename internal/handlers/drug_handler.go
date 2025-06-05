package handlers

import (
	"net/http"

	"github.com/AryaJayadi/MedTrace_api_org4/internal/models"
	"github.com/AryaJayadi/MedTrace_api_org4/internal/services"
	"github.com/labstack/echo/v4"
)

type DrugHandler struct {
	Service *services.DrugService
}

func NewDrugHandler(service *services.DrugService) *DrugHandler {
	return &DrugHandler{Service: service}
}

// GetHistoryDrug godoc
// @Summary Get history for drugs
// @Description Retrieve history records for all drugs, including creation and deletion events.
// @Tags drugs
// @Produce json
// @Success 200 {object} response.BaseListResponse[entity.HistoryDrug]
// @Failure 401 {object} response.BaseResponse "Unauthorized - JWT invalid or missing"
// @Failure 500 {object} response.BaseResponse "Internal server error or Fabric error"
// @Router /history/drug/{drugID} [get]
func (h *DrugHandler) GetHistoryDrug(c echo.Context) error {
	drugID := c.Param("drugID")
	if drugID == "" {
		resp := models.ErrorListResponse[models.HistoryDrug](http.StatusBadRequest, "Drug ID parameter is required")
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := h.Service.GetHistoryDrug(c.Request().Context(), drugID)

	status := http.StatusOK
	if !resp.Success {
		status = resp.Error.Code
		if status == 0 {
			status = http.StatusInternalServerError
		}
	}
	return c.JSON(status, resp)
}
