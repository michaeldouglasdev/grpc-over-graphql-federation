package data

import v1 "github.com/wundergraph/service/pkg/generated/service/v1"

// SERVICES - Serviços farmacêuticos
var SERVICES = []*v1.Service{
	{
		Id:          "service-1",
		Name:        "Aplicação de Vacinas",
		Description: "Aplicação de vacinas contra COVID-19, gripe e outras doenças",
		Price:       25.9,
		Duration:    30,
		Category:    "Vacinação",
		Available:   true,
		Location:    "Unidade Centro",
	},
	{
		Id:          "service-2",
		Name:        "Teste de Glicemia",
		Description: "Medição de glicose no sangue para controle do diabetes",
		Price:       15.5,
		Duration:    15,
		Category:    "Exames",
		Available:   true,
		Location:    "Unidade Centro",
	},
	{
		Id:          "service-3",
		Name:        "Medição de Pressão",
		Description: "Verificação da pressão arterial",
		Price:       8.9,
		Duration:    10,
		Category:    "Exames",
		Available:   true,
		Location:    "Todas as unidades",
	},
	{
		Id:          "service-4",
		Name:        "Consulta Farmacêutica",
		Description: "Orientação sobre medicamentos e tratamentos",
		Price:       35.0,
		Duration:    45,
		Category:    "Consultoria",
		Available:   true,
		Location:    "Unidade Centro",
	},
	{
		Id:          "service-5",
		Name:        "Aplicação de Injeções",
		Description: "Aplicação de injeções intramusculares e subcutâneas",
		Price:       18.9,
		Duration:    20,
		Category:    "Procedimentos",
		Available:   true,
		Location:    "Unidade Centro",
	},
}
