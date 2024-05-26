ARG PORT

FROM python:3.12-slim

# Set the working directory in the container
WORKDIR /motif-service

# Copy the Go service and Python scripts into the container
COPY . .

# Install opencv dependencies
RUN apt-get update && \
    apt-get install ffmpeg libsm6 libxext6 -y && \
    apt-get clean
    
# Upgrade pip
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# Install Go depending on arg
ADD https://golang.org/dl/go1.21.4.linux-amd64.tar.gz go1.21.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
RUN rm -rf go1.21.4.linux-amd64.tar.gz

# Install Go dependencies
RUN go mod download
# Build the Go service
RUN go build -o ./server presentation/grpc/server.go

# Define an environment variable for the port (you can set a default if needed)
ENV PORT=9000

# Run the Go service
CMD ./server -p "$PORT"

