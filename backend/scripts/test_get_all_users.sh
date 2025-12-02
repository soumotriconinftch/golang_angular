#!/bin/bash

BASE_URL="http://localhost:8080"
TS=$(date +%s)
EMAIL="get_all_users_${TS}@example.com"
USERNAME="get_all_users_${TS}"
PASSWORD="password123"

echo "Testing Get All Users Flow"
echo "========================"

echo -e "\n1. Sign Up..."
rm -f cookies.txt
curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/sign-up" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"$USERNAME\", \"email\": \"$EMAIL\", \"password\": \"$PASSWORD\"}" > /dev/null

echo "User created and logged in"

echo -e "\n2. Get All Users..."
curl -s -c cookies.txt -b cookies.txt "$BASE_URL/user/all" | jq '.'

echo -e "\nGet All Users test completed!"
