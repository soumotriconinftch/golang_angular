#!/bin/bash

API_URL="http://localhost:8080"
USERNAME="soumo2"
EMAIL="soumosaha@example.com"
PASSWORD="helloworld@123"

echo "=============================================="
echo "  USER REGISTRATION AND LOGIN TEST SCRIPT"
echo "=============================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}========== STEP 1: CREATE NEW USER ==========${NC}"
echo -e "${YELLOW}Creating user with:${NC}"
echo "  Username: $USERNAME"
echo "  Email: $EMAIL"
echo "  Password: $PASSWORD"
echo ""

CREATE_RESPONSE=$(curl -s -w "\n%{http_code}" -X POST "$API_URL/users/" \
  -H "Content-Type: application/json" \
  -d "{
    \"username\": \"$USERNAME\",
    \"email\": \"$EMAIL\",
    \"password\": \"$PASSWORD\"
  }")

HTTP_CODE=$(echo "$CREATE_RESPONSE" | tail -n1)
RESPONSE_BODY=$(echo "$CREATE_RESPONSE" | sed '$d')

echo -e "${YELLOW}Response Status Code:${NC} $HTTP_CODE"
echo -e "${YELLOW}Response Body:${NC}"
echo "$RESPONSE_BODY" | jq '.' 2>/dev/null || echo "$RESPONSE_BODY"
echo ""

if [ "$HTTP_CODE" -eq 201 ]; then
    echo -e "${GREEN}User created successfully!${NC}"
    CREATE_TOKEN=$(echo "$RESPONSE_BODY" | jq -r '.token' 2>/dev/null)
    USER_ID=$(echo "$RESPONSE_BODY" | jq -r '.user.id' 2>/dev/null)
    echo -e "${YELLOW}User ID:${NC} $USER_ID"
    echo -e "${YELLOW}Token (from registration):${NC} $CREATE_TOKEN"
else
    echo -e "${RED}User creation failed!${NC}"
    if [ "$HTTP_CODE" -eq 500 ]; then
        echo -e "${RED}Note: User might already exist. Proceeding with login test...${NC}"
    fi
fi

echo ""
echo "=============================================="
echo ""

echo -e "${BLUE}========== STEP 2: LOGIN WITH USER ==========${NC}"
echo -e "${YELLOW}Logging in with:${NC}"
echo "  Email: $EMAIL"
echo "  Password: $PASSWORD"
echo ""

LOGIN_RESPONSE=$(curl -s -w "\n%{http_code}" -X POST "$API_URL/users/login" \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$EMAIL\",
    \"password\": \"$PASSWORD\"
  }")

HTTP_CODE=$(echo "$LOGIN_RESPONSE" | tail -n1)
RESPONSE_BODY=$(echo "$LOGIN_RESPONSE" | sed '$d')

echo -e "${YELLOW}Response Status Code:${NC} $HTTP_CODE"
echo -e "${YELLOW}Response Body:${NC}"
echo "$RESPONSE_BODY" | jq '.' 2>/dev/null || echo "$RESPONSE_BODY"
echo ""

if [ "$HTTP_CODE" -eq 200 ]; then
    echo -e "${GREEN}Login successful!${NC}"
    LOGIN_TOKEN=$(echo "$RESPONSE_BODY" | jq -r '.token' 2>/dev/null)
    LOGGED_IN_USER=$(echo "$RESPONSE_BODY" | jq -r '.user.username' 2>/dev/null)
    LOGGED_IN_EMAIL=$(echo "$RESPONSE_BODY" | jq -r '.user.email' 2>/dev/null)
    echo -e "${YELLOW}Logged in as:${NC} $LOGGED_IN_USER ($LOGGED_IN_EMAIL)"
    echo -e "${YELLOW}Token (from login):${NC} $LOGIN_TOKEN"
else
    echo -e "${RED}Login failed!${NC}"
    exit 1
fi

echo ""
echo "=============================================="
echo ""

echo -e "${BLUE}========== STEP 3: TEST AUTHENTICATED ENDPOINT ==========${NC}"
echo -e "${YELLOW}Fetching current user info using token...${NC}"
echo ""

ME_RESPONSE=$(curl -s -w "\n%{http_code}" -X GET "$API_URL/users/me" \
  -H "Authorization: Bearer $LOGIN_TOKEN")

HTTP_CODE=$(echo "$ME_RESPONSE" | tail -n1)
RESPONSE_BODY=$(echo "$ME_RESPONSE" | sed '$d')

echo -e "${YELLOW}Response Status Code:${NC} $HTTP_CODE"
echo -e "${YELLOW}Response Body:${NC}"
echo "$RESPONSE_BODY" | jq '.' 2>/dev/null || echo "$RESPONSE_BODY"
echo ""

if [ "$HTTP_CODE" -eq 200 ]; then
    echo -e "${GREEN}Authenticated request successful!${NC}"
    CURRENT_USER=$(echo "$RESPONSE_BODY" | jq -r '.username' 2>/dev/null)
    CURRENT_EMAIL=$(echo "$RESPONSE_BODY" | jq -r '.email' 2>/dev/null)
    echo -e "${YELLOW}Current User:${NC} $CURRENT_USER"
    echo -e "${YELLOW}Email:${NC} $CURRENT_EMAIL"
else
    echo -e "${RED}Authenticated request failed!${NC}"
fi

echo ""
echo "=============================================="
echo -e "${GREEN}TEST COMPLETED!${NC}"
echo "=============================================="
echo ""
echo -e "${YELLOW}Check the server logs for detailed step-by-step execution${NC}"
