# photo-ops
A pipeline passing an image between different micro services to crop, caption, watermark, hashtagify and finally post it to social media.

NOTE: Currently all models used are cpu based.

## Dependencies

- [Imaginary](https://github.com/h2non/imaginary)
  - Responsibilities:
    - Image cropping
    - Image resizing
    - Image compression
- [rembg](https://github.com/danielgatis/rembg)
  - Responsibilities:
    - Image background removal -> provides subject of image

## Process

1. [x] Receive request with image
2. [x] Store original image in local storage
3. [ ] Compress image and resize it to a point where it is just visible what the image is
4. [x] Retrieve the bounding box of the motif of the image
5. [ ] Use the bounding box and given configuration params to crop the original image
6. [x] Afterwards the final post image is given to the caption service to retrieve a caption
7. [ ] The caption is then used in cohesion with the LLM model to generate hashtags
8. [ ] The caption and hashtags are then stored as IPTC data in the image

### Running the application

- docker
- docker-compose

### Development

- docker
- docker-compose
- Go
- bash
