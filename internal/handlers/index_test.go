package handlers

import (
	"github.com/labstack/echo"
	"github.com/mozey/logutil"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestIndex(t *testing.T) {
	// Create request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record handler response
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	Index(c)

	// Verify response
	dump, err := httputil.DumpResponse(rec.Result(), true)
	require.NoError(t, err)
	logutil.Debug(string(dump)) // Print raw response
	require.Equal(t, rec.Code, http.StatusOK, "invalid status code")
}
