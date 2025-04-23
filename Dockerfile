FROM golang:1.23.4-alpine AS build
WORKDIR /app

COPY . .

ENV CONFIG_PATH=./config/config-docker.yaml
ENV MIGRATIONS_PATH=./migrations
RUN go run ./cmd/migrator/main.go

RUN mkdir -p ./build && go build -o ./build/main ./cmd/main/main.go

FROM alpine:latest AS production
WORKDIR /app

COPY --from=build /app/build .
COPY ./config/config-docker.yaml ./config/config-docker.yaml

ENV CONFIG_PATH=./config/config-docker.yaml

ENTRYPOINT ["/app/main"]