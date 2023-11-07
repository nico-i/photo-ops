# Use a base image with cron installed, for example, an Ubuntu image
FROM alpine:3.18.4

# Install curl, bash, and cron
RUN apk add --no-cache curl bash dcron

# Set the directory for the operations
WORKDIR /app

# Copy the script into the container
COPY watcher.sh .
# Make sure the script is executable
RUN chmod +x watcher.sh

# Setup the cron job to run the script every minute
RUN echo '* * * * * /app/watcher.sh' > /etc/crontabs/root

# Start the cron daemon
CMD ["crond", "-f", "-d", "8"]
