package main

import (
	"encoding/json"
	"net/http"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type getcompanyRequest struct {
	S string `json:"s"`
}

type getcompanyResponse struct {
	Result   string `json:"resp"`
	Err string `json:"err,omitempty"`
}

func decodeGetcompanyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getcompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeGetcompanyEndpoint(svc QxxService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getcompanyRequest)
		resp, err := svc.Getcompany(req.S)
		if err != nil {
			return getcompanyResponse{resp, err.Error()}, nil
		}
		return getcompanyResponse{resp, ""}, nil
	}
}