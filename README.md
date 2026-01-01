# Student Management System

A full-stack CRUD application with Go backend and React frontend.

## Quick Start (2 minutes)

### Prerequisites
- Go 1.20+
- Node.js 18+

### Terminal 1: Start Backend
```bash
cd Go_Project
go run ./cmd/students-api/main.go
```
Backend runs at `http://localhost:8080`

### Terminal 2: Start Frontend
```bash
cd students-frontend
npm install  # First time only
npm run dev
```
Frontend runs at `http://localhost:3000`

### Open Browser
Navigate to `http://localhost:3000`

---

## Features

### ✅ Implemented
- **Create**: Add students with name, age, email
- **Read**: View all students in responsive grid
- **Update**: Edit student details via modal
- **Delete**: Remove students with confirmation
- Beautiful UI with animations
- Form validation
- Error handling
- CORS support

---

## Tech Stack

| Layer | Tech |
|-------|------|
| **Backend** | Go 1.20+, SQLite, HTTP |
| **Frontend** | React 18, TypeScript, Vite, Axios |
| **Styling** | CSS3, Gradients, Animations |

---

## Project Structure

```
Go_Project/
├── cmd/students-api/main.go          # Backend entry point
├── internal/
│   ├── config/config.go              # Configuration
│   ├── http/handlers/student/         # API handlers
│   ├── storage/sqlite/sqlite.go       # Database
│   ├── types/types.go                 # Data types
│   └── utils/responses/               # Response helpers
├── config/local.yaml                 # Backend config
├── storage/storage.db                # SQLite database
├── students-frontend/
│   ├── src/
│   │   ├── components/               # React components
│   │   ├── pages/                    # Pages
│   │   ├── services/api.ts           # API client
│   │   └── types/index.ts            # Types
│   ├── package.json
│   ├── vite.config.ts
│   └── tsconfig.json
└── README.md                         # This file
```

---

## API Endpoints

**Base URL**: `http://localhost:8080/api/students`

| Method | Endpoint | Purpose |
|--------|----------|---------|
| POST | `/api/students` | Create student |
| GET | `/api/students` | Get all students |
| GET | `/api/students/{id}` | Get student by ID |
| PATCH | `/api/students/{id}` | Update student |
| DELETE | `/api/students/{id}` | Delete student |

### Example Request
```json
{
  "name": "John Doe",
  "age": 20,
  "email": "john@example.com"
}
```

---

## Database

SQLite schema:
```sql
CREATE TABLE students (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  age INTEGER NOT NULL
)
```

---

## Commands

### Backend
```bash
# Development
go run ./cmd/students-api/main.go

# Build
go build -o students-api ./cmd/students-api/main.go

# Run built binary
./students-api
```

### Frontend
```bash
cd students-frontend

# Development with hot reload
npm run dev

# Production build
npm run build

# Preview build
npm run preview
```

---

## Configuration

Backend config (`config/local.yaml`):
```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8080"
```

---

## Troubleshooting

| Issue | Solution |
|-------|----------|
| Backend won't start | Ensure port 8080 is free |
| Frontend shows "Loading..." | Verify backend is running |
| API requests fail | Check CORS headers are set |
| Database errors | Delete `storage/storage.db` and restart |

---

## Development Workflow

1. **Backend development**: Changes require restart
2. **Frontend development**: Vite hot reload enabled
3. **Database changes**: Edit SQLite directly or modify schema
4. **CORS**: Already configured in backend middleware

---

## Deployment

### Backend
```bash
# Build
go build -o students-api ./cmd/students-api/main.go

# Deploy executable to server
# Update config/local.yaml with production settings
```

### Frontend
```bash
# Build
npm run build

# Deploy dist/ folder to static hosting
```

---

## Learning Points

- RESTful API design
- Database schema and SQL
- Frontend-backend integration
- React hooks and component composition
- TypeScript type safety
- CORS and cross-origin requests
- HTTP error handling
- Form validation

---

## Notes

- Database auto-creates on first run
- Email validation enforced on backend
- Age constraints: 1-150
- Name length: 2-100 characters
- All endpoints return JSON
- DELETE endpoint sends no response body on success (204)

---

## Author

Created as a learning project demonstrating full-stack development with Go and React.
