package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/pkg/handler"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Handler for  mozey/gateway
type Handler struct {
	*hutil.Handler
	Config *config.Config
	Router *httprouter.Router
}

// Response to be used or extended by handlers
type Response struct {
	Message string `json:"message"`
}

// RespondJSON can be used by route handler to respond to requests
func RespondJSON(w http.ResponseWriter, r *http.Request, i interface{}) {
	j, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		log.Panic().Err(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(j))
}

// NewHandler creates a new handler and initialises services
// that are shared between handlers.
// Remember to close services like the database connection by
// calling h.Cleanup before the application exits
func NewHandler(conf *config.Config) (h *Handler) {
	h = &Handler{}
	h.Config = conf
	h.Handler = hutil.NewHandler(&hutil.Config{
		Debug:      conf.Debug(),
		Region:     conf.Region(),
		AwsProfile: conf.AwsProfile(),
	})
	h.Router = httprouter.New()
	return h
}
