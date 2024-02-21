# Use the official golang image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go get -d -v ./...
RUN go install -v ./...

# Expose the Lambda port
EXPOSE 8080

# Set the environment variable to run the Lambda handler
CMD ["app"]
