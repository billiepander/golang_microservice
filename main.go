package main

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)


func main() {
	var svc QxxService
	svc = qxxService{}

	getcompanycaseHandler := httptransport.NewServer(
		// Server wraps an endpoint and implements http.Handler. 分别传入的参数是endpoint.Endpoint   DecodeRequestFunc   EncodeResponseFunc
		makeGetcompanyEndpoint(svc),
		decodeGetcompanyRequest,
		encodeResponse,
	)

	http.Handle("/getcompany", getcompanycaseHandler)
	http.ListenAndServe(":8080", nil)
}