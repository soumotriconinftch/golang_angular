# Blog System

A full-stack blog application built with **Angular** (frontend) and **Go** (backend), featuring user authentication, blog post management, and a modern, responsive UI.

---

## 🚀 Features

### Authentication
- **Login Form**: Reactive form with email and password validation
- **Signup Form**: Reactive form with name, email, password, and password confirmation
- Custom validators for email format and password matching
- Modern, premium UI with gradient backgrounds and smooth animations

### Blog Management
- View all blog posts
- Read individual blog posts
- Dashboard for authenticated users
- Blog post components with custom styling

### UI/UX
- Responsive design for all screen sizes
- Modern glassmorphism effects
- Smooth hover and focus animations
- Real-time form validation with user-friendly error messages
- Premium gradient color schemes

---

## 🛠️ Tech Stack

### Frontend
- **Framework**: Angular 15.2.0
- **Forms**: Reactive Forms with custom validators
- **Styling**: Vanilla CSS with modern design patterns
- **Routing**: Angular Router

### Backend
- **Language**: Go (Golang)
- **Architecture**: RESTful API

---

## 📁 Project Structure

```
golang_angular/
├── frontend/
│   └── frontend/
│       ├── src/
│       │   └── app/
│       │       ├── components/
│       │       │   ├── login-form/      # Login reactive form
│       │       │   ├── signup-form/     # Signup reactive form
│       │       │   └── blogbox/         # Blog post component
│       │       ├── pages/
│       │       │   ├── login/           # Login page
│       │       │   ├── signup/          # Signup page
│       │       │   ├── dashboard/       # User dashboard
│       │       │   ├── blog/            # Blog listing
│       │       │   └── blog-open/       # Individual blog view
│       │       ├── app.module.ts        # Main app module
│       │       └── app-routing.module.ts
│       └── package.json
└── backend/                             # Go backend server
```

---

## 🔧 Setup Instructions

### Prerequisites
- **Node.js** (v14 or higher)
- **npm** (v6 or higher)
- **Go** (v1.16 or higher)
- **Git**

### Frontend Setup

1. **Navigate to the frontend directory**:
   ```bash
   cd frontend/frontend
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Start the development server**:
   ```bash
   npm start
   ```

4. **Access the application**:
   Open your browser and navigate to `http://localhost:4200`

### Backend Setup

1. **Navigate to the backend directory**:
   ```bash
   cd backend
   ```

2. **Install Go dependencies**:
   ```bash
   go mod download
   ```

3. **Run the server**:
   ```bash
   go run main.go
   ```

---

## 📝 Usage

### Authentication

#### Login
1. Navigate to the login page
2. Enter your email and password
3. The form validates:
   - Email is required and must be a valid format
   - Password is required and must be at least 6 characters
4. Submit button is disabled until all validations pass
5. On successful validation, form data is logged to console (ready for backend integration)

#### Signup
1. Navigate to the signup page
2. Fill in all required fields:
   - Full Name (required)
   - Email (required, valid format)
   - Password (required, min 6 characters)
   - Confirm Password (required, must match password)
3. Real-time validation ensures:
   - All fields are filled
   - Email format is correct
   - Passwords match
4. Submit button is disabled until all validations pass
5. On successful validation, form data is logged to console (ready for backend integration)

### Blog Features
- **Dashboard**: View and manage your blog posts
- **Blog Listing**: Browse all available blog posts
- **Read Blog**: Click on any blog post to read the full content

---

## 🎨 Form Validation Details

### Login Form Validators
| Field    | Validation Rules                    | Error Messages                                      |
|----------|-------------------------------------|-----------------------------------------------------|
| Email    | Required, Valid email format        | "Email is required", "Please enter a valid email"   |
| Password | Required, Minimum 6 characters      | "Password is required", "Password must be at least 6 characters" |

### Signup Form Validators
| Field            | Validation Rules                    | Error Messages                                      |
|------------------|-------------------------------------|-----------------------------------------------------|
| Name             | Required                            | "Name is required"                                  |
| Email            | Required, Valid email format        | "Email is required", "Please enter a valid email"   |
| Password         | Required, Minimum 6 characters      | "Password is required", "Password must be at least 6 characters" |
| Confirm Password | Required, Must match password       | "Please confirm your password", "Passwords do not match" |

---

## 🔌 Backend Integration

The forms are currently set up to log data to the console. To integrate with the backend:

1. **Create an authentication service**:
   ```typescript
   // auth.service.ts
   import { Injectable } from '@angular/core';
   import { HttpClient } from '@angular/common/http';
   
   @Injectable({
     providedIn: 'root'
   })
   export class AuthService {
     private apiUrl = 'http://localhost:8080/api';
     
     constructor(private http: HttpClient) {}
     
     login(credentials: any) {
       return this.http.post(`${this.apiUrl}/login`, credentials);
     }
     
     signup(userData: any) {
       return this.http.post(`${this.apiUrl}/signup`, userData);
     }
   }
   ```

2. **Update form components** to use the service instead of console.log

---

## 🚧 Development

### Available Scripts

#### Frontend
- `npm start` - Start development server
- `npm run build` - Build for production
- `npm test` - Run unit tests
- `npm run watch` - Build in watch mode

#### Backend
- `go run main.go` - Start the Go server
- `go test ./...` - Run tests
- `go build` - Build the application

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

## 👨‍💻 Author

Built with ❤️ using Angular and Go

---

## 📞 Support

For issues or questions, please open an issue in the repository.