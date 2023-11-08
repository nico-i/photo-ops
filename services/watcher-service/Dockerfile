# Use the Alpine base image
FROM alpine:3.18.4

# Update package lists and install inotify-tools
RUN apk add --no-cache inotify-tools curl

# Copy the monitor script into the image
COPY watcher.sh /watcher.sh

# Make the script executable
RUN chmod +x /watcher.sh

# Set the entrypoint to the script
ENTRYPOINT ["/watcher.sh"]