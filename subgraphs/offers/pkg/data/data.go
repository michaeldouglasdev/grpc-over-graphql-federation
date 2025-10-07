package data

import v1 "github.com/wundergraph/service/pkg/generated/service/v1"

var OFFER_SECTIONS = []*v1.OfferSection{
	{
		Id:    "offer-1",
		Title: "Ofertas da Semana",
		Products: []*v1.Product{
			{
				Id: "prod-2",
			},
			{
				Id: "prod-4",
			},
			{
				Id: "prod-5",
			},
		},
	},
	{
		Id:    "offer-2",
		Title: "Promoções de Vacinas",
		Products: []*v1.Product{
			{
				Id: "prod-6",
			},
		},
	},
}
