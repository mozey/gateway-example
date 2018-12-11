package middleware

import (
	"fmt"
	"github.com/mozey/gateway/pkg/response"
	"github.com/rs/zerolog/log"
	"net/http"
	"runtime/debug"
)

type PanicHandlerFunc func(http.ResponseWriter, *http.Request, interface{})

type PanicHandlerOptions struct {
	PrintStack bool
}

func PanicHandler(o *PanicHandlerOptions) PanicHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, rcv interface{}) {
		response.JSON(http.StatusInternalServerError, w, r, response.Response{
			Message: fmt.Sprintf("%s", rcv),
		})
		log.Error().Msg(rcv.(string))
		if o.PrintStack {
			debug.PrintStack()
		}
	}
}
