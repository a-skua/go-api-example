#!/bin/sh

ADDR='api:80'

# ユーザー追加
RESPONSE=$(curl -s -X 'POST' -d '{"user":{"name":"Bob","password":"12345678"}}' "$ADDR/user")
USER_ID=$(echo $RESPONSE | jq -r '.user.id')
echo "  POST $RESPONSE"

RESPONSE=$(curl -s -X 'GET' "$ADDR/user/$USER_ID")
echo "   GET $RESPONSE"

RESPONSE=$(curl -s -X 'PUT' -d '{"user":{"name":"Alice","password":"12345678"}}' "$ADDR/user/$USER_ID")
echo "   PUT $RESPONSE"

RESPONSE=$(curl -s -X 'DELETE' "$ADDR/user/$USER_ID")
echo "DELETE $RESPONSE"
