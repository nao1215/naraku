package api

import (
	"net/http"

	"github.com/go-spectest/naraku/app/domain/model"
	"github.com/labstack/echo/v4"
)

// OSController is a controller for /os API.
type OSController struct{}

// NewOSController returns a new OSController struct.
func NewOSController() *OSController {
	return &OSController{}
}

// OSResponse is response for GET /os
type OSResponse struct {
	// @Description	Items is os information list.
	Items []model.OS `json:"items"`
}

// handleOS handle /os API.
func (ctrl *OSController) handleGetOS(c echo.Context) error {
	distribution := c.Param("distribution")
	switch distribution {
	case "android":
		return ctrl.androidInfo(c)
	default:
		return c.String(http.StatusNotFound, "Unknown OS")
	}
}

// androidInfo return android information.
//
// @Summary		Get android information
// @Description	This API is for getting android information.
// @Tags		os
// @Produce		json
// @Success		200	{object}	OSResponse
// @Router		/os/android [get]
func (ctrl *OSController) androidInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
