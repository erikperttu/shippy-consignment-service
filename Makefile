build:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto
image:
	docker build -t shippy-consignment-service .
mongo: 
	docker run -p 27017:27017 mongo
run:
	docker run --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		-e DISABLE_AUTH=true \
		-e DB_HOST=192.168.99.100\
		shippy-consignment-service