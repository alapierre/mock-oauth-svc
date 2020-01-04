package main

import (
	"github.com/go-kit/kit/log"
	"net/http"
)

type loggingMiddleware struct {
	logger log.Logger
	next   http.Handler
}
