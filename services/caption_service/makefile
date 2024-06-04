HOST := localhost:9091
GET_CAPTION := caption_service.v1.CaptionService/get_caption
IMG_PATH := test/images
SERVICE := caption_service

include ../../shared/makefile

### Examples ###

# Full image examples

req_animal:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/animal.jpg"}}' $(HOST) $(GET_CAPTION)

req_obj:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/object.jpg"}}' $(HOST) $(GET_CAPTION)

req_human:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/human.jpg"}}' $(HOST) $(GET_CAPTION)


# Cropped examples (currently the test data equals the returned bounding box from the motif service)

req_animal_crop:
	grpcurl -plaintext -d '{ \
		"image":{ \
			"path":"$(IMG_PATH)/animal.jpg" \
		}, \
		"crop": { \
			"x": 219, \
			"y": 65, \
			"width": 680, \
			"height": 655 \
		}}' $(HOST) $(GET_CAPTION)

req_obj_crop:
	grpcurl -plaintext -d '{ \
		"image":{ \
			"path":"$(IMG_PATH)/object.jpg" \
		}, \
		"crop": { \
			"x": 428, \
			"y": 195, \
			"width": 286, \
			"height": 463 \
		}}' $(HOST) $(GET_CAPTION)

req_human_crop:
	grpcurl -plaintext -d '{ \
		"image":{ \
			"path":"$(IMG_PATH)/human.jpg" \
		}, \
		"crop": { \
			"x": 208, \
			"y": 176, \
			"width": 537, \
			"height": 371 \
		}}' $(HOST) $(GET_CAPTION)