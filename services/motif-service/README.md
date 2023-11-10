# motif-service

This service provides the motif i.e. the main subject of a given image via an http or gRPC calls the user is able to either upload or specify the local path of an image to be processed.

Depending on the type of call the service can have the following outputs:

- a PNG with the background of the input image removed and only the main subject (motif) of the image remaining
- the X- and Y-coordinates of the corners of a bounding box encasing the motif
- the input image with the motif highlighted by a bounding box
- any combination of the above

## Dependencies

This service depends on the following repositories:

- `Go`
  - [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway#grpc-gateway) — generation of gRPC stubs, OpenAPI specification and http gateway
- `Python`
  - [rembg](https://github.com/danielgatis/rembg#rembg) — removal of an image's background
  - [Pillow](https://github.com/python-pillow/Pillow#pillow) — denoising, bounding box calculation and drawing

## Credits

Example images used in the tests are from [Unsplash](https://unsplash.com/).

- [Object example photo](./docs/example/obj.jpg) by <a href="https://unsplash.com/@atulvi?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Atul Vinayak</a> on <a href="https://unsplash.com/photos/text-9ZvjWPDb0fE?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
  