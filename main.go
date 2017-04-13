package main

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"os"
	"github.com/go-kit/kit/log"
)


func main() {
	f, _ := os.OpenFile("/home/pd/service_qxx_log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer f.Close()

	logger := log.NewLogfmtLogger(f)

	var svc QxxService
	svc = qxxService{}

	getcompanycaseHandler := httptransport.NewServer(
		// Server wraps an endpoint and implements http.Handler. 分别传入的参数是endpoint.Endpoint   DecodeRequestFunc   EncodeResponseFunc
		makeGetcompanyEndpoint(svc),
		decodeGetcompanyRequest,
		encodeResponse,
	)

	http.Handle("/getcompany", getcompanycaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}