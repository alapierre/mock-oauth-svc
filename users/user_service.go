package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Login       string   `json:"login"`
	Password    string   `json:"password"`
	Authorities []string `json:"authorities"`
}

type Client struct {
	Client    string   `json:"client"`
	Secret    string   `json:"secret"`
	Scope     []string `json:"scope"`
	GrantType []string `json:"grant_type"`
}

type ServerConfig struct {
	Clients []Client
	Users   []User
}

type UserServiceFileImpl struct {
	Clients map[string]Client
	Users   map[string]User
}

func NewServiceFromFile(fileName string) (*UserServiceFileImpl, error) {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config ServerConfig

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	return &UserServiceFileImpl{
		Users:   mapUsers(config.Users),
		Clients: mapClients(config.Clients),
	}, nil
}

func (svc *UserServiceFileImpl) AuthenticateUser(login, password string) ([]string, error) {
	user, ok := svc.Users[login+":"+password]
	if !ok {
		return nil, fmt.Errorf("cannot find user %s", user)
	}
	return user.Authorities, nil
}

func mapClients(clients []Client) map[string]Client {

	res := make(map[string]Client)
	for _, i := range clients {
		res[i.Client+":"+i.Secret] = i
	}

	return res
}

func mapUsers(clients []User) map[string]User {

	res := make(map[string]User)
	for _, i := range clients {
		res[i.Login+":"+i.Password] = i
	}

	return res
}
