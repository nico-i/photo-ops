module github.com/nico-i/photo-ops/services/motif_service_gateway

go 1.22.3

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/nico-i/photo-ops/shared v0.0.0
	google.golang.org/grpc v1.64.0
)

require (
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240513163218-0867130af1f8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240513163218-0867130af1f8 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

replace github.com/nico-i/photo-ops/shared v0.0.0 => ../../shared
