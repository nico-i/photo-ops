HOST := localhost:9091
GET_CAPTION := caption_service.v1.CaptionService/get_caption
IMG_PATH := test/images

dev:
	. ./env/bin/activate && cd ../.. && python -m services.caption_service.src.presentation.grpc

venv:
	python3 -m venv env

install:
	. env/bin/activate && pip install --force-reinstall --no-deps -r requirements.txt

freeze:
	. ./env/bin/activate && pip freeze > requirements.txt

clean:
	rm -rf env

### Examples ###

req_animal:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/animal.jpg"}}' $(HOST) $(GET_CAPTION)

req_obj:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/object.jpg"}}' $(HOST) $(GET_CAPTION)

req_human:
	grpcurl -plaintext -d '{"image":{"path":"$(IMG_PATH)/human.jpg"}}' $(HOST) $(GET_CAPTION)
