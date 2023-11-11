FROM python:3.11.6-alpine3.18

ARG GO_VERSION=1.21.4

# Install Go depending on arg
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
RUN rm -rf go${GO_VERSION}.linux-amd64.tar.gz

# Set the working directory in the container
WORKDIR /app

# Copy the Go service and Python scripts into the container
COPY . .

# Install Go dependencies
# Ensure your Go project has a go.mod and go.sum file
RUN go mod download

# Upgrade pip
RUN pip install --upgrade pip

# Install Python dependencies
# Ensure you have a requirements.txt file
RUN pip install -r requirements.txt

# Build the Go service
RUN go build -o motif-service

# Make the Python script executable
RUN chmod +x myscript.py