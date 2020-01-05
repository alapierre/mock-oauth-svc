package main

import (
	"context"
	"encoding/json"
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

type checkTokenRequest struct {
	token string
}

type RestError struct {
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
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

func makeCheckTokenEndpoint(svc OAuthService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(checkTokenRequest)
		tokenInfo, err := svc.CheckToken(req.token)
		if err != nil {
			return RestError{
				Error:            "invalid_token",
				ErrorDescription: "Token was not recognised",
			}, nil
		}
		return tokenInfo, nil
	}
}

func decodeTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return tokenRequest{
		r.URL.Query().Get("username"),
		r.URL.Query().Get("password"),
		r.URL.Query().Get("grant_type"),
	}, nil
}

func decodeCheckTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return checkTokenRequest{
		r.URL.Query().Get("token"),
	}, nil
}

// Encodes response as a JSON
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	headers := w.Header()
	headers.Set("Content-Type", "application/json; charset=utf-8")
	headers.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	headers.Set("X-Content-Type-Options", "nosniff")
	headers.Set("X-XSS-Protection", "1; mode=block")
	headers.Set("Pragma", "no-cache")
	headers.Set("Expires", "0")
	headers.Set("X-Frame-Options", "DENY")
	return json.NewEncoder(w).Encode(response)
}

func MakeError(r *http.Request) RestError {
	return RestError{
		Error:            "invalid_token",
		ErrorDescription: "Token was not recognised",
	}

}
