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

func (s *Service) QueryCategories(ctx context.Context, c *connect.Request[v1.QueryCategoriesRequest]) (*connect.Response[v1.QueryCategoriesResponse], error) {
	response := &v1.QueryCategoriesResponse{
		Categories: []*v1.Category{
			{Id: "1", Name: "Category 1"},
		},
	}
	return connect.NewResponse(response), nil
}

func (s *Service) LookupCategoryById(ctx context.Context, c *connect.Request[v1.LookupCategoryByIdRequest]) (*connect.Response[v1.LookupCategoryByIdResponse], error) {
	panic("unimplemented")
	/*response := &v1.LookupCategoryByIdResponse{
		Result: &v1.Category{Id: "1", Name: "Category 1"},
	}
	return connect.NewResponse(response), nil*/
}

func (s *Service) Start() error {
	service := &Service{}
	mux := http.NewServeMux()
	path, handler := servicev1.NewServiceServiceHandler(service)
	mux.Handle(path, handler)
	slog.Info(`Listening on http://localhost:4012`)
	return http.ListenAndServe(
		"localhost:4012",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
