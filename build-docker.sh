#!/bin/bash

CGO_ENABLED=0 go build -a -installsuffix cgo -o mock-oauth-svr .
docker build -t lapierre/mock-oauth:0.0.1 .
