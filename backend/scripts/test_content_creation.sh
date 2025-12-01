#!/bin/bash

BASE_URL="http://localhost:8080"
TS=$(date +%s)
EMAIL="content_user_${TS}@example.com"
USERNAME="content_user_${TS}"
PASSWORD="password123"

echo "Testing Content Creation Flow"
echo "=============================="

echo -e "\n1. Sign Up..."
rm -f cookies.txt
curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/sign-up" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"$USERNAME\", \"email\": \"$EMAIL\", \"password\": \"$PASSWORD\"}" > /dev/null

echo "User created and logged in"

echo -e "\n2. Create Content..."
curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/me/content/" \
  -H "Content-Type: application/json" \
  -d '{"title": "My First Blog Post", "body": "This is the content of my first blog post. It contains some interesting information."}' | jq '.'

echo -e "\n3. Create Another Content..."
curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/me/content/" \
  -H "Content-Type: application/json" \
  -d '{"title": "Second Post", "body": "Another blog post with different content."}' | jq '.'

echo -e "\nContent creation test completed!"
