### motif_service ###

env-motif:
	echo "./services/motif_service/env/bin/activate"

dev-motif:
	python -m services.motif_service.src.presentation.grpc

dev-motif-gw:
	cd ./services/motif_service_gateway && go run src/presentation/http/server.go


### caption_service ###

env-cap:
	echo "./services/caption_service/env/bin/activate"

dev-cap:
	python -m services.caption_service.src.presentation.grpc

dev-cap-gw:
	cd ./services/caption_service_gateway && go run src/presentation/http/server.go

### hashtag_service ###

env-hash:
	echo "./services/hashtag_service/env/bin/activate"

dev-hash:
	python -m services.hashtag_service.src.presentation.grpc

dev-hash-gw:
	cd ./services/hashtag_service_gateway && go run src/presentation/http/server.go
	