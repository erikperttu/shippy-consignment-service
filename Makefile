build:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto
image:
	docker build -t shippy-consignment-service .
run:
	docker run -d --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		-e DISABLE_AUTH=true \
		shippy-consignment-service