#!/bin/bash

# Start the backend in the background
echo "Starting backend..."
/backend/app &
BACKEND_PID=$!

# Start the frontend in the background
echo "Starting frontend..."
cd /frontend && node -r dotenv/config build &
FRONTEND_PID=$!

# Wait for both processes and exit if any one fails
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
