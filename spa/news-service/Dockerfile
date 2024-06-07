FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o news-service ./server

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/news-service .

RUN chmod +x /root/news-service .

EXPOSE 50054

CMD ["./news-service"]
