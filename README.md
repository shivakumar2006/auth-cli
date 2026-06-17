# Auth CLI

A secure, containerized Command-Line Authentication System built with Go, PostgreSQL, Docker, Cobra CLI, and TOTP-based Multi-Factor Authentication (Google Authenticator compatible).

---

## Features

### Authentication

- User Registration
- User Login
- Secure Password Hashing using bcrypt
- Account Lockout after multiple failed login attempts
- Last Login Tracking

### Multi-Factor Authentication (MFA)

- Optional TOTP-based MFA
- Google Authenticator Compatible
- QR Code Generation
- Enable / Disable MFA

### Session Management

- Session Creation
- Session Validation
- Session Expiration
- Session Logout

### Interactive CLI

- Command History
- Tab Completion
- Help Command
- Clear Success/Error Feedback

### Database

- PostgreSQL
- Persistent Storage
- SQL Migrations

### Containerization

- Docker
- Docker Compose
- Persistent PostgreSQL Volume

---

# Tech Stack

| Component          | Technology     |
| ------------------ | -------------- |
| Language           | Go             |
| Database           | PostgreSQL     |
| CLI Framework      | Cobra          |
| Interactive Prompt | go-prompt      |
| Password Hashing   | bcrypt         |
| MFA                | TOTP           |
| QR Code Generation | go-qrcode      |
| Database Driver    | pgx            |
| Migrations         | golang-migrate |
| Containers         | Docker         |

---

# Project Structure

```text
auth-cli/
│
├── cmd/
│   ├── root.go
│   ├── register.go
│   ├── login.go
│   ├── logout.go
│   ├── whoami.go
│   ├── enable2fa.go
│   ├── disable2fa.go
│   └── help.go
│
├── internal/
│
│   ├── auth/
│   │   ├── models.go
│   │   ├── password.go
│   │   ├── service.go
│   │   └── database.go
│   │
│   ├── session/
│   │   ├── models.go
│   │   ├── manager.go
│   │   └── database.go
│   │
│   ├── totp/
│   │   └── totp.go
│   │
│   ├── cli/
│   │   ├── context.go
│   │   ├── completer.go
│   │   └── shell.go
│   │
│   └── database/
│       ├── postgres.go
│       └── migrations.go
│
├── migrations/
│   ├── 001_create_users.up.sql
│   ├── 001_create_users.down.sql
│   ├── 002_create_sessions.up.sql
│   └── 002_create_sessions.down.sql
│
├── qrcodes/
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

---

# Database Schema

## Users Table

| Column          | Type      |
| --------------- | --------- |
| id              | BIGSERIAL |
| username        | VARCHAR   |
| password_hash   | TEXT      |
| mfa_enabled     | BOOLEAN   |
| mfa_secret      | TEXT      |
| failed_attempts | INTEGER   |
| locked_until    | TIMESTAMP |
| created_at      | TIMESTAMP |
| last_login      | TIMESTAMP |

---

## Sessions Table

| Column     | Type      |
| ---------- | --------- |
| id         | BIGSERIAL |
| user_id    | BIGINT    |
| token      | UUID      |
| created_at | TIMESTAMP |
| expires_at | TIMESTAMP |

---

# Security Features

## Password Security

Passwords are never stored in plaintext.

Passwords are hashed using bcrypt before being stored in the database.

```text
Password
↓
bcrypt
↓
Database
```

---

## Account Lockout

After 5 consecutive failed login attempts:

```text
Account Locked
↓
15 Minutes
↓
Login Disabled
```

This helps protect against brute-force attacks.

---

## Multi-Factor Authentication

Uses Time-Based One-Time Passwords (TOTP).

Compatible with:

- Google Authenticator
- Microsoft Authenticator
- Authy

Flow:

```text
Enable MFA
↓
Generate Secret
↓
Generate QR Code
↓
Scan QR
↓
Enter OTP
↓
Verify
↓
MFA Enabled
```

---

## Session Security

A unique UUID session token is generated after login.

```text
Login
↓
UUID Token
↓
Database Session
↓
Expiration
```

Expired sessions are automatically invalidated.

---

# Available Commands

## Before Login

### register

Create a new user account.

```bash
register
```

### login

Login using username and password.

```bash
login
```

### help

Show available commands.

```bash
help
```

### exit

Exit the application.

```bash
exit
```

---

## After Login

### whoami

Display current user information.

```bash
whoami
```

Shows:

- Username
- Registration Date
- MFA Status
- Session Expiration
- Last Login

---

### enable-2fa

Enable Multi-Factor Authentication.

```bash
enable-2fa
```

---

### disable-2fa

Disable Multi-Factor Authentication.

```bash
disable-2fa
```

---

### logout

Terminate current session.

```bash
logout
```

---

# Running Locally

## Clone Repository

```bash
git clone <repository-url>

cd auth-cli
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Environment Variables

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=authdb
```

---

## Run Application

```bash
go run .
```

---

# Running with Docker

Build and start containers:

```bash
docker compose up --build
```

Stop containers:

```bash
docker compose down
```

---

# Example Workflow

## Register

```text
auth-cli > register

Username: shiva
Password: ********

✓ User registered successfully
```

---

## Login

```text
auth-cli > login

Username: shiva
Password: ********

✓ Login successful
```

---

## Enable MFA

```text
auth-cli > enable-2fa

QR Code: qrcodes/shiva.png

Enter OTP:
```

---

## View User

```text
auth-cli > whoami
```

---

## Logout

```text
auth-cli > logout

✓ Logged out successfully
```

---

# Future Improvements

- Refresh Tokens
- Role-Based Access Control (RBAC)
- Password Reset Flow
- Email Verification
- Audit Logging
- Session Cleanup Scheduler
- Unit Tests
- CI/CD Pipeline

---

# Author

Shiva

Backend Developer | Go Developer

Built as part of the Containerized CLI Login System Assignment.
