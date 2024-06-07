package main

import (
    "context"
    "log"
    "time"

    newspb "news-service"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:82", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := newspb.NewNewsServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    req := &newspb.GetNewsRequest{
        NewsId: "1",
    }
    res, err := client.GetNews(ctx, req)
    if err != nil {
        log.Fatalf("could not get news: %v", err)
    }
    log.Printf("News: %v", res.News)
}