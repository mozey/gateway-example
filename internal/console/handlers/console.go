package handlers

import (
	"fmt"
	"github.com/mozey/gateway/pkg/response"
	"github.com/mozey/gateway/web/console"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Bar route handler
func (h *Handler) Console(w http.ResponseWriter, r *http.Request) {
	response.JSON(http.StatusOK, w, r, "todo")
}

// Static route handler
func (h *Handler) Static(w http.ResponseWriter, r *http.Request) {
	asset := strings.Replace(
		r.URL.Path, "console", "", 1)
	file, err := console.VFS.Open(asset)
	if err != nil {
		response.JSON(http.StatusInternalServerError, w, r,
			errors.Wrap(err, fmt.Sprintf("open %s", asset)))
		return
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		response.JSON(http.StatusInternalServerError, w, r,
			errors.Wrap(err, fmt.Sprintf("read %s", asset)))
		return
	}
	response.HTML(http.StatusOK, w, r, b)
	return
}
