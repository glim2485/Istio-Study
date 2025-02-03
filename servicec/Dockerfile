FROM golang:1.23.4

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory into the container
COPY . .

# Update and install netstat, curl, iproute2 (for ip route), and iputils-ping (for ping)
RUN apt-get update && apt-get install -y \
    net-tools \
    curl \
    iproute2 \
    iputils-ping && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Go dependencies
RUN go get .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 6789

# Run the application
CMD ["./main"]