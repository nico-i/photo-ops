#!/bin/bash

# Check if MONITORED_DIR is set in the environment, if not exit with an error
if [ -z "$MONITORED_DIR" ]; then
    echo "Error: The environment variable MONITORED_DIR is not set." >&2
    exit 1
fi

# Check if STATE_FILE is set in the environment, if not exit with an error
if [ -z "$STATE_FILE" ]; then
    echo "Error: The environment variable STATE_FILE is not set." >&2
    exit 1
fi

# Check if UPLOAD_URL is set in the environment, if not exit with an error
if [ -z "$UPLOAD_URL" ]; then
    echo "Error: The environment variable UPLOAD_URL is not set." >&2
    exit 1
fi

# Check if STATE_FILE exists, if not, create it with the current directory state
if [ ! -f "$STATE_FILE" ]; then
    ls -l "$MONITORED_DIR" > "$STATE_FILE"
fi

# Check for new files
NEW_FILES=$(diff <(ls -l "$MONITORED_DIR") "$STATE_FILE" | grep "^>" | awk '{print $NF}')

# Loop over new files and upload them
for file in $NEW_FILES; do
    curl -F "file=@$MONITORED_DIR/$file" "$UPLOAD_URL"
done

# Update the state file
ls -l "$MONITORED_DIR" > "$STATE_FILE"
