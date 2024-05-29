HOST := localhost:9090
GET_B_BOX := motif_service.v1.MotifService/get_b_box
GET_B_BOX_DEBUG := motif_service.v1.MotifService/get_b_box_debug
IMG_PATH := test/images

dev:
	. env/bin/activate && cd ../.. && python -m services.motif_service.src.presentation.grpc

venv:
	python3 -m venv env

install:
	. env/bin/activate && pip install --force-reinstall --no-deps -r requirements.txt

freeze:
	. env/bin/activate && pip freeze > requirements.txt

clean:
	rm -rf env

### Examples ###

# Bounding box JSON

req_animal:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/animal.jpg"}}' $(HOST) $(GET_B_BOX)

req_obj:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/object.jpg"}}' $(HOST) $(GET_B_BOX)

req_human:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/human.jpg"}}' $(HOST) $(GET_B_BOX)

# Debug bounding box images

req_animal_debug:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/animal.jpg"}}' $(HOST) $(GET_B_BOX_DEBUG)

req_obj_debug:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/object.jpg"}}' $(HOST) $(GET_B_BOX_DEBUG)

req_human_debug:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/human.jpg"}}' $(HOST) $(GET_B_BOX_DEBUG)