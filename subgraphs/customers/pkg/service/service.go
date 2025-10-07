package service

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/wundergraph/service/pkg/data"
	v1 "github.com/wundergraph/service/pkg/generated/service/v1"
	v1connect "github.com/wundergraph/service/pkg/generated/service/v1/v1connect"
)

type Service struct{}

// LookupAddressById implements v1connect.ServiceServiceHandler.
func (s *Service) LookupAddressById(context.Context, *connect.Request[v1.LookupAddressByIdRequest]) (*connect.Response[v1.LookupAddressByIdResponse], error) {
	panic("unimplemented")
}

// LookupUserById implements v1connect.ServiceServiceHandler.
func (s *Service) LookupUserById(context.Context, *connect.Request[v1.LookupUserByIdRequest]) (*connect.Response[v1.LookupUserByIdResponse], error) {
	panic("unimplemented")
}

func (s *Service) QueryMe(ctx context.Context, c *connect.Request[v1.QueryMeRequest]) (*connect.Response[v1.QueryMeResponse], error) {
	response := &v1.QueryMeResponse{
		Me: data.ME,
	}
	return connect.NewResponse(response), nil
}
func (s *Service) Start() error {
	// Create the service handler
	path, handler := v1connect.NewServiceServiceHandler(s)

	// Create the HTTP server
	mux := http.NewServeMux()

	// Mount the service handler
	mux.Handle(path, handler)

	// Create the server with h2c support for gRPC
	server := &http.Server{
		Addr:    ":4011",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	fmt.Printf("Starting gRPC server on port 4011...")

	// Start the server
	return server.ListenAndServe()
}
