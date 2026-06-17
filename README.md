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
- Enable MFA
- Verify MFA
- Disable MFA

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

## Demo Video

```text
https://your-demo-video-link
```

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
│   ├── verify2fa.go
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
├── data/
│   └── qrcodes/
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
Generate OTP
↓
Verify OTP
↓
MFA Enabled
```

---

## Session Security

A unique UUID session token is generated after login.

```text
Login
↓
UUID Session Token
↓
Database Session
↓
Expiration
```

Expired sessions are automatically invalidated.

---

# Available Commands

## Before Login

### Register User

```bash
register <username> <password>
```

Example:

```bash
register shiva Shiva123
```

---

### Login User

```bash
login <username> <password>
```

Example:

```bash
login shiva Shiva123
```

---

### Show Help

```bash
help
```

---

### Exit Application

```bash
exit
```

---

## After Login

### Show Current User

```bash
whoami
```

Displays:

- Username
- Registration Date
- MFA Status
- Session Expiration Time
- Last Login Time

---

### Enable MFA

```bash
enable-2fa
```

Generates:

- TOTP Secret
- QR Code

---

### Verify MFA

```bash
verify-2fa <otp>
```

Example:

```bash
verify-2fa 123456
```

---

### Disable MFA

```bash
disable-2fa yes
```

---

### Logout

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

## Configure Environment Variables

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

## Build Containers

```bash
docker compose build
```

---

## Start Services

```bash
docker compose up -d
```

This starts:

- PostgreSQL Container
- Auth CLI Container

---

## Launch Interactive CLI

Open a shell inside the running CLI container:

```bash
docker exec -it auth-cli-auth-cli-1 sh
```

Then start the application:

```bash
./auth-cli
```

Expected output:

```text
===================================
      Auth CLI Login System
===================================
Type 'help' to see commands

auth-cli >
```

---

## Stop Services

```bash
docker compose down
```

---

# QR Code Generation

When MFA is enabled, a QR code is generated inside the container.

Example:

```text
auth-cli > enable-2fa

===================================
2FA Setup
===================================

QR Code : data/qrcodes/shiva.png
Secret  : XXXXXXXXXXXXXXXXX

Next step:
verify-2fa <otp>
```

---

## Copy QR Code to Host Machine

```bash
docker cp auth-cli-auth-cli-1:/app/data/qrcodes/shiva.png .
```

This copies the generated QR code to your current directory.

---

## Open QR Code

### macOS

```bash
open shiva.png
```

### Linux

```bash
xdg-open shiva.png
```

### Windows PowerShell

```powershell
start shiva.png
```

Scan the QR code using:

- Google Authenticator
- Microsoft Authenticator
- Authy

Then verify MFA:

```bash
verify-2fa 123456
```

---

# Example Workflow

## Register

```text
auth-cli > register shiva Shiva123

✓ User registered successfully
```

---

## Login

```text
auth-cli > login shiva Shiva123

✓ Login successful
```

---

## Enable MFA

```text
auth-cli > enable-2fa

===================================
2FA Setup
===================================

QR Code : data/qrcodes/shiva.png
Secret  : XXXXXXXXXXXXXXXXX

Next step:
verify-2fa <otp>
```

---

## Verify MFA

```text
auth-cli > verify-2fa 123456

✓ MFA enabled successfully
```

---

## View User

```text
auth-cli > whoami
```

---

## Disable MFA

```text
auth-cli > disable-2fa yes

✓ MFA disabled successfully
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

**Shiva**

Backend Developer | Go Developer

Built as part of the **Containerized CLI Login System Assignment**.
