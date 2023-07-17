gen:
	wire gen ./di/injector.go
	protoc --proto_path=model/proto \
	--openapiv2_out=docs \
    --openapiv2_opt=allow_merge=true \
	--go_out=model/pb --go_opt=paths=source_relative \
    --go-grpc_out=model/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=model/pb --grpc-gateway_opt=paths=source_relative \
    model/proto/*.proto
	
clean:
	rm ./di/wire_gen.go
	rm ./model/pb/*.go
	docker stop $$(docker ps -q)
	docker rm -f $$(docker ps -a -q)
	docker rmi $$(docker images -a -q)
	docker volume rm $$(docker volume ls -q)
	
build:
	docker-compose up -d
run:
	go run .