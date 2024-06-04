# photo-ops
A pipeline passing an image between different micro services to crop, caption, watermark, hashtagify and finally post it to social media.

NOTE: Currently all models used are cpu based.

## Services

- [Motif Service](./services/motif_service/README.md)
- [Caption Service](./services/caption_service/README.md)
- [Hashtag Service](./services/hashtag_service/README.md)
- [Watermark Service](./services/watermark_service/README.md)
- [Crop Service](./services/crop_service/README.md)

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

## FRs

- [x] User can upload an image
- [ ] User can crop the image
- [ ] User gets a recommended crop based on the subject of the image (if the bounding box is not too large)
- [ ] User can edit the caption
- [x] User gets a recommended caption based on the cropped image
- [ ] User can edit the hashtags
- [x] User gets a recommended hashtags based on the caption
- [ ] User can specify the watermark
- [ ] User can freely move the watermark
- [ ] User can specify the resolution of the image

## Technical Requirements

- [ ] The application should be able to run on a local machine
- [ ] The application should be able to run on a server
- [ ] The application should be able to run on a cloud provider
- [ ] The application should be able to run on a container
- [ ] The application should be able to run on a container orchestrator
- [ ] All but the service accessed by the FE should only communicate via gRPC

## Running the application

- docker
- docker-compose

## Development

- docker
- docker-compose
- Go
- bash
