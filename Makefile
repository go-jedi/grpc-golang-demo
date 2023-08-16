gen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

clean:
	rm pb/*.go

test:
	go test -cover -race ./...

run-server:
	go run server/*.go

run-client:
	go run client/*.go


update-packages:
	go get -u ./...