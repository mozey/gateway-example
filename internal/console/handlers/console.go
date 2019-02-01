package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mozey/gateway/pkg/response"
	"github.com/mozey/gateway/web/console"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"path"
)

// Bar route handler
func (h *Handler) Console(w http.ResponseWriter, r *http.Request) {
	response.JSON(http.StatusOK, w, r, "todo")
}

// Static route handler
func (h *Handler) Static(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	file := params.ByName("file")
	file = path.Join("static", file)
	f, err := console.VFS.Open(file)
	if err != nil {
		response.JSON(http.StatusInternalServerError, w, r,
			errors.Wrap(err, fmt.Sprintf("open %s", file)))
		return
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		response.JSON(http.StatusInternalServerError, w, r,
			errors.Wrap(err, fmt.Sprintf("read %s", file)))
		return
	}
	response.HTML(http.StatusOK, w, r, b)
	return
}
