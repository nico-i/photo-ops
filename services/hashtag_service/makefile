HOST := localhost:9092
GET_HASHTAGS := hashtag_service.v1.HashtagService/get_hashtags
IMG_PATH := test/images
SERVICE := hashtag_service

include ../../shared/makefile

### Examples ###

req_animal:
	grpcurl -plaintext -d '{ "caption": "a cat is lying on a chair" }' $(HOST) $(GET_HASHTAGS)

req_obj:
	grpcurl -plaintext -d '{ "caption": "a sign on a fence" }' $(HOST) $(GET_HASHTAGS)

req_human:
	grpcurl -plaintext -d '{ "caption": "a girl with a flower in her hair" }' $(HOST) $(GET_HASHTAGS)