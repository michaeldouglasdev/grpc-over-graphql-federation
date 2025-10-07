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

func (s *Service) LookupOfferSectionById(ctx context.Context, c *connect.Request[v1.LookupOfferSectionByIdRequest]) (*connect.Response[v1.LookupOfferSectionByIdResponse], error) {
	panic("unimplemented")
}

func (s *Service) QueryOfferSections(ctx context.Context, c *connect.Request[v1.QueryOfferSectionsRequest]) (*connect.Response[v1.QueryOfferSectionsResponse], error) {
	response := &v1.QueryOfferSectionsResponse{
		OfferSections: data.OFFER_SECTIONS,
	}
	return connect.NewResponse(response), nil
}

func (s *Service) Start() error {
	service := &Service{}
	mux := http.NewServeMux()
	path, handler := servicev1.NewServiceServiceHandler(service)
	mux.Handle(path, handler)
	slog.Info(`Listening on http://localhost:4014`)
	return http.ListenAndServe(
		"localhost:4014",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
