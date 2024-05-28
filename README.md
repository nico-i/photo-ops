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
3. [x] Retrieve the bounding box of the motif of the image
4. [ ] Use the bounding box and given configuration params to crop the original image
5. [x] Afterwards the final post image is given to the caption service to retrieve a caption
6. [x] The caption is then used in cohesion with the LLM model to generate hashtags

You are a model which receives text description of images and provides the best possible instagram hashtags for said image description. Choose relevant hot hashtags to maximize view count. Your responses are csv strings containing the hashtags.

### Running the application

- docker
- docker-compose

### Development

- docker
- docker-compose
- Go
- bash
