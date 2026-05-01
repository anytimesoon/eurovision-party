#!/bin/bash

set -e

echo "Building config"
mkdir -p /backend/conf
echo "
ASSET_DIR=/backend/assets/img
BACKEND_HOST=${BACKEND_HOST:-0.0.0.0:8080}
BOT_ID=${BOT_ID:-}
CHAT_BOT_NAME=${CHAT_BOT_NAME:-Eurobot}
DB_PATH=storage/
PUBLIC_DOMAIN_NAME=${DOMAIN_NAME}
MAX_INVITES=${MAX_INVITES:-5}
SECRET=${SECRET:-}
VOTE_COUNT_TRIGGER=${VOTE_COUNT_TRIGGER:-5}
" > /backend/conf/app.env

echo "Starting backend..."
cd /backend && ./app &
BACKEND_PID=$!

echo "Starting frontend..."
cd /frontend && node -r dotenv/config build &
FRONTEND_PID=$!

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
    exit 0
}

# Set up signal handlers
trap cleanup SIGINT SIGTERM

echo "Both services started. Backend PID: $BACKEND_PID, Frontend PID: $FRONTEND_PID"
echo "Use Ctrl+C to stop both services"

# Wait for either process to exit
wait
