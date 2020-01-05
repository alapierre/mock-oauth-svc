package main

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"github.com/allegro/bigcache"
	"log"
	"time"
)

var cache *bigcache.BigCache

func init() {
	var err error
	cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(24 * time.Hour))

	if err != nil {
		log.Fatal("Problem with cache", err)
	}
}

type OAuthService interface {
	Token(grantType, username, password string) (*Token, error)
	CheckToken(token string) (*TokenInfo, error)
}

type oAuthService struct{}

func (oAuthService) Token(grantType, username, password string) (*Token, error) {

	token, err := GenerateRandomString(16)

	if err != nil {
		log.Fatal("error generating random string:", err)
	}

	stamp := time.Now().Unix() + 86_400

	cache.Set(token, SerializeStruct(TokenInfo{
		UserName:    username,
		Active:      true,
		Exp:         stamp,
		ClientId:    "",
		Scope:       []string{"api"},
		Authorities: []string{"ROLE_ADMIN", "ROLE_SUPERUSER"},
	}))

	return &Token{
		AccessToken:  token,
		TokenType:    "Bearer",
		RefreshToken: "",
		ExpiresIn:    stamp,
	}, nil
}

func (oAuthService) CheckToken(token string) (*TokenInfo, error) {
	entry, err := cache.Get(token)
	if err != nil {
		return nil, err
	}

	tokenInfo, err := deserializeTokenInfo(entry)
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}

func SerializeStruct(obj interface{}) []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)

	err := enc.Encode(obj)

	if err != nil {
		log.Fatal("encode error:", err)
	}

	return buff.Bytes()
}

func deserializeTokenInfo(source []byte) (*TokenInfo, error) {

	bufor := bytes.NewReader(source)
	dec := gob.NewDecoder(bufor)
	res := &TokenInfo{}
	err := dec.Decode(res)

	return res, err
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, hex encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	randomBytes, err := GenerateRandomBytes(s)
	res := hex.EncodeToString(randomBytes)
	return res[:8] + "-" + res[8:12] + "-" + res[12:16] + "-" + res[16:20] + "-" + res[20:24] + "-" + res[24:], err
}
