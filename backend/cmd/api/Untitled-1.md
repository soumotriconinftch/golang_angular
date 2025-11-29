Learn:
* Indexing(Later)
* SQL queries (Constraints, foreign Key - Later)
* Go Validators
* Req next
* Cors allowedHeaders
* Timeout of 5 configured on server level and api takes 5sec


To Do:
* API => Update API to get all users (Pass JWT from login api as authorization header), Move the username check to Service or repo
* UI => Login & Register (Name, username, password)
* API => Login (Create access token JWT) 
* UI => To store JWT after login success 
* API => Create new api to get current logged-in user data (GET) (Pass JWT from login api as authorization header) (Res: Name, isAdmin)
* Backend => Create table for admin users and check against this for above

Done:
* Backend => Move connection URL to env file and env should be ignored by git



1. JWT Generator and Validator
2. Validator
2. Register API changes
    -Create user
        -Validate the payload
            -check if user exists or not
                -if yes return error
                -if no do a DB query
            -Responsd with a JWT token

3. Create login API
    -Get User
        -Validate the payload
            -check if user exists or not
                -if yes return error
                -if no do a DB query
            -Responsd with a JWT token


Chat