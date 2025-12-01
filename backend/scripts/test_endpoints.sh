#!/bin/bash

BASE_URL="http://localhost:8080"
EMAIL="test_$(date +%s)@example.com"
PASSWORD="password123"

echo "1. Testing Health Check..."
curl -s "$BASE_URL/v1/health"
echo -e "\n"

echo "2. Testing Registration..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/users" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"testuser\", \"email\": \"$EMAIL\", \"password\": \"$PASSWORD\"}")
echo "Response: $REGISTER_RESPONSE"
echo -e "\n"

echo "3. Testing Login..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/users/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$EMAIL\", \"password\": \"$PASSWORD\"}")
echo "Response: $LOGIN_RESPONSE"

TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
  echo "Failed to extract token"
  exit 1
fi

echo -e "\nToken: $TOKEN\n"

echo "4. Testing Get Current User (Protected)..."
curl -s -X GET "$BASE_URL/users/me" \
  -H "Authorization: Bearer $TOKEN"
echo -e "\n"
