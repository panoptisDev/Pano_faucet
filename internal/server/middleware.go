package server

import (
	"net/http"
	"os"
)

type Redirect struct{}

func (redirect *Redirect) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	proto := req.Header.Get("X-Forwarded-Proto")
	if os.Getenv("ENV") != "production" || proto == "https" {
		next(w, req)
		return
	}
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func NewRedirect() *Redirect {
	return &Redirect{}
}
