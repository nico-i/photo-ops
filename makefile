env-motif:
	echo "./services/motif_service/env/bin/activate"

dev-motif:
	python -m services.motif_service.src.presentation.grpc

dev-motif-gw:
	cd ./services/motif_service_gateway && go run src/presentation/http/server.go

env-caption:
	echo "./services/caption_service/env/bin/activate"

dev-caption:
	python -m services.caption_service.src.presentation.grpc

dev-caption-gw:
	cd ./services/caption_service_gateway && go run src/presentation/http/server.go