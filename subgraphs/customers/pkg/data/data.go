package data

import (
	v1 "github.com/wundergraph/service/pkg/generated/service/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ME - Dados do usuário principal
var ME = &v1.User{
	Id:          "id-1",
	Name:        "Michael",
	Cpf:         "123.456.789-10",
	Email:       "md@md.com.br",
	DateOfBirth: "1970/01/01",
	Avatar:      "avatar.png",
	Address:     ADDRESS,
}

// ADDRESS - Endereço do usuário
var ADDRESS = &v1.Address{
	Id:         "id-address-1",
	Name:       "My house",
	Country:    "United States",
	State:      "ABC",
	City:       "XYZ",
	District:   "ABC",
	Street:     "Street Lorem Ipsum",
	Number:     "01",
	Complement: &wrapperspb.StringValue{Value: "Near of Center Park"},
}
