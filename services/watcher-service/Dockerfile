# Use the Alpine base image
FROM alpine:3.20.0

# Update package lists and install inotify-tools
RUN apk add --no-cache inotify-tools curl

# Copy the monitor script into the image
COPY watcher.sh /watcher.sh

# Make the script executable
RUN chmod +x /watcher.sh

ENV MONITORED_DIR=./images

# Set the entrypoint to the script
ENTRYPOINT ["/watcher.sh"]