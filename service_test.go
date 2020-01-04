package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {

	now := time.Now().Unix()
	stamp := now + 86_400

	fmt.Println(now)
	fmt.Println(stamp)

}

func TestToken(t *testing.T) {

	var svc oAuthService

	token, err := svc.Token("password", "admin", "password")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", token)

	if token.AccessToken == "" {
		t.Errorf("AccessToken = %s", token.AccessToken)
	}

	tokenInfo, err := svc.CheckToken(token.AccessToken)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", tokenInfo)

	if !tokenInfo.Active {
		t.Errorf("TokenInfo = %#v", tokenInfo)
	}
}
