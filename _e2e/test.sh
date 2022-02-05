#!/bin/sh

ADDR='localhost:8080'

curl -X 'POST' -d '{"user":{"name":"Bob","password":"12345678"}}' "$ADDR/user"
curl -X 'GET' -H 'X-User-Id: 1' "$ADDR/user/1"
