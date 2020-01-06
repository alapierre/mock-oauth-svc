// Simple mock oAuth service for microservice integration test
package main

import (
	"github.com/go-kit/kit/auth/basic"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	users "mock-oauth-svr/users"
	"net/http"
	"os"
)

type config struct {
	Client string
	Secret string
}

func main() {

	var client string
	var secret string

	if client = os.Getenv("AUTH_CLIENT"); client == "" {
		client = "client"
		log.Println("Using default client for BASIC-AUTH")
	}

	if secret = os.Getenv("AUTH_SECRET"); secret == "" {
		secret = "secret"
		log.Println("Using default secret for BASIC-AUTH")
	}

	var usvc = users.UserServiceSimple{}
	var svc OAuthService = NewOAuthService(usvc)

	tokenHandler := httptransport.NewServer(
		//makeTokenEndpoint(svc),
		basic.AuthMiddleware("client", "secret", "oAuth client Realm")(makeTokenEndpoint(svc)),
		decodeTokenRequest,
		encodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
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
