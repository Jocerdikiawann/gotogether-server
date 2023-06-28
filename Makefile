gen:
	protoc --proto_path=model/proto \
	--openapiv2_out=docs \
    --openapiv2_opt=allow_merge=true \
	--go_out=model/pb --go_opt=paths=source_relative \
    --go-grpc_out=model/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=model/pb --grpc-gateway_opt=paths=source_relative \
    model/proto/*.proto
	wire gen ./di/injector.go
clean:
	rm -f model/pb/*.go
	rm ./di/wire_gen.go
	docker stop $$(docker ps -q)
	docker rm -f $$(docker ps -a -q)
	docker rmi $$(docker images -a -q)
	docker volume rm $$(docker volume ls -q)
build:
	docker-compose up -d
	@$(shell sleep 10)
	docker exec -it mongo1 mongosh --eval "rs.initiate({ \
		_id: \"myReplicaSet\", \
		members: [ \
			{_id: 0, host: \"mongo1\"}, \
			{_id: 1, host: \"mongo2\"}, \
			{_id: 2, host: \"mongo3\"} \
		] \
	})"
	@$(shell sleep 10)
	docker exec -it mongo1 mongosh --eval "rs.status()"
	@$(shell sleep 10)
run:
	go run .