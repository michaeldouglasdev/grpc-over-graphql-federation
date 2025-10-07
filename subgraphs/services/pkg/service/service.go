package service

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/wundergraph/service/pkg/data"
	v1 "github.com/wundergraph/service/pkg/generated/service/v1"
	servicev1 "github.com/wundergraph/service/pkg/generated/service/v1/v1connect"
)

type Service struct{}

func (s *Service) LookupServiceById(ctx context.Context, c *connect.Request[v1.LookupServiceByIdRequest]) (*connect.Response[v1.LookupServiceByIdResponse], error) {
	for _, service := range data.SERVICES {
		if service.Id == c.Msg.Keys[0].Id {
			return connect.NewResponse(&v1.LookupServiceByIdResponse{
				Result: []*v1.Service{service},
			}), nil
		}
	}
	return nil, nil
}

func (s *Service) QueryServices(ctx context.Context, c *connect.Request[v1.QueryServicesRequest]) (*connect.Response[v1.QueryServicesResponse], error) {
	response := &v1.QueryServicesResponse{
		Services: data.SERVICES,
	}

	return connect.NewResponse(response), nil
}

func (s *Service) Start() error {
	service := &Service{}
	mux := http.NewServeMux()
	path, handler := servicev1.NewServiceServiceHandler(service)
	mux.Handle(path, handler)
	slog.Info(`Listening on http://localhost:4013`)
	return http.ListenAndServe(
		"localhost:4013",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
