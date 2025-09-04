package middleware

import "github.com/dgrijalva/jwt-go"

type UserSessionsClaims struct {
	HashCID       string `json:"hash_cid"`
	ProviderID    string `json:"provider_id"`
	ProviderToken string `json:"provider_token"`
	EncryptedCID  string `json:"encrypted_cid,omitempty"`
	HCode         string `json:"hcode,omitempty"`
	HNameTH       string `json:"hname_th,omitempty"`
	IsAdmin       bool   `json:"is_admin"`
	FirstnameTH   string `json:"firstname_th,omitempty"`
	LastnameTH    string `json:"lastname_th,omitempty"`
	Role          string `json:"role,omitempty"`
	Position      string `json:"position,omitempty"`
	UserAgent     string `json:"user_agent"`
	IsDPH         bool   `json:"is_dph"`
	IsPPH         bool   `json:"is_pph"`
	HeathZone     int    `json:"heath_zone"`
	Province      string `json:"province"`
	District      string `json:"district"`
	Version       string `json:"version"`
}

type Claims struct {
	jwt.StandardClaims
}
