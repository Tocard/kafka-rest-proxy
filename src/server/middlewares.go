package server

import (
	"net/http"
	"strings"
)

// AddJSONHeader is a middleware to set Content-Type in headers if Accept is application/json.
func AddJSONHeader(res http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/json" {
		res.Header().Add("Content-type", "application/json")
	}
}

// AcceptCORS is a middleware to accept OPTIONS request to get CORS.
func AcceptCORS(res http.ResponseWriter, r *http.Request) {
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, HEAD, TRACE, DELETE, PATCH, COPY, HEAD, LINK, OPTIONS")
	if r.Method == "OPTIONS" {
		for n, h := range r.Header {
			if strings.Contains(n, "Access-Control-Request") {
				for _, h := range h {
					k := strings.Replace(n, "Request", "Allow", 1)
					res.Header().Add(k, h)
				}
			}
		}
		res.WriteHeader(http.StatusNoContent)
		// we should write something to stop martini to respond. This is an "options"
		// for cors, so we MUST NOT respond anything else.
		res.Write([]byte(""))
	}
}
