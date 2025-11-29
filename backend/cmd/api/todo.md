# **Step by step approach:**
(
    BACKEND: Golang
    FRONTEND: Angular
)

1. BACKEND: JWT Generator("*github.com/golang-jwt/jwt/v5*") and Validator(Go validator)
    - Define signing method
    - Define secret management
    - Implement token creation
    - Implement token parsing
    - Implement middleware extraction from Authorization header

2. BACKEND: Rewrite the user table with Required information:
    - Username
    - email
    - Password

3. BACKEND: Register API
    - Validate payload(validator func)
    - Check if user exists
    - If exists return conflict
    - If not, insert user
    - Return success response
    - Generate JWT

3. BACKEND: Login API
    - Validate payload
    - Fetch user
    - If user missing return unauthorized
    - If user present verify password
    - Generate JWT
    - Return token

4. BACKEND: Admin Table Creation
    - Define schema
    - (Add isAdmin column) or separate admin table
    - Populate initial admin entry
    - Create repo method to check admin status

5. BACKEND: Current User API (GET)
    - Extract JWT from header
    - Validate token
    - Fetch user from repo
    - Fetch admin status
    - Respond with name and isAdmin

6. UI: Login
    - Create form
    - POST credentials to login API
    - Store JWT on success

7. UI: Register
    - Create form
    - POST data to register API
    - Handle success and JWT if returned

8. UI: Token Storage and Attach
    - Persist JWT in secure storage
    - Attach token to API calls

9. API: Get All Users
    - Require JWT
    - Validate token in middleware
    - Move username validation to service
    - Return user list

10. HOME

    - Login is successful
        - Redirect "GET users/id/content"
(Might need modification as we dive deep into each problem)
(CODE optimisation and testing to be done alongside or after feedback from Logesh.)