package data

import "products-subgraph/graph/model"

// FEATURED_PRODUCTS - Produtos em destaque
var FEATURED_PRODUCTS = []*model.FeaturedProductSection{
	{
		ID:    "featured-1",
		Title: "Medicamentos em Destaque",
		Products: []*model.Product{
			{
				ID:              "prod-1",
				Name:            "Paracetamol 500mg",
				Brand:           "Tylenol",
				QuantityInStock: 150,
				Price:           12.9,
				Gallery:         []string{"paracetamol1.jpg", "paracetamol2.jpg"},
				Description:     "Analgésico e antitérmico para dores e febre",
				Category:        "Medicamentos",
				Available:       true,
			},
			{
				ID:              "prod-2",
				Name:            "Ibuprofeno 400mg",
				Brand:           "Advil",
				QuantityInStock: 200,
				Price:           18.5,
				Gallery:         []string{"ibuprofeno1.jpg"},
				Description:     "Anti-inflamatório para dores e inflamações",
				Category:        "Medicamentos",
				Available:       true,
			},
		},
	},
	{
		ID:    "featured-2",
		Title: "Suplementos Populares",
		Products: []*model.Product{
			{
				ID:              "prod-3",
				Name:            "Vitamina D3 2000UI",
				Brand:           "Addera",
				QuantityInStock: 80,
				Price:           45.9,
				Gallery:         []string{"vitamina-d1.jpg"},
				Description:     "Suplemento de vitamina D para ossos e imunidade",
				Category:        "Suplementos",
				Available:       true,
			},
		},
	},
}

// PRODUTOS_ADICIONAIS - Produtos referenciados por outros subgraphs
var PRODUTOS_ADICIONAIS = []*model.Product{
	{
		ID:              "prod-4",
		Name:            "Dipirona 500mg",
		Brand:           "Novalgina",
		QuantityInStock: 120,
		Price:           8.9,
		Gallery:         []string{"dipirona1.jpg"},
		Description:     "Analgésico e antitérmico",
		Category:        "Medicamentos",
		Available:       true,
	},
	{
		ID:              "prod-5",
		Name:            "Omeprazol 20mg",
		Brand:           "Losec",
		QuantityInStock: 90,
		Price:           25.5,
		Gallery:         []string{"omeprazol1.jpg"},
		Description:     "Protetor gástrico",
		Category:        "Medicamentos",
		Available:       true,
	},
	{
		ID:              "prod-6",
		Name:            "Vacina Antigripal",
		Brand:           "Butantan",
		QuantityInStock: 50,
		Price:           89.9,
		Gallery:         []string{"vacina1.jpg"},
		Description:     "Vacina contra gripe",
		Category:        "Vacinas",
		Available:       true,
	},
}
