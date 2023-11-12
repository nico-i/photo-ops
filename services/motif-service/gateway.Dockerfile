FROM golang:1.21.4-alpine3.18

# Set the working directory in the container
WORKDIR /motif-service

# Copy the Go service and Python scripts into the container
COPY . .

# Install Go dependencies
RUN go mod download

# Build the Go service
RUN go build -o ./gateway presentation/http/gateway.go

# Define an environment variable for the port and the grpc server port
ENV GATEWAY_PORT=8080
ENV SERVER_PORT=9000

# Run the Go service
CMD ./gateway -p "$GATEWAY_PORT" -serv-p "$SERVER_PORT"