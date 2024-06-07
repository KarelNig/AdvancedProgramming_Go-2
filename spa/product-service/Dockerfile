FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o product-service ./server

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/product-service .

RUN chmod +x /root/product-service .

EXPOSE 50053

CMD ["./product-service"]
