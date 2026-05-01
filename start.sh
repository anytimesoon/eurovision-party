#!/bin/bash

set -e

DOMAIN_NAME="${DOMAIN_NAME:-localhost}"
PUBLIC_DOMAIN_NAME="http://${DOMAIN_NAME}"
PUBLIC_GO_HOST="http://${DOMAIN_NAME}"
PUBLIC_CHAT="ws://${DOMAIN_NAME}"
CONFIG_FILE=/backend/conf/app.env

export PUBLIC_DOMAIN_NAME
export PUBLIC_GO_HOST
export PUBLIC_CHAT

set_config_value() {
    local key="$1"
    local value="$2"

    if grep -q "^${key}=" "$CONFIG_FILE"; then
        sed -i "s|^${key}=.*|${key}=${value}|" "$CONFIG_FILE"
    else
        echo "${key}=${value}" >> "$CONFIG_FILE"
    fi
}

echo "Building configs"
mkdir -p /backend/conf

if [ ! -f "$CONFIG_FILE" ]; then
    echo "Creating initial backend config"
    echo "
    ASSET_DIR=/backend/assets/img
    BACKEND_HOST=${BACKEND_HOST:-0.0.0.0:8080}
    BOT_ID=${BOT_ID:-}
    CHAT_BOT_NAME=${CHAT_BOT_NAME:-Eurobot}
    DB_PATH=storage/
    PUBLIC_DOMAIN_NAME=${PUBLIC_DOMAIN_NAME}
    MAX_INVITES=${MAX_INVITES:-5}
    SECRET=${SECRET:-}
    VOTE_COUNT_TRIGGER=${VOTE_COUNT_TRIGGER:-5}
    " > "$CONFIG_FILE"
else
    echo "Using existing backend config"
fi

set_config_value "ASSET_DIR" "/backend/assets/img"
set_config_value "BACKEND_HOST" "${BACKEND_HOST:-0.0.0.0:8080}"
set_config_value "DB_PATH" "storage/"
set_config_value "PUBLIC_DOMAIN_NAME" "${PUBLIC_DOMAIN_NAME}"
set_config_value "MAX_INVITES" "${MAX_INVITES:-5}"
set_config_value "VOTE_COUNT_TRIGGER" "${VOTE_COUNT_TRIGGER:-5}"

echo "
:80 {
	handle /api/* {
		reverse_proxy 127.0.0.1:8080
	}

	handle /chat/* {
		reverse_proxy 127.0.0.1:8080
	}

	handle /ws/* {
		reverse_proxy 127.0.0.1:8080
	}

  handle /restricted/* {
    reverse_proxy 127.0.0.1:8080
  }

  handle /content/* {
    reverse_proxy 127.0.0.1:8080
  }

	handle {
		reverse_proxy 127.0.0.1:3000
	}
}
" > /etc/caddy/Caddyfile

echo "Starting backend..."
cd /backend && ./app &
BACKEND_PID=$!

echo "Starting frontend..."
cd /frontend && node -r dotenv/config build &
FRONTEND_PID=$!

echo "Starting Caddy..."
caddy run --config /etc/caddy/Caddyfile --adapter caddyfile &
CADDY_PID=$!

wait_for_process() {
    local pid=$1
    local name=$2
    
    while kill -0 $pid 2>/dev/null; do
        sleep 1
    done
    
    echo "$name has stopped"
}

# Function to cleanup processes
cleanup() {
    echo "Shutting down..."
    kill $BACKEND_PID 2>/dev/null
    kill $FRONTEND_PID 2>/dev/null
    kill $CADDY_PID 2>/dev/null
    exit 0
}

# Set up signal handlers
trap cleanup SIGINT SIGTERM

echo "Both services started. Backend PID: $BACKEND_PID, Frontend PID: $FRONTEND_PID, Caddy PID: $CADDY_PID"
echo "Application is ready"

# Wait for either process to exit
wait
