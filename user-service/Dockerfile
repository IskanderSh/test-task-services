FROM golang:alpine AS builder

WORKDIR /usr/local/src

ADD go.mod .

COPY . .

RUN go build -o ./bin/user-service cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/user-service .
COPY config/local.yaml .

ENV CONFIG_PATH=./local.yaml

CMD ["/user-service"]