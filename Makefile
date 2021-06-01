build:
	GOOS=windows go build&&GOOS=darwin go build -o wormhole.mac&&GOOS=linux go build
grpc:
	@protoc protos/redux/redux.proto --go-grpc_out=. --go_out=.
	@#protoc protos/redux/redux.proto --go_out=plugins=grpc:./
    #protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/redux/redux.proto
grpc-n:
	@protoc protos/get.proto --go-grpc_out=. --go_out=.