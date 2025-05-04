# Split - Backend API (Go)

Split is a backend service built in Go to manage shared expenses and simplify bill splitting among users. It supports user and group management, transaction tracking, balance calculation, and email notifications. The backend uses PostgreSQL for persistent storage and Redis for session management.

---

## ✨ Features

- User and Group Management
- Expense Tracking and Bill Splitting
- Transaction Recording and History
- Session Management with Redis
- Email Notifications via SMTP
- Environment-specific Configuration (Dev / Prod)

---

## 🛠 Tech Stack

- Go (Golang)
- PostgreSQL
- Redis
- Gorilla Mux (Routing)
- Viper (Config management)
- slog (Structured logging)
- SMTP (Email support)

---

## 📁 Project Structure

```

.
├── db/
│   ├── postgres/             # PostgreSQL repository implementations
│   └── redis/                # Redis session manager
├── pkg/
│   ├── balance/              # Balance calculation logic
│   ├── groupmember/          # Group-member relations
│   ├── groups/               # Group-related logic
│   ├── mail/                 # Email sending utility
│   ├── payment/              # Payment tracking
│   ├── request/              # User requests (e.g., joining groups)
│   ├── sessmanager/          # Session logic
│   ├── transaction/          # Transactions
│   ├── transaction\_split/    # Splitting logic
│   └── users/                # User registration and login
├── server/                   # Main HTTP server logic
└── .split/                   # Development config directory

````

---

## 🧪 Requirements

- Go 1.21 or newer
- PostgreSQL
- Redis
- SMTP-compatible email account

---

## 🔧 Configuration

### Development

Create a file at `.split/dev-split.json` with the following structure:

```json
{
  "mail_id": "your_email@example.com",
  "mail_pass": "your_email_password",
  "app_pass": "your_app_specific_password"
}
````

### Production

Create a file at `/app/prod-split.json` (for Docker/Koyeb) with the same structure.

---

## ▶️ Running the Server

To start the server in development mode:

```
go run main.go dev
```

To start in production mode:

```
go run main.go prod
```

The server will run on:

* `http://localhost:8080` (Development)
* `http://localhost:8194` (Production)

---



## 🧾 Logging

Structured logging is implemented using Go's `slog` package.
Log level is based on the environment:

* **Development**: Debug level
* **Production**: Info level

---

