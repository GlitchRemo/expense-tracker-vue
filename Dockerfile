# ----- Build Frontend -----
    FROM node:18-alpine AS frontend-builder
    WORKDIR /app
    COPY frontend/ ./
    RUN npm install && npm run build
    
    # ----- Build Backend -----
    FROM golang:1.21-alpine AS backend-builder
    WORKDIR /app
    COPY backend/ ./
    COPY --from=frontend-builder /app/dist ./frontend
    RUN go build -o main .
    
    # ----- Final Minimal Image -----
    FROM alpine:latest
    WORKDIR /root/
    COPY --from=backend-builder /app/main .
    COPY --from=backend-builder /app/frontend ./frontend
    
    # If your Go app serves frontend (e.g., via embed or static handler)
    EXPOSE 8080
    CMD ["./main"]
    