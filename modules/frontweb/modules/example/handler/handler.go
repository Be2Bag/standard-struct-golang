package handler

import (
	example_port "standard-struct-golang/modules/frontweb/modules/example/ports"
)

type ExampleHandler struct {
	svc example_port.ExampleService
}

func NewExampleHandler(svc example_port.ExampleService) *ExampleHandler {
	return &ExampleHandler{
		svc: svc,
	}
}
