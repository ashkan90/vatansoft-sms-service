FROM golang:1.18-alpine as build

WORKDIR /consumer

VOLUME ["/consumer"]

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY .. .

RUN GOOS=linux CGO_ENABLED=0 go build -o /vatansoft-sms-consumer-service ./cmd/consumer

FROM alpine
COPY --from=build /vatansoft-sms-consumer-service ./consumer

COPY ./server.yaml ./
EXPOSE 3000

CMD [ "./consumer" ]