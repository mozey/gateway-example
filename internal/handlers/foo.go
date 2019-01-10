package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mozey/gateway/internal/config"
	"github.com/mozey/gateway/pkg/middleware"
	"github.com/mozey/gateway/pkg/response"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Foo route handler
func (h *Handler) Foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := httprouter.ParamsFromContext(r.Context())
	foo := params.ByName("foo")
	if foo == "" {
		response.JSON(http.StatusUnauthorized, w, r, response.Response{
			Message: "missing foo in the query string",
		})
		return
	}
	if foo == "panic" {
		//time.Sleep(1 * time.Second)
		// Pass in foo=panic to see the middleware.RecoveryHandler in action
		log.Ctx(ctx).Panic().Msg("oops!")
	}
	if foo == "config" {
		conf := config.New()
		resp := response.Response{
			Message: fmt.Sprintf("conf.Debug %v", conf.Debug())}
		response.JSON(http.StatusOK, w, r, resp)
		return
	}

	// Auth middleware sets user on the context
	user, ok := r.Context().Value(middleware.User{}).(middleware.User)
	if !ok {
		response.JSON(http.StatusInternalServerError, w, r, response.Response{
			Message: "user not set",
		})
		return
	}

	resp := response.Response{
		Message: fmt.Sprintf("foo: %v user: %v", foo, user.Name)}
	response.JSON(http.StatusOK, w, r, resp)
	return
}
