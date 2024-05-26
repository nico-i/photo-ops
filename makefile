dev-motif:
	python -m services.motif_service.src.presentation.grpc

dev-motif-gw:
	go run serivces/motif_service_gateway/src/presentation/http/server.go

dev-caption:
	python -m services.caption_service.src.presentation.grpc

dev-caption-gw:
	go run serivces/caption_service_gateway/src/presentation/http/server.go