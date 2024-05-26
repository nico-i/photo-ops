module github.com/nico-i/photo-ops/services/caption_service_gateway

go 1.22.3

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/nico-i/photo-ops/shared v0.0.0
	google.golang.org/grpc v1.64.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.34.1
)

require (
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240521202816-d264139d666e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240515191416-fc5f0ca64291 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/nico-i/photo-ops/shared v0.0.0 => ../../shared
