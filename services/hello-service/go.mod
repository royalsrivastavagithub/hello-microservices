module github.com/royalsrivastavagithub/hello-microservices/services/hello-service

replace github.com/royalsrivastavagithub/hello-microservices/pb => ../../pb

go 1.24.5

require (
	github.com/royalsrivastavagithub/hello-microservices/pb v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.73.0
)

require (
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
