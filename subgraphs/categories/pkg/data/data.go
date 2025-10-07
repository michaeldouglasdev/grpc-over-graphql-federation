package data

import v1 "github.com/wundergraph/service/pkg/generated/service/v1"

// CATEGORIES - Categorias de produtos
var CATEGORIES = []*v1.Category{
	{
		Id:           "cat-1",
		Name:         "Medicamentos",
		Description:  "Medicamentos controlados e de venda livre",
		Icon:         "💊",
		ProductCount: 1250,
		ServiceCount: 0,
	},
	{
		Id:           "cat-2",
		Name:         "Vacinas",
		Description:  "Vacinas para prevenção de doenças",
		Icon:         "💉",
		ProductCount: 45,
		ServiceCount: 3,
	},
	{
		Id:           "cat-3",
		Name:         "Dermocosméticos",
		Description:  "Produtos para cuidados com a pele",
		Icon:         "🧴",
		ProductCount: 320,
		ServiceCount: 0,
	},
	{
		Id:           "cat-4",
		Name:         "Suplementos",
		Description:  "Vitaminas e suplementos alimentares",
		Icon:         "💊",
		ProductCount: 180,
		ServiceCount: 0,
	},
	{
		Id:           "cat-5",
		Name:         "Cuidados Pessoais",
		Description:  "Produtos de higiene e cuidados pessoais",
		Icon:         "🧼",
		ProductCount: 450,
		ServiceCount: 0,
	},
}
