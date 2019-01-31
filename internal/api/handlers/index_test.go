package handlers_test

import (
	"github.com/mozey/gateway/internal/api/handlers"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestHandler_Index(t *testing.T) {
	// Create shared handler
	h := handlers.NewHandler(conf)
	defer h.Cleanup()

	// Record handler response
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	h.Index(rec, req)

	// Verify response
	dump, err := httputil.DumpResponse(rec.Result(), true)
	require.NoError(t, err)
	log.Debug().Msg(string(dump))
	require.Equal(t, rec.Code, http.StatusOK, "invalid status code")
}
