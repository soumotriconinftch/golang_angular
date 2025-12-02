#!/bin/bash

BASE_URL="http://localhost:8080"
TS=$(date +%s)
EMAIL="get_content_user_${TS}@example.com"
USERNAME="get_content_user_${TS}"
PASSWORD="password123"

echo "Testing Get Content Flow"
echo "========================"

echo -e "\n1. Sign Up..."
rm -f cookies.txt
curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/sign-up" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"$USERNAME\", \"email\": \"$EMAIL\", \"password\": \"$PASSWORD\"}" > /dev/null

echo "User created and logged in"

echo -e "\n2. Create Content 1..."
CONTENT1=$(curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/me/content/" \
  -H "Content-Type: application/json" \
  -d '{"title": "Post 1", "body": "Body 1"}')
echo $CONTENT1 | jq '.'
ID1=$(echo $CONTENT1 | jq -r '.id')

echo -e "\n3. Create Content 2..."
CONTENT2=$(curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/me/content/" \
  -H "Content-Type: application/json" \
  -d '{"title": "Post 2", "body": "Body 2"}')
echo $CONTENT2 | jq '.'
ID2=$(echo $CONTENT2 | jq -r '.id')

echo -e "\n4. Get All Content..."
curl -s -c cookies.txt -b cookies.txt "$BASE_URL/user/me/content/" | jq '.'

echo -e "\n5. Get Content 1 by ID ($ID1)..."
curl -s -c cookies.txt -b cookies.txt "$BASE_URL/user/me/content/$ID1" | jq '.'

echo -e "\n6. Get Content 2 by ID ($ID2)..."
curl -s -c cookies.txt -b cookies.txt "$BASE_URL/user/me/content/$ID2" | jq '.'

echo -e "\nGet Content test completed!"
