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

## User Journey

1. Add images
   1. press the + button or drag images into the home page
   2. prompt appears which asks how the images should be processed (selections are stored as defaults for later additions)
2. begin processing (multithreaded => start new service instances)
   1. resize image whilst keeping the aspect ratio so that width is 1080px (maximum instagram width)
   2. read location from image data => if exists ? save to DB : skip location data
   3. losslessly compress images using squoosh
   4. if AI crop is active
      1. determine motif using motif service and save bbox to DB
      2. if defined crop ratio is fittable with motif ? auto-crop and save crop to DB : **mark image as input needed**
   5. else
      1. **mark image as input is needed**
   6. ai generate image caption and save to DB as alttext
   7. if ai hashtag is active ? generate ai caption with hashtags and save to db : **mark image as input needed**
3. MVP: User can press show JSON button to receive a JSON overview of all the details relating to an image when it is done processing
