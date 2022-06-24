FROM golang:1.18.1-alpine

RUN apk --no-cache add build-base git gcc openssh-client curl

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY .. .


RUN go install github.com/cosmtrek/air@latest
COPY ./.air.toml ./

ENTRYPOINT ["air", "-c", ".air.toml"]