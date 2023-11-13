#!/bin/sh

# Check if MONITORED_DIR is set in the environment, if not exit with an error
if [ -z "$MONITORED_DIR" ]; then
    echo "Error: The environment variable MONITORED_DIR is not set." >&2
    exit 1
fi

# Check if UPLOAD_URL is set in the environment, if not exit with an error
if [ -z "$UPLOAD_URL" ]; then
    echo "Error: The environment variable UPLOAD_URL is not set." >&2
    exit 1
fi

# Ensure the monitored directory exists
mkdir -p ${MONITORED_DIR}

inotifywait -m -e create -e modify -e delete -e close_write --format '%w%f %e' "${MONITORED_DIR}" | while read FILE EVENT
do
    echo "File ${FILE} had event ${EVENT}"
    curl -F "file=@$MONITORED_DIR/$file" "$UPLOAD_URL"
done

echo "Error: inotifywait exited unexpectedly." >&2