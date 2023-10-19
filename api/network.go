package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NetworkController is a controller for network API.
type NetworkController struct{}

// NewNetworkController returns a new NetworkController struct.
func NewNetworkController() *NetworkController {
	return &NetworkController{}
}

// IPResponse is response for GET /ip
type IPResponse struct {
	//	@Description GlobalIP is client ip address.
	GlobalIP string `json:"global_ip"`
}

// ip return client global ip address.
//
//	@Summary		Get client global ip address
//	@Description	This API is for getting client global ip address.
//	@Tags			network
//	@Produce		json
//	@Success		200	{object}	IPResponse
//	@Router			/ip [get]
func (ctrl *NetworkController) ip(c echo.Context) error {
	return c.JSON(http.StatusOK, &IPResponse{
		GlobalIP: c.RealIP(),
	})
}
