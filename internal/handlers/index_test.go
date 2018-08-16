package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"github.com/stretchr/testify/require"
	"github.com/labstack/echo"
	"fmt"
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
	// TODO Why is log not printed with -v flag?
	//logutil.Debug(string(dump)) // Print raw response
	fmt.Println(string(dump))
	require.Equal(t, rec.Code, http.StatusOK, "invalid status code")
}
