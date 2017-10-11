package main

import (
	"fmt"
	"github.com/omise/omise-go"
	"net/http"
)

type TamboonHandler struct {
	client *omise.Client
}

func (handler *TamboonHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	method, path := req.Method, req.URL.Path
	fmt.Printf("%s %s\n", method, path)

	if method == "GET" && path == "/" {
		handler.GetCharityList(resp, req)

	} else if method == "POST" && path == "/donate" {
        handler.PostDonate(resp, req)
        
	} else {
		http.NotFound(resp, req)

	}
}