// Simple mock oAuth service for microservice integration test
package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func main() {

	var svc OAuthService = oAuthService{}

	tokenHandler := httptransport.NewServer(
		makeTokenEndpoint(svc),
		decodeTokenRequest,
		encodeResponse,
	)

	http.Handle("/oauth/token/", tokenHandler)

	panic(http.ListenAndServe(":9005", nil))
}
