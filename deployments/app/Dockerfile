FROM golang:1.18-alpine as build

WORKDIR /app

VOLUME ["/app"]

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY .. .

RUN GOOS=linux CGO_ENABLED=0 go build -o /vatansoft-sms-service ./cmd/app

FROM alpine
COPY --from=build /vatansoft-sms-service ./app

COPY ./server.yaml ./
EXPOSE 8080

CMD [ "./app" ]