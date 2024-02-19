run:
	go run cmd/main.go --config=./config/local.yaml

app-up:
	docker build -t user-service .
	docker run \
	--name user-service \
	-p 1111:1111 \
	-d user-service
