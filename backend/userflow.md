# 🔄 Full Stack User & Developer Flow
 
This document maps every user interaction to the specific code running on both the Frontend (Angular) and Backend (Go).
 
---
 
## 1️⃣ Registration Flow
 
### 👤 User Perspective
1. User clicks **"Sign Up"** button on the landing page.
2. Enters `Username`, `Email`, and `Password`.
3. Clicks **"Create Account"**.
4. Sees a success message or is redirected to the dashboard.
 
### 💻 Developer Perspective
 
#### A. Frontend (Angular)
**File:** `register.component.ts` (Illustrative)
```typescript
// 1. Capture user input
const userData = {
  username: this.registerForm.value.username,
  email: this.registerForm.value.email,
  password: this.registerForm.value.password
};
 
// 2. Call Service
this.authService.register(userData).subscribe(response => {
  // 3. Handle Success
  localStorage.setItem('token', response.token); // Save JWT
  this.router.navigate(['/dashboard']);
});
```
 
#### B. Network Request
- **Method:** `POST`
- **URL:** `http://localhost:9000/users/`
- **Body:** `{ "username": "...", "email": "...", "password": "..." }`
 
#### C. Backend (Go)
**File:** `cmd/api/api.go` (Router)
```go
// Route definition
r.Route("/users", func(r chi.Router) {
    r.Post("/", app.createUserHandler) // <--- Maps POST /users/ to this handler
})
```
 
**File:** `cmd/api/users.go` (Handler)
```go
func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Decode JSON
    var payload NewUserPayload
    json.NewDecoder(r.Body).Decode(&payload)
 
    // 2. Validate
    if err := Validate.Struct(payload); err != nil { ... }
 
    // 3. Create User & Hash Password
    user := &store.Users{ ... }
    user.Password.Set(payload.Password) // Bcrypt hashing
 
    // 4. Save to DB
    app.store.Users.Create(r.Context(), user)
 
    // 5. Generate Token
    token, _ := auth.GenerateToken(user.ID)
 
    // 6. Send Response
    json.NewEncoder(w).Encode(map[string]interface{}{
        "user": user,
        "token": token,
    })
}
```
 
---
 
## 2️⃣ Login Flow
 
### 👤 User Perspective
1. User clicks **"Login"**.
2. Enters `Email` and `Password`.
3. Clicks **"Sign In"**.
4. Is redirected to their personal dashboard.
 
### 💻 Developer Perspective
 
#### A. Frontend (Angular)
**File:** `login.component.ts` (Illustrative)
```typescript
// 1. Call Service
this.authService.login(email, password).subscribe(response => {
  // 2. Save Token
  localStorage.setItem('token', response.token);
  // 3. Save User Info
  localStorage.setItem('user', JSON.stringify(response.user));
  // 4. Redirect
  this.router.navigate(['/dashboard']);
});
```
 
#### B. Network Request
- **Method:** `POST`
- **URL:** `http://localhost:9000/users/login`
- **Body:** `{ "email": "...", "password": "..." }`
 
#### C. Backend (Go)
**File:** `cmd/api/api.go` (Router)
```go
r.Post("/login", app.loginHandler) // <--- Maps POST /users/login
```
 
**File:** `cmd/api/auth.go` (Handler)
```go
func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Decode & Validate
    var payload LoginPayload
    json.NewDecoder(r.Body).Decode(&payload)
 
    // 2. Find User by Email
    user, err := app.store.Users.GetByEmail(..., payload.Email)
 
    // 3. Check Password
    if err := user.ComparePassword(payload.Password); err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
 
    // 4. Generate New Token
    token, _ := auth.GenerateToken(user.ID)
 
    // 5. Send Response
    json.NewEncoder(w).Encode(...)
}
```
 
---
 
## 3️⃣ Authenticated Dashboard Flow (Protected Route)
 
### 👤 User Perspective
1. User lands on **Dashboard**.
2. Sees **"Welcome, [Username]"**.
3. Sees their private data.
 
### 💻 Developer Perspective
 
#### A. Frontend (Angular)
**File:** `dashboard.component.ts` (Illustrative)
```typescript
// 1. Get Token from Storage
const token = localStorage.getItem('token');
 
// 2. Prepare Headers
const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
 
// 3. Make Request
this.http.get('http://localhost:9000/users/me', { headers }).subscribe(
  user => {
    this.username = user.username; // Display on screen
  },
  error => {
    // If 401, redirect to login
    this.router.navigate(['/login']);
  }
);
```
 
#### B. Network Request
- **Method:** `GET`
- **URL:** `http://localhost:9000/users/me`
- **Headers:** `Authorization: Bearer <jwt_token>`
- **Body:** *Empty*
 
#### C. Backend (Go)
**File:** `cmd/api/api.go` (Router)
```go
r.Group(func(r chi.Router) {
    r.Use(app.AuthTokenMiddleware) // <--- 1. Middleware runs first
    r.Get("/me", app.getCurrentUserHandler) // <--- 2. Handler runs if middleware passes
})
```
 
**File:** `cmd/api/middleware.go` (Middleware)
```go
func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1. Extract Header
        authHeader := r.Header.Get("Authorization")
        
        // 2. Parse "Bearer <token>"
        tokenString := strings.Split(authHeader, " ")[1]
 
        // 3. Validate Token
        token, _ := auth.ValidateToken(tokenString)
        
        // 4. Extract User ID from Claims
        claims := token.Claims.(jwt.MapClaims)
        userID := int64(claims["user_id"].(float64))
 
        // 5. Inject into Context & Call Next
        ctx := context.WithValue(r.Context(), userIDKey, userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```
 
**File:** `cmd/api/users.go` (Handler)
```go
func (app *application) getCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Get User ID from Context (set by middleware)
    userID := r.Context().Value(userIDKey).(int64)
 
    // 2. Fetch User from DB
    user, _ := app.store.Users.GetByID(..., userID)
 
    // 3. Send JSON
    json.NewEncoder(w).Encode(user)
}
```
 
---
 
## 4️⃣ Logout Flow
 
### 👤 User Perspective
1. User clicks **"Logout"**.
2. Is redirected to Login page.
 
### 💻 Developer Perspective
 
#### A. Frontend (Angular)
**File:** `auth.service.ts`
```typescript
logout() {
  // 1. Remove Token
  localStorage.removeItem('token');
  localStorage.removeItem('user');
  
  // 2. Redirect
  this.router.navigate(['/login']);
}
```
 
#### B. Backend (Go)
*Note: Since JWT is stateless, the backend does not need to do anything for logout. The token simply stops being used by the frontend.*
 
---
 
## 🧩 Summary of Connections
 
| Action | Frontend Responsibility | Backend Responsibility |
| :--- | :--- | :--- |
| **Register** | Collect input, POST to `/users/` | Validate, Hash Password, Insert DB, Return Token |
| **Login** | Collect input, POST to `/users/login` | Verify Password, Return Token |
| **Access Data** | Attach `Authorization: Bearer <token>` | Middleware verifies token, Handler fetches data |
| **Logout** | Delete token from LocalStorage | None (Stateless) |
 
 