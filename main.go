// Simple mock oAuth service for microservice integration test
package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main() {

	var svc OAuthService = oAuthService{}

	tokenHandler := httptransport.NewServer(
		makeTokenEndpoint(svc),
		decodeTokenRequest,
		encodeResponse,
	)

	checkTokenHandler := httptransport.NewServer(
		makeCheckTokenEndpoint(svc),
		decodeCheckTokenRequest,
		encodeResponse,
	)

	http.Handle("/oauth/token/", tokenHandler)
	http.Handle("/oauth/check_token/", checkTokenHandler)

	log.Println("Server is starting on port 9005")

	panic(http.ListenAndServe(":9005", nil))
}
