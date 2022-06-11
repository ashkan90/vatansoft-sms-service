FROM golang:1.18-alpine as build

WORKDIR /app

VOLUME ["/app"]

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY .. .

RUN GOOS=linux CGO_ENABLED=0 go build -o /bit-driver-location-service ./cmd/app

FROM alpine
COPY --from=build /bit-driver-location-service ./app

COPY ./server.yaml ./
EXPOSE 8080

CMD [ "./app" ]