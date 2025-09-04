package services

import (
	example_port "standard-struct-golang/modules/frontweb/modules/example/ports"
)

type ExampleService struct {
	repo example_port.ExampleRepositories
}

func NewExampleService(repo example_port.ExampleRepositories) *ExampleService {
	return &ExampleService{repo: repo}
}
