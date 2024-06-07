package main

import (
	pb "auth"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

var integrationServer *server

func truncateTables(db *sql.DB) {
    _, err := db.Exec(`DO
    $$
    DECLARE
        r RECORD;
    BEGIN
        FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
            EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' CASCADE';
        END LOOP;
    END
    $$;`)
    if err != nil {
        log.Fatalf("failed to truncate tables: %v", err)
    }
}

func TestIntegrationRegisterAndLogin(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:1111@localhost:5432/spa_salon?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func(){
		truncateTables(db)
		db.Close()
	}()

	integrationServer = &server{db: db}

	registerReq := &pb.RegisterRequest{
		Username: "integuser",
		Email:    "integuser@example.com",
		Password: "password",
	}

	registerRes, err := integrationServer.Register(context.Background(), registerReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if registerRes.UserId == "" {
		t.Fatalf("Expected user ID, got empty string")
	}

	loginReq := &pb.LoginRequest{
		Username: "integuser",
		Password: "password",
	}

	loginRes, err := integrationServer.Login(context.Background(), loginReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if loginRes.UserId != registerRes.UserId {
		t.Fatalf("Expected user ID to be %v, got %v", registerRes.UserId, loginRes.UserId)
	}

}

func TestIntegrationValidateAndRefreshToken(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:1111@localhost:5432/spa_salon?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func(){
		truncateTables(db)
		db.Close()
	}()

	integrationServer = &server{db: db}

	registerReq := &pb.RegisterRequest{
		Username: "integuser2",
		Email:    "integuser2@example.com",
		Password: "password",
	}

	_, err = integrationServer.Register(context.Background(), registerReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	loginReq := &pb.LoginRequest{
		Username: "integuser2",
		Password: "password",
	}

	loginRes, err := integrationServer.Login(context.Background(), loginReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	validateReq := &pb.TokenRequest{
		Token: loginRes.Token,
	}

	validateRes, err := integrationServer.ValidateToken(context.Background(), validateReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !validateRes.Valid {
		t.Fatalf("Expected token to be valid")
	}

	refreshReq := &pb.TokenRequest{
		Token: loginRes.Token,
	}

	refreshRes, err := integrationServer.RefreshToken(context.Background(), refreshReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !refreshRes.Valid {
		t.Fatalf("Expected token to be valid")
	}

	if refreshRes.NewToken == "" {
		t.Fatalf("Expected new token, got empty string")
	}

}
