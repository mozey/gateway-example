package controllers

import (
	"net/http"
	"io/ioutil"
	"io"
	"log"
	"net/url"
	"github.com/mozey/gateway/internal/middleware"
)

const kilobyte = 1024

// EchoRequest is the same as http.Request minus the bits that break json.Marshall
type EchoRequest struct {
	Method     string
	URL        *url.URL
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0
	Header     http.Header
	Body struct {
		String string
	}
	ContentLength    int64
	TransferEncoding []string
	Host             string
	//Form url.Values
	//PostForm url.Values
	//MultipartForm *multipart.Form
	Trailer    http.Header
	RemoteAddr string
	RequestURI string
	//TLS *tls.ConnectionState
}

// Echo the http request
func Echo(w http.ResponseWriter, r *http.Request) {
	e := EchoRequest{}
	e.Method = r.Method
	e.URL = r.URL
	e.Proto = r.Proto
	e.ProtoMajor = r.ProtoMajor
	e.ProtoMinor = r.ProtoMinor
	e.Header = r.Header
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1*kilobyte))
	if err != nil {
		log.Panic(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Panic(err)
	}
	e.Body.String = string(body)
	e.ContentLength = r.ContentLength
	e.TransferEncoding = r.TransferEncoding
	e.Host = r.Host
	e.Trailer = r.Trailer
	e.RemoteAddr = r.RemoteAddr
	e.RequestURI = r.RequestURI

	middleware.RespondWithCode(http.StatusNotFound, w, r, e)
}
