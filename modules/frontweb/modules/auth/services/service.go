package services

import (
	auth_port "standard-struct-golang/modules/frontweb/modules/auth/ports"
)

type AuthService struct {
	repo auth_port.AuthRepositories
}

func NewAuthService(repo auth_port.AuthRepositories) *AuthService {
	return &AuthService{repo: repo}
}
