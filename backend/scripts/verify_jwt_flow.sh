#!/bin/bash

# Port is 8080 based on server logs
BASE_URL="http://localhost:8080"
TS=$(date +%s)
EMAIL="user_${TS}@example.com"
USERNAME="user_${TS}"
PASSWORD="password123"

echo "Using BASE_URL: $BASE_URL"
echo "Creating user: $USERNAME / $EMAIL"

echo -e "\n1. Testing Health Check (Expect 404)..."
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/v1/health")
echo "Health Check Status: $HTTP_CODE"
if [ "$HTTP_CODE" -eq 404 ]; then
    echo "PASS: Health check removed."
else
    echo "FAIL: Health check exists or other error."
fi

echo -e "\n2. Testing Sign Up..."
rm -f cookies.txt
curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/sign-up" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"$USERNAME\", \"email\": \"$EMAIL\", \"password\": \"$PASSWORD\"}" > signup_response.json

cat signup_response.json
echo ""

if grep -q "accessToken" cookies.txt; then
    echo "PASS: Access token cookie found."
else
    echo "FAIL: Access token cookie NOT found."
    cat cookies.txt
fi

echo -e "\n3. Testing Get Current User (Protected)..."
curl -s -c cookies.txt -b cookies.txt "$BASE_URL/user/me" > me_response.json
cat me_response.json
echo ""

if grep -q "$USERNAME" me_response.json; then
    echo "PASS: User details retrieved."
else
    echo "FAIL: Could not retrieve user details."
fi

echo -e "\n4. Testing Sign In..."
# Clear cookies to simulate fresh login
rm -f cookies.txt
curl -s -c cookies.txt -b cookies.txt -X POST "$BASE_URL/user/sign-in" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$EMAIL\", \"password\": \"$PASSWORD\"}" > login_response.json

cat login_response.json
echo ""

if grep -q "accessToken" cookies.txt; then
    echo "PASS: Access token cookie found after login."
else
    echo "FAIL: Access token cookie NOT found after login."
fi

echo -e "\n5. Testing Get Current User after Login..."
curl -s -c cookies.txt -b cookies.txt "$BASE_URL/user/me" > me_login_response.json
cat me_login_response.json
echo ""

if grep -q "$USERNAME" me_login_response.json; then
    echo "PASS: User details retrieved after login."
else
    echo "FAIL: Could not retrieve user details after login."
fi
