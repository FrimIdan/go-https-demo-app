#!/bin/sh

swagger generate client -f swagger.yaml -t ./client
swagger generate server -f swagger.yaml -t ./server
