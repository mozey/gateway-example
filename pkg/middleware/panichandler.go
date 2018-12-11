package middleware

import (
	"fmt"
	"github.com/mozey/gateway/pkg/response"
	"net/http"
	"runtime"
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
		if o.PrintStack {
			var stackSize = 4 << 10 // 4 KB
			b := make([]byte, stackSize)
			length := runtime.Stack(b, false)
			fmt.Println(string(b[:length]))
			// TODO Use zerolog to print stack trace
			// https://github.com/rs/zerolog/pull/35/files
		}
	}
}
