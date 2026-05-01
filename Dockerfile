# Multi-stage Dockerfile that builds both frontend and backend

# STAGE 1: Build Go backend
FROM golang:alpine AS backend-build
RUN apk add --no-cache git ca-certificates

# add a user for security
RUN addgroup -S myapp && adduser -S -u 10000 -g myapp myapp

WORKDIR /src
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend ./

# Build the Go executable
RUN CGO_ENABLED=0 go build \
    -C pkg/cmd \
    -installsuffix 'static' \
    -o /app .

# STAGE 2: Build SvelteKit frontend
FROM node:19 as frontend-build
ENV NODE_ENV=production

WORKDIR /app
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

COPY frontend ./
RUN npm run build

# STAGE 3: Create runtime environment with both applications
FROM node:19-alpine AS final
LABEL maintainer="anytimesoon"

RUN apk add --no-cache bash caddy

# Create directories for both applications
RUN mkdir -p /backend /frontend

# Copy backend executable and required files
COPY --from=backend-build /app /backend/app
COPY --from=backend-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=backend-build /etc/passwd /etc/passwd
ADD backend/assets/img.tar.gz /backend/assets


# Copy frontend build
COPY --from=frontend-build /app /frontend

# Copy startup script
COPY start.sh /start.sh

# Set permissions
RUN chmod +x /backend/app
RUN chmod +x /start.sh

# Create volumes for backend
VOLUME /backend/conf
VOLUME /backend/tmp
VOLUME /backend/storage

# Set environment variables
ENV BODY_SIZE_LIMIT=0
ENV NODE_ENV=production

ENV DOMAIN_NAME=localhost
ENV BACKEND_HOST=0.0.0.0:8080
ENV CHAT_BOT_NAME=Eurobot
ENV MAX_INVITES=5
ENV VOTE_COUNT_TRIGGER=5
ENV INSECURE=false

EXPOSE 80

ENTRYPOINT ["/start.sh"]
