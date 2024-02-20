run:
	go run cmd/main.go --config=./config/local.yaml

app-up:
	docker build -t user-service .
	docker run \
	--name user-service \
	-p 1111:1111 \
	-d user-service

postgres-up:
	docker run \
    --rm --name postgres \
  	--net mynetwork \
	-e POSTGRES_PASSWORD=postgres \
 	-e POSTGRES_USER=postgres \
	-e POSTGRES_DB=postgres \
	-d postgres:latest

migrations-up:
	goose -dir migrations postgres "port=5432 host=localhost user=postgres password=postgres dbname=postgres" up

docker-up:
	docker-compose up -d
