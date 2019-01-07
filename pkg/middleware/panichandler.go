package middleware

import (
	"fmt"
	"github.com/mozey/gateway/pkg/response"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

type PanicHandlerFunc func(http.ResponseWriter, *http.Request, interface{})

type PanicHandlerOptions struct {
	PrintStack bool
}

func PanicHandler(o *PanicHandlerOptions) PanicHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, rcv interface{}) {
		err := fmt.Errorf("%s", rcv)
		response.JSON(http.StatusInternalServerError, w, r, response.Response{
			Message: err.Error(),
		})
		if o.PrintStack {
			// Use zerolog to print stack trace
			// https://github.com/rs/zerolog/pull/35
			err := errors.Wrap(err, "recovered panic")
			log.Error().Stack().Err(err).Msg("-")
		}
	}
}
