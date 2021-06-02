FROM golang:1.12-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/orderbook-server

COPY . .

# run tests
RUN CGO_ENABLED=0 go test ./...

# build application
RUN go build -o ./out/orderbook-server .

# run the application with smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

# copy binary from build image
COPY --from=build_base /tmp/orderbook-server/out/orderbook-server /app/orderbook-server

EXPOSE 8080

CMD ["/app/orderbook-server"]