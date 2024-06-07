package main

import (
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	pb "user-service"
)

var integrationUserServer *server


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

func TestIntegrationCreateAndGetUser(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:1111@localhost:5432/spa_salon?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func(){
		truncateTables(db)
		db.Close()
	}()

	integrationUserServer = &server{db: db}

	createReq := &pb.CreateUserRequest{
		Name:    "Jane Doe",
		Email:   "jane@example.com",
		Phone:   "0987654321",
		Address: "456 Avenue",
	}

	createRes, err := integrationUserServer.CreateUser(context.Background(), createReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if createRes.UserId == "" {
		t.Fatalf("Expected user ID, got empty string")
	}

	getReq := &pb.GetUserRequest{
		UserId: createRes.UserId,
	}

	getRes, err := integrationUserServer.GetUser(context.Background(), getReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if getRes.User.UserId != createRes.UserId {
		t.Fatalf("Expected user ID to be %v, got %v", createRes.UserId, getRes.User.UserId)
	}
}

func TestIntegrationUpdateAndDeleteUser(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:1111@localhost:5432/spa_salon?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func(){
		truncateTables(db)
		db.Close()
	}()

	integrationUserServer = &server{db: db}

	createReq := &pb.CreateUserRequest{
		Name:    "Jane Smith",
		Email:   "jane.smith@example.com",
		Phone:   "0987654322",
		Address: "789 Boulevard",
	}

	createRes, err := integrationUserServer.CreateUser(context.Background(), createReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	updateReq := &pb.UpdateUserRequest{
		UserId:  createRes.UserId,
		Name:    "Jane Doe",
		Email:   "jane.doe@example.com",
		Phone:   "1234567890",
		Address: "123 Street",
	}

	updateRes, err := integrationUserServer.UpdateUser(context.Background(), updateReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if updateRes.User.Name != "Jane Doe" {
		t.Fatalf("Expected name to be 'Jane Doe', got %v", updateRes.User.Name)
	}

	deleteReq := &pb.DeleteUserRequest{
		UserId: createRes.UserId,
	}

	deleteRes, err := integrationUserServer.DeleteUser(context.Background(), deleteReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !deleteRes.Success {
		t.Fatalf("Expected success to be true")
	}
}

func TestIntegrationListUsers(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:1111@localhost:5432/spa_salon?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func(){
		truncateTables(db)
		db.Close()
	}()

	integrationUserServer = &server{db: db}

	createReq := &pb.CreateUserRequest{
		Name:    "Jane Doe",
		Email:   "jane.doe@example.com",
		Phone:   "1234567890",
		Address: "123 Street",
	}

	_, err = integrationUserServer.CreateUser(context.Background(), createReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	listReq := &pb.ListUsersRequest{
		Filter:   "Jane",
		Page:     1,
		PageSize: 10,
	}

	listRes, err := integrationUserServer.ListUsers(context.Background(), listReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(listRes.Users) == 0 {
		t.Fatalf("Expected at least one user, got %v", len(listRes.Users))
	}
}
