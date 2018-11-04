package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

// Pass message
const Pass = "ok"
// Skip message
const Skip = "-"

// StatusResponse for normal operation fields will be set to Pass,
// if the check was not performed the field is set to Skip,
// otherwise the response will contain an error message
type StatusResponse struct {
	Connectivity string `json:"connectivity"`
}

// Status can be used to check if services the server depends on are available
func (h *Handler) Status(c echo.Context) error {
	resp := StatusResponse{}
	code := http.StatusOK

	check := map[string]bool{
		"connectivity": false,
	}

	metrics := c.QueryParam("metrics")
	if metrics != "" {
		// Only show status for the specified metrics
		for _, m := range strings.Split(c.QueryParam("metrics"), ",") {
			if _, ok := check[m]; ok {
				check[m] = true
			}
		}
	} else {
		// Check all
		for m := range check {
			check[m] = true
		}
	}

	// Connectivity
	if check["connectivity"] {
		u := "http://example.com/"
		getResp, err := http.Get(u)
		if err != nil {
			code = http.StatusBadGateway
			resp.Connectivity = err.Error()
		} else if getResp.StatusCode != http.StatusOK {
			code = http.StatusBadGateway
			resp.Connectivity = fmt.Sprintf(
				"%v status code %v", u, getResp.StatusCode)
		} else {
			resp.Connectivity = Pass
		}
	} else {
		resp.Connectivity = Skip
	}

	return c.JSON(code, resp)
}
