package users

import (
	"fmt"
	"testing"
)

func TestNewServiceFromFile(t *testing.T) {

	userService, err := NewServiceFromFile("users.json")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(userService.Users)
	fmt.Println(userService.Clients)

	auth, err := userService.AuthenticateUser("admin", "admin")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(auth)
}
