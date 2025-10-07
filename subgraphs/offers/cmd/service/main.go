package main

import (
	"github.com/wundergraph/service/pkg/service"
)

func main() {
	svc := &service.Service{}
	if err := svc.Start(); err != nil {
		panic(err)
	}
}
