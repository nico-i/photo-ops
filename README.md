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

1. [x] Receive a new image
2. [x] Store original image in local storage and save its path
3. [x] Retrieve the bounding box of the motif of the image and store it
4. [ ] Crop (and compress) the image (resolution is based on a setting) so that the bounding box is centered
   1. [ ] Provide the user the option to re-crop the image, especially if the bounding box is too large for the resolution
5. [x] Generate a caption for the image
   1. [ ] Give the user the option to edit the caption
6. [ ] Watermark the image (location is again based on a setting)
7. [x] Generate hashtags based on the caption

You are a model which receives text description of images and provides the best possible instagram hashtags for said image description. Choose relevant hot hashtags to maximize view count. Your responses are csv strings containing the hashtags.

### Running the application

- docker
- docker-compose

### Development

- docker
- docker-compose
- Go
- bash
