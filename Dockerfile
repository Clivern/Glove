FROM golang:1.23 as build

RUN apt-get update
RUN mkdir -p /app

WORKDIR /app

COPY ./ ./

RUN go build -v -ldflags="-X 'main.version=v1.0.0' -X 'main.builtBy=docker'" main.go

FROM ubuntu:22.04

RUN apt-get update
RUN apt-get install curl -y

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs

WORKDIR /app

COPY --from=build /app/main /app/main
COPY --from=build /app/config.dist.yml /app/configs/config.dist.yml

EXPOSE 9253

VOLUME /app/configs
VOLUME /app/var

RUN ./main version

CMD ["./main", "server", "-c", "/app/configs/config.dist.yml"]
