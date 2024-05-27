module github.com/brunodev09/industrial-oms-go/gateway

go 1.22.0

require github.com/brunodev09/industrial-oms-go/common v0.0.0

replace github.com/brunodev09/industrial-oms-go/common => ../common

require (
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.64.0
)

require (
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
