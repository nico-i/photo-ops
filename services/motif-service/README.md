# motif-service

This service provides a bounding box for the motif i.e. the main subject of a given image via an http or gRPC calls.

## Dependencies

This service depends on the following repositories:

- `Go`
  - [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway#grpc-gateway) — generation of gRPC stubs, OpenAPI specification and http gateway
- `Python`
  - [rembg](https://github.com/danielgatis/rembg#rembg) — removal of an image's background
  - [Pillow](https://github.com/python-pillow/Pillow#pillow) — denoising, bounding box calculation and drawing

## Credits

Example images used in the tests are from [Unsplash](https://unsplash.com/):

- [Object example photo](./docs/example/obj.jpg) by [Atul Vinayak](https://unsplash.com/@atulvi?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) on [Unsplash](https://unsplash.com/photos/text-9ZvjWPDb0fE?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
  