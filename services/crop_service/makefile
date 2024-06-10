HOST := localhost:9093
CROP_IMAGE := crop_service.v1.CropService/crop_image
CROP_SUGGESTION := crop_service.v1.CropService/get_crop_suggestion
IMG_PATH := test/images
OUTPUT_PATH := ./services/crop_service/docs/img
SERVICE := crop_service
TARGET_ASPECT_RATIO := 0.8

include ../../shared/makefile


req_animal_rec:
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

req_animal_crop:
	grpcurl -plaintext -d '{ \
		"output_path": "$(OUTPUT_PATH)/animal_crop.jpg", \
		"input_image": { \
			"path": "$(IMG_PATH)/animal.jpg" \
		}, \
		"crop_spec": { \
			"x": 296, \
			"y": 65, \
			"width": 524, \
			"height": 655 \
		}}' $(HOST) $(CROP_IMAGE)

req_obj_rec:
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

req_obj_crop:
	grpcurl -plaintext -d '{ \
		"output_path": "$(OUTPUT_PATH)/object_crop.jpg", \
		"input_image": { \
			"path": "$(IMG_PATH)/object.jpg" \
		}, \
		"crop_spec": { \
			"x": 336, \
			"y": 133, \
			"width": 469, \
			"height": 587 \
		}}' $(HOST) $(CROP_IMAGE)

req_human_rec:
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

req_human_crop:
	grpcurl -plaintext -d '{ \
		"output_path": "$(OUTPUT_PATH)/human_crop.jpg", \
		"input_image": { \
			"path": "$(IMG_PATH)/human.jpg" \
		}, \
		"crop_spec": { \
			"x": 188, \
			"y": 2, \
			"width": 575, \
			"height": 718 \
		}}' $(HOST) $(CROP_IMAGE)