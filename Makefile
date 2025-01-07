# Makefile for running go in the background

# Define the app command as a variable for easier reuse
CMD = nohup go run main.go &

# Target to start the go app in the background
start:
    $(CMD)
    @echo "go app started on port 8013."

# Optional: Target to stop the go app (find and kill the process)
stop:
    @pkill -f "go run main.go"
    @echo "go app stopped."

# Optional: Target to restart the go app
restart: stop start