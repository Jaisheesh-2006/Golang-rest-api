# Stage 1: Build Frontend
FROM node:18-alpine AS frontend-builder

WORKDIR /frontend

COPY students-frontend/package*.json ./

RUN npm ci

COPY students-frontend/ .

RUN npm run build

# Stage 2: Build Backend
FROM golang:1.25-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o students-api ./cmd/students-api/main.go

RUN mkdir -p storage

# Stage 3: Final Runtime Image
FROM gcr.io/distroless/cc-debian11

WORKDIR /app

# Copy built backend binary
COPY --from=backend-builder /app/students-api .

# Copy backend configuration
COPY config/local.yaml ./config/

# Copy built frontend files
COPY --from=frontend-builder /frontend/dist ./public

# Copy storage directory
COPY --from=backend-builder /app/storage ./storage

# Expose port
EXPOSE 8080

# Run the application with config path
CMD ["./students-api", "-config", "./config/local.yaml"]

