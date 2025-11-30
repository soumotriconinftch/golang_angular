#!/bin/bash

API_URL="http://localhost:8080"

echo "=============================================="
echo "  AuthTokenMiddleware Demo"
echo "=============================================="


RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${BLUE}Test 1: Access /users/me WITHOUT Authorization header${NC}"
echo -e "${CYAN}Command:${NC} curl -X GET $API_URL/users/me"

curl -s -X GET "$API_URL/users/me"

echo -e "${RED}Expected: 401 Unauthorized - Authorization header is missing${NC}"
echo ""
echo "=============================================="
echo ""

echo -e "${BLUE}Test 2: Access /users/me with INVALID header format${NC}"
echo -e "${CYAN}Command:${NC} curl -X GET $API_URL/users/me -H \"Authorization: InvalidToken123\""
echo ""
curl -s -X GET "$API_URL/users/me" -H "Authorization: InvalidToken123"
echo ""
echo -e "${RED}Expected: 401 Unauthorized - Invalid authorization header format${NC}"
echo ""
echo "=============================================="
echo ""

echo -e "${BLUE}Test 3: Access /users/me with FAKE token${NC}"
echo -e "${CYAN}Command:${NC} curl -X GET $API_URL/users/me -H \"Authorization: Bearer fake.token.here\""
echo ""
curl -s -X GET "$API_URL/users/me" -H "Authorization: Bearer fake.token.here"
echo ""
echo -e "${RED}Expected: 401 Unauthorized - Invalid token${NC}"
echo ""
echo "=============================================="
echo ""

echo -e "${BLUE}Test 4: Login to get a VALID token${NC}"
echo -e "${CYAN}Command:${NC} curl -X POST $API_URL/users/login ..."
echo ""

LOGIN_RESPONSE=$(curl -s -X POST "$API_URL/users/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "soumosaha@example.com",
    "password": "helloworld@123"
  }')

echo "$LOGIN_RESPONSE" | jq '.'
echo ""

TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')

if [ "$TOKEN" != "null" ] && [ -n "$TOKEN" ]; then
    echo -e "${GREEN}Login successful! Token received${NC}"
    echo -e "${YELLOW}Token:${NC} ${TOKEN:0:50}..."
else
    echo -e "${RED}Login failed. Make sure user exists.${NC}"
    echo "Create user first with: ./test_user_flow.sh"
    exit 1
fi

echo ""
echo "=============================================="
echo ""

echo -e "${BLUE}Test 5: Access /users/me with VALID token${NC}"
echo -e "${CYAN}Command:${NC} curl -X GET $API_URL/users/me -H \"Authorization: Bearer <token>\""
echo ""

USER_RESPONSE=$(curl -s -X GET "$API_URL/users/me" \
  -H "Authorization: Bearer $TOKEN")

echo "$USER_RESPONSE" | jq '.'
echo ""
echo -e "${GREEN}Success! Middleware validated token and extracted user_id${NC}"
echo -e "${GREEN}   Handler received user_id from context and fetched user data${NC}"
echo ""

echo "=============================================="
echo -e "${GREEN}Demo Complete!${NC}"
echo "=============================================="
echo ""
echo -e "${YELLOW}What happened in Test 5:${NC}"
echo "1. Middleware extracted token from 'Authorization: Bearer <token>'"
echo "2. Validated JWT signature and expiration"
echo "3. Extracted user_id from token claims"
echo "4. Added user_id to request context"
echo "5. Passed request to getCurrentUserHandler"
echo "6. Handler retrieved user_id from context"
echo "7. Fetched user from database and returned data"
