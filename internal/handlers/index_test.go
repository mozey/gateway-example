package handlers_test

import (
	"github.com/labstack/echo"
	"github.com/mozey/gateway/internal/handlers"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestHandler_Index(t *testing.T) {
	// Create shared handler
	e := echo.New()
	h := handlers.NewHandler(e, conf)
	defer h.Cleanup()

	// Create request
	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	// Record handler response
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	require.NoError(t, h.Index(c))

	// Verify response
	dump, err := httputil.DumpResponse(rec.Result(), true)
	require.NoError(t, err)
	h.Debugj(c, "raw response ", dump)
	require.Equal(t, rec.Code, http.StatusOK, "invalid status code")
}
