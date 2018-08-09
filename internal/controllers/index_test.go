package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/require"
	"github.com/mozey/gateway/internal/middleware"
	"github.com/mozey/logutil"
	"net/http/httputil"
)

func TestIndex(t *testing.T) {
	// Create request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record handler response
	rr := httptest.NewRecorder()
	var handler http.Handler
	handler = http.HandlerFunc(Index)
	handler = middleware.ResponseHeaders(handler)
	handler.ServeHTTP(rr, req)

	// Verify response
	dump, err := httputil.DumpResponse(rr.Result(), true)
	require.NoError(t, err)
	logutil.Debug(string(dump)) // Print raw response
	require.Equal(t, rr.Code, http.StatusOK, "invalid status code")
}

