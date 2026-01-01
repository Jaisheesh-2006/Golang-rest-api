# Project Setup Guide

## Overview

This is a full-stack application with:

- **Backend**: Go REST API server on port 8080
- **Frontend**: React + TypeScript application on port 3000

## Prerequisites

- Go 1.20 or higher
- Node.js 18+ and npm 9+
- SQLite (included with most systems)

## Backend Setup

### 1. Navigate to Backend Directory

```bash
cd Go_Project
```

### 2. Install Go Dependencies

```bash
go mod download
```

### 3. Configuration

- Configuration file: `config/local.yaml`
- Ensure the database path and port are correctly set

### 4. Run the Backend

```bash
go run cmd/students-api/main.go
```

The backend will start on `http://localhost:8080`

Expected output:

```
Server started at :8080
```

## Frontend Setup

### 1. Navigate to Frontend Directory

```bash
cd students-frontend
```

### 2. Install Dependencies (if not already done)

```bash
npm install
```

### 3. Development Mode

```bash
npm run dev
```

The frontend will start on `http://localhost:3000`

### 4. Production Build

```bash
npm run build
```

Output will be in the `dist/` directory.

## API Endpoints

All endpoints are prefixed with `/api/students`

### Create Student

```bash
POST /api/students
Content-Type: application/json

{
  "name": "John Doe",
  "age": 20,
  "grade": "A"
}
```

### Get All Students

```bash
GET /api/students
```

### Get Student by ID

```bash
GET /api/students/{id}
```

### Update Student

```bash
PATCH /api/students/{id}
Content-Type: application/json

{
  "name": "Jane Doe",
  "age": 21,
  "grade": "A+"
}
```

## Common Issues

### Port Already in Use

If port 8080 or 3000 is already in use:

- Go backend: Update `config/local.yaml` to use a different port
- React frontend: Set `VITE_PORT=3001` before running

### CORS Issues

The Vite dev server proxies `/api` requests to the backend, so CORS should work automatically.

### Database Connection Issues

- Ensure the SQLite database file path is correct in `config/local.yaml`
- Make sure the `storage/` directory exists

## Development Workflow

1. **Start Backend Terminal**

   ```bash
   cd Go_Project
   go run cmd/students-api/main.go
   ```

2. **Start Frontend Terminal**

   ```bash
   cd Go_Project/students-frontend
   npm run dev
   ```

3. Open `http://localhost:3000` in your browser

4. Make changes to either backend or frontend - they'll auto-reload

## Build for Production

### Backend

```bash
go build -o students-api cmd/students-api/main.go
./students-api
```

### Frontend

```bash
cd students-frontend
npm run build
# Serve the dist/ directory with any static server
```

## Troubleshooting

### Frontend shows "Loading students..." forever

- Check that the Go backend is running
- Verify the API URL in `src/services/api.ts` matches your backend address
- Check browser console for network errors

### Form submission fails

- Ensure the backend is running
- Check that the student data matches the expected format
- Look at the Go server logs for error messages

### Build fails

```bash
# Clear node_modules and reinstall
rm -rf node_modules
npm install
npm run build
```

## File Structure

```
Go_Project/
├── cmd/
│   └── students-api/
│       └── main.go          # Backend entry point
├── internal/
│   ├── config/              # Configuration
│   ├── http/                # HTTP handlers
│   ├── storage/             # Database layer
│   ├── types/               # Type definitions
│   └── utils/               # Utilities
├── storage/                 # Database files
├── config/
│   └── local.yaml          # Configuration file
├── go.mod                   # Go module definition
├── README.md               # Backend documentation
└── students-frontend/       # React frontend
    ├── src/
    │   ├── components/      # React components
    │   ├── pages/           # Page components
    │   ├── services/        # API client
    │   ├── types/           # TypeScript types
    │   ├── App.tsx
    │   └── main.tsx
    ├── public/              # Static assets
    ├── index.html           # HTML entry point
    ├── package.json         # npm configuration
    ├── tsconfig.json        # TypeScript configuration
    ├── vite.config.ts       # Vite configuration
    └── README.md            # Frontend documentation
```

## Next Steps

1. Start both the backend and frontend servers
2. Navigate to `http://localhost:3000`
3. Add some students using the form
4. View the student list automatically update
5. Explore the code and make modifications

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [React Documentation](https://react.dev/)
- [Vite Documentation](https://vitejs.dev/)
- [TypeScript Documentation](https://www.typescriptlang.org/)
