HOST := localhost:9092
GET_HASHTAGS := hashtag_service.v1.HashtagService/get_hashtags
IMG_PATH := test/images

dev:
	. env/bin/activate && cd ../.. && python -m services.hashtag_service.src.presentation.grpc

venv:
	python3 -m venv env

install:
	. env/bin/activate && pip install --force-reinstall --no-deps -r requirements.txt

freeze:
	. env/bin/activate && pip freeze > requirements.txt

clean:
	rm -rf env

### Examples ###

req_animal:
	grpcurl -plaintext -d '{ "caption": "a cat is lying on a chair" }' $(HOST) $(GET_HASHTAGS)

req_obj:
	grpcurl -plaintext -d '{ "caption": "a sign on a fence" }' $(HOST) $(GET_HASHTAGS)

req_human:
	grpcurl -plaintext -d '{ "caption": "a girl with a flower in her hair" }' $(HOST) $(GET_HASHTAGS)