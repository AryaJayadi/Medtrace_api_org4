package handlers

import (
	"net/http"

	"github.com/AryaJayadi/MedTrace_api_org4/internal/services"
	"github.com/labstack/echo/v4"
)

type OrganizationHandler struct {
	Service *services.OrganizationService
}

func NewOrganizationHandler(service *services.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{Service: service}
}

// GetOrganizations godoc
// @Summary Get all organizations
// @Description Retrieve all organizations from the ledger
// @Tags organizations
// @Produce json
// @Success 200 {object} response.BaseListResponse[entity.Organization]
// @Failure 401 {object} response.BaseResponse "Unauthorized - JWT invalid or missing"
// @Failure 500 {object} response.BaseResponse "Internal server error or Fabric error"
// @Router /organizations [get]
// @Security BearerAuth
func (h *OrganizationHandler) GetOrganizations(c echo.Context) error {
	resp := h.Service.GetOrganizations(c.Request().Context())

	status := http.StatusOK
	if !resp.Success {
		status = resp.Error.Code
		if status == 0 {
			status = http.StatusInternalServerError
		}
	}

	return c.JSON(status, resp)
}
