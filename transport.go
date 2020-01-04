package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type TokenInfo struct {
	UserName    string   `json:"user_name,omitempty"`
	Active      bool     `json:"active,omitempty"`
	Exp         int64    `json:"exp,omitempty"`
	ClientId    string   `json:"client_id,omitempty"`
	Scope       []string `json:"scope,omitempty"`
	Authorities []string `json:"authorities,omitempty"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type tokenRequest struct {
	User      string
	Password  string
	GrantType string
}

func makeTokenEndpoint(svc OAuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(tokenRequest)
		token, err := svc.Token(req.GrantType, req.User, req.Password)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
}

func decodeTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return tokenRequest{
		r.URL.Query().Get("username"),
		r.URL.Query().Get("password"),
		r.URL.Query().Get("grant_type"),
	}, nil
}
