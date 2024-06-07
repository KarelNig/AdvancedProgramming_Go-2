FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service ./server

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/auth-service .

RUN chmod +x /root/auth-service .

EXPOSE 50051

CMD ["./auth-service"]
