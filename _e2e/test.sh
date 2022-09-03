#!/bin/sh

ADDR='api:80'

# ユーザー
echo "[USER]"
URI="$ADDR/user"
echo "\tPOST $URI"
RESPONSE=$(curl -s -X 'POST' -d '{"user":{"name":"Bob","password":"12345678"}}' "$URI")
echo $RESPONSE | jq -Cc
if [ $? -ne 0 ] || [ "$(echo $RESPONSE | jq -r '.error')" != "null" ]; then exit 1; fi

USER_ID=$(echo $RESPONSE | jq -r '.user.id')
if [ $USER_ID = "null" ]; then exit 1; fi

URI="$ADDR/user/$USER_ID"
echo "\tGET $URI"
RESPONSE=$(curl -s -X 'GET' "$URI")
echo $RESPONSE | jq -Cc
if [ $? -ne 0 ] || [ "$(echo $RESPONSE | jq -r '.error')" != "null" ]; then exit 1; fi

URI=$ADDR/user/$USER_ID
RESPONSE=$(curl -s -X 'PUT' -d '{"user":{"name":"Alice","password":"12345678"}}' "$URI")
echo "\tPUT $URI"
echo $RESPONSE | jq -Cc
if [ $? -ne 0 ] || [ "$(echo $RESPONSE | jq -r '.error')" != "null" ]; then exit 1; fi

URI=$ADDR/user/$USER_ID
RESPONSE=$(curl -s -X 'DELETE' "$URI")
echo "\tDELETE $URI"
echo $RESPONSE | jq -Cc
if [ $? -ne 0 ] || [ "$(echo $RESPONSE | jq -r '.error')" != "null" ]; then exit 1; fi

# 企業
echo "[COMPANY]"
URI=$ADDR/company
echo "\tPOST $URI"
RESPONSE=$(curl -s -X 'POST' -d '{"company":{"name":"GREATE COMPANY","owner_id":1}}' "$URI")
echo $RESPONSE | jq -Cc
if [ $? -ne 0 ] || [ "$(echo $RESPONSE | jq -r '.error')" != "null" ]; then exit 1; fi

COMPANY_ID=$(echo $RESPONSE | jq -r '.company.id')
if [ $COMPANY_ID = "null" ]; then exit 1; fi

URI=$ADDR/company/$COMPANY_ID
echo "\tGET $URI"
RESPONSE=$(curl -s -X 'GET' "$URI")
echo $RESPONSE | jq -Cc
if [ $? -ne 0 ] || [ "$(echo $RESPONSE | jq -r '.error')" != "null" ]; then exit 1; fi

URI=$ADDR/company/$COMPANY_ID
echo "\tPUT $URI"
RESPONSE=$(curl -s -X 'PUT' -d '{"company":{"name":"greate company","owner_id":2}}' "$URI")
echo $RESPONSE | jq -Cc
if [ $? -ne 0 ] || [ "$(echo $RESPONSE | jq -r '.error')" != "null" ]; then exit 1; fi

URI=$ADDR/company/$COMPANY_ID
echo "DELETE $URI"
RESPONSE=$(curl -s -X 'DELETE' "$URI")
echo $RESPONSE | jq -Cc
if [ $? -ne 0 ] || [ "$(echo $RESPONSE | jq -r '.error')" != "null" ]; then exit 1; fi
