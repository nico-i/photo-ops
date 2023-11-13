# image-service

This service is responsible for receiving the initial image and passing the image through each of the image processing services.
After the image has been processed and the necessary metadata has been acquired, a resulting image with the metadata stored as attached IPTC data is then either returned to the user or saved to a local file.
