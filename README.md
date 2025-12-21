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
3. The form validates input requirements locally
4. On submission, credentials are sent to the Go backend
5. Successful login stores the session and redirects to the dashboard

#### Signup
1. Navigate to the signup page
2. Fill in all required fields
3. Real-time validation ensures data integrity
4. On submission, the user is created in the backend database
5. Successful signup prompts the user to login

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

The frontend is fully integrated with the Go backend via `AuthService`.
- **Authentication**: Login and Signup endpoints are connected.
- **Blog Data**: Blog posts are fetched from the backend API.
- **CORS**: Configured to allow requests from the Angular frontend.

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