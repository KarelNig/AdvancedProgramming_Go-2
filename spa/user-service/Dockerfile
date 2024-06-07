FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service ./server

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/user-service .

RUN chmod +x /root/user-service .

EXPOSE 50052

CMD ["./user-service"]
