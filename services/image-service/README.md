# image-service

This service is responsible for receiving the initial image and passing the image through each of the image processing services.
After the image has been processed and the necessary metadata has been acquired, a resulting image with the metadata stored as attached IPTC data is then either returned to the user or saved to a local file.

## Process

1. Receive request with image
2. Store original image in local storage
3. Compress image and resize it to a point where it is just visible what the image is
4. Retrieve the bounding box of the motif of the image
5. Use the bounding box and given configuration params to crop the original image
6. Afterwards the final post image is given to the caption service to retrieve a caption
7. The caption is then used in cohesion with the LLM model to generate hashtags
8. The caption and hashtags are then stored as IPTC data in the image

Conclusion: Due to the sequentiality of this process, a micro service architecture is not necessary. The image service can be implemented as a single service.