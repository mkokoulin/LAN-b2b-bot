# # Use an official Golang runtime as a parent image
# FROM golang:latest

# # Set the working directory to /app
# WORKDIR /app

# # Copy the current directory contents into the container at /app
# COPY . /app

# # Download and install any required dependencies
# RUN go mod download

# # Build the Go app
# # RUN go build -o main .

# RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# # Expose port 8080 for incoming traffic
# EXPOSE 8081

# # Define the command to run the app when the container starts
# CMD ["/app/main"]

# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Disable Go Modules
ENV GO111MODULE=off

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]