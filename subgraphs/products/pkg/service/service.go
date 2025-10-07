package service

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	v1 "github.com/wundergraph/service/pkg/generated/service/v1"
	servicev1 "github.com/wundergraph/service/pkg/generated/service/v1/v1connect"
)

type Service struct{}

func (s *Service) LookupFeaturedProductSectionById(ctx context.Context, c *connect.Request[v1.LookupFeaturedProductSectionByIdRequest]) (*connect.Response[v1.LookupFeaturedProductSectionByIdResponse], error) {
	panic("unimplemented")
}

func (s *Service) LookupProductById(ctx context.Context, c *connect.Request[v1.LookupProductByIdRequest]) (*connect.Response[v1.LookupProductByIdResponse], error) {
	panic("unimplemented")
}

func (s *Service) QueryFeaturedProducts(ctx context.Context, c *connect.Request[v1.QueryFeaturedProductsRequest]) (*connect.Response[v1.QueryFeaturedProductsResponse], error) {
	response := &v1.QueryFeaturedProductsResponse{
		FeaturedProducts: []*v1.FeaturedProductSection{
			{Id: "1", Title: "Featured Product Section 1"},
		},
	}
	return connect.NewResponse(response), nil
}

func (s *Service) Start() error {
	service := &Service{}
	mux := http.NewServeMux()
	path, handler := servicev1.NewServiceServiceHandler(service)
	mux.Handle(path, handler)
	slog.Info(`Listening on http://localhost:4015`)
	return http.ListenAndServe(
		"localhost:4015",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
