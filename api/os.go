package api

import (
	"net/http"
	"time"

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
	//	@Description	Items is os information list.
	Items []model.OS `json:"items"`
}

// handleOS handle /os API.
func (ctrl *OSController) handleGetOS(c echo.Context) error {
	distribution := c.Param("distribution")
	switch distribution {
	case "android":
		return ctrl.androidInfo(c)
	case "debian":
		return ctrl.debianInfo(c)
	default:
		return c.String(http.StatusNotFound, "Unknown OS")
	}
}

// androidInfo return android information.
//
//	@Summary		Get android information
//	@Description	This API is for getting android information.
//	@Tags			os
//	@Produce		json
//	@Success		200	{object}	OSResponse
//	@Router			/os/android [get]
//
// TODO: This is mock code. Implement the datastore code.
func (ctrl *OSController) androidInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, &OSResponse{
		Items: []model.OS{
			{
				Name:     "Android",
				Version:  "1.0",
				CodeName: "",
				Release:  time.Date(2008, 9, 23, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{1},
			},
			{
				Name:     "Android",
				Version:  "1.1",
				CodeName: "",
				Release:  time.Date(2009, 2, 9, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{2},
			},
			{
				Name:     "Android",
				Version:  "1.5",
				CodeName: "Cupcake",
				Release:  time.Date(2009, 4, 27, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{3},
			},
			{
				Name:     "Android",
				Version:  "1.6",
				CodeName: "Donut",
				Release:  time.Date(2009, 9, 15, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{4},
			},
			{
				Name:     "Android",
				Version:  "2.0 - 2.1",
				CodeName: "Eclair",
				Release:  time.Date(2009, 10, 26, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{5, 6, 7},
			},
			{
				Name:     "Android",
				Version:  "2.2 - 2.2.3",
				CodeName: "Froyo",
				Release:  time.Date(2010, 5, 20, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{8},
			},
			{
				Name:     "Android",
				Version:  "2.3 - 2.3.7",
				CodeName: "Gingerbread",
				Release:  time.Date(2010, 12, 6, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{9, 10},
			},
			{
				Name:     "Android",
				Version:  "3.0 - 3.2.6",
				CodeName: "Honeycomb",
				Release:  time.Date(2011, 2, 22, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{11, 12, 13},
			},
			{
				Name:     "Android",
				Version:  "4.0 - 4.0.4",
				CodeName: "Ice Cream Sandwich",
				Release:  time.Date(2011, 10, 18, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{14, 15},
			},
			{
				Name:     "Android",
				Version:  "4.1 - 4.3.1",
				CodeName: "Jelly Bean",
				Release:  time.Date(2012, 7, 9, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{16, 17, 18},
			},
			{
				Name:     "Android",
				Version:  "4.4 - 4.4.4",
				CodeName: "KitKat",
				Release:  time.Date(2013, 10, 31, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{19},
			},
			{
				Name:     "Android",
				Version:  "4.4W - 4.4W.2",
				CodeName: "4.4W",
				Release:  time.Date(2014, 6, 25, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{20},
			},
			{
				Name:     "Android",
				Version:  "5.0 - 5.1.1",
				CodeName: "Lollipop",
				Release:  time.Date(2014, 11, 12, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{21, 22},
			},
			{
				Name:     "Android",
				Version:  "6.0 - 6.0.1",
				CodeName: "Marshmallow",
				Release:  time.Date(2015, 10, 5, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{23},
			},
			{
				Name:     "Android",
				Version:  "7.0 - 7.1.2",
				CodeName: "Nougat",
				Release:  time.Date(2016, 8, 22, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{24, 25},
			},
			{
				Name:     "Android",
				Version:  "8.0 - 8.1",
				CodeName: "Oreo",
				Release:  time.Date(2017, 8, 21, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{26, 27},
			},
			{
				Name:     "Android",
				Version:  "9",
				CodeName: "Pie",
				Release:  time.Date(2018, 8, 6, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{28},
			},
			{
				Name:     "Android",
				Version:  "10",
				CodeName: "Q",
				Release:  time.Date(2019, 9, 3, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{29},
			},
			{
				Name:     "Android",
				Version:  "11",
				CodeName: "R",
				Release:  time.Date(2020, 9, 8, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{30},
			},
			{
				Name:     "Android",
				Version:  "12",
				CodeName: "S",
				Release:  time.Date(2021, 9, 8, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{31},
			},
			{
				Name:     "Android",
				Version:  "12L",
				CodeName: "Sv2",
				Release:  time.Date(2022, 3, 7, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{32},
			},
			{
				Name:     "Android",
				Version:  "13",
				CodeName: "Tiramisu",
				Release:  time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{33},
			},
			{
				Name:     "Android",
				Version:  "14",
				CodeName: "Upside Down Cake",
				Release:  time.Date(2023, 10, 4, 0, 0, 0, 0, time.UTC),
				APILevel: []uint{34},
			},
		},
	})
}

// debianInfo return debian information.
//
//	@Summary		Get debian information
//	@Description	This API is for getting debian information.
//	@Tags			os
//	@Produce		json
//	@Success		200	{object}	OSResponse
//	@Router			/os/debian [get]
//
// TODO: This is mock code. Implement the datastore code.
func (ctrl *OSController) debianInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, &OSResponse{
		Items: []model.OS{
			{
				Name:     "Debian",
				Version:  "1.1",
				CodeName: "Buzz",
				Release:  time.Date(1996, 6, 16, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "1.2",
				CodeName: "Rex",
				Release:  time.Date(1996, 12, 12, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "1.3",
				CodeName: "Bo",
				Release:  time.Date(1997, 6, 5, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "2.0",
				CodeName: "Hamm",
				Release:  time.Date(1998, 7, 24, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "2.1",
				CodeName: "Slink",
				Release:  time.Date(1999, 3, 9, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "2.2",
				CodeName: "Potato",
				Release:  time.Date(2000, 8, 15, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "3.0",
				CodeName: "Woody",
				Release:  time.Date(2002, 7, 19, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "3.1",
				CodeName: "Sarge",
				Release:  time.Date(2005, 6, 6, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "4.0",
				CodeName: "Etch",
				Release:  time.Date(2007, 4, 8, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "5.0",
				CodeName: "Lenny",
				Release:  time.Date(2009, 2, 14, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "6.0",
				CodeName: "Squeeze",
				Release:  time.Date(2011, 2, 6, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "7.0",
				CodeName: "Wheezy",
				Release:  time.Date(2013, 5, 4, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "8.0",
				CodeName: "Jessie",
				Release:  time.Date(2015, 4, 25, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "9.0",
				CodeName: "Stretch",
				Release:  time.Date(2017, 6, 17, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "10.0",
				CodeName: "Buster",
				Release:  time.Date(2019, 7, 6, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "11.0",
				CodeName: "Bullseye",
				Release:  time.Date(2021, 8, 14, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:     "Debian",
				Version:  "12.0",
				CodeName: "Bookworm",
				Release:  time.Date(2023, 6, 10, 0, 0, 0, 0, time.UTC),
			},
		},
	})
}
