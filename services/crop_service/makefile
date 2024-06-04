HOST := localhost:9093
CROP_IMAGE := crop_service.v1.CropService/crop_image
CROP_SUGGESTION := crop_service.v1.CropService/get_crop_suggestion
IMG_PATH := test/images
SERVICE := crop_service
TARGET_ASPECT_RATIO := 0.8

include ../../shared/makefile


req_animal:
	grpcurl -plaintext -d '{ \
		"image_width": 1080, \
		"image_height": 720, \
		"target_aspect_ratio": $(TARGET_ASPECT_RATIO), \
		"image_motif": { \
			"x": 219, \
			"y": 65, \
			"width": 680, \
			"height": 655 \
		}}' $(HOST) $(CROP_SUGGESTION)

req_obj:
	grpcurl -plaintext -d '{ \
		"image_width": 1080, \
		"image_height": 720, \
		"target_aspect_ratio": $(TARGET_ASPECT_RATIO), \
		"image_motif": { \
			"x": 428, \
			"y": 195, \
			"width": 286, \
			"height": 463 \
		}}' $(HOST) $(CROP_SUGGESTION)

req_human:
	grpcurl -plaintext -d '{ \
		"image_width": 1080, \
		"image_height": 721, \
		"target_aspect_ratio": $(TARGET_ASPECT_RATIO), \
		"image_motif":  { \
			"x": 208, \
			"y": 176, \
			"width": 537, \
			"height": 371 \
		}}' $(HOST) $(CROP_SUGGESTION)