package handler

import (
	auth_port "standard-struct-golang/modules/frontweb/modules/auth/ports"
)

type AuthHandler struct {
	svc auth_port.AuthService
}

func NewAuthHandler(svc auth_port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}
