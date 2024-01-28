run:
	go run cmd/main.go --config=./config/local.yaml

docker-run:
	docker run \
	--rm --name postgres \
	-e POSTGRES_PASSWORD=postgres \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_DB=postgres \
	-p 5432:5432 \
	-d postgres:latest

migrations-up:
	goose -dir migrations postgres "port=5432 host=localhost user=postgres password=postgres dbname=postgres" up
