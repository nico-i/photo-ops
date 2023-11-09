# photo-ops
A pipeline passing an image between different micro services to crop, caption, watermark, hashtagify and finally post it to social media.

## Dependencies

- [Imaginary](https://github.com/h2non/imaginary)
  - Responsibilities:
    - Image cropping
    - Image resizing
    - Image compression
- [rembg](https://github.com/danielgatis/rembg)
  - Responsibilities:
    - Image background removal -> provides subject of image

## Prerequisites

### Running the application

- docker
- docker-compose

### Development

- docker
- docker-compose
- Go
- bash
