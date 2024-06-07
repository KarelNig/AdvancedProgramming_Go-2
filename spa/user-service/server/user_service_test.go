package main

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"os"
	"testing"

	pb "user-service"
)

var testUserServer *server
var userMock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mck, _ := sqlmock.New()
	userMock = mck
	defer db.Close()
	testUserServer = &server{db: db}
	os.Exit(m.Run())
}

func TestGetUser(t *testing.T) {
	userMock.ExpectQuery("SELECT id, name, email, phone, address FROM users_create WHERE id=\\$1").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "address"}).AddRow("1", "John Doe", "john@example.com", "1234567890", "123 Street"))

	req := &pb.GetUserRequest{
		UserId: "1",
	}

	res, err := testUserServer.GetUser(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.User.UserId != "1" {
		t.Fatalf("Expected user ID to be '1', got %v", res.User.UserId)
	}
}

func TestCreateUser(t *testing.T) {
	userMock.ExpectQuery("INSERT INTO users_create \\(name, email, phone, address\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\) RETURNING id").
		WithArgs("John Doe", "john@example.com", "1234567890", "123 Street").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	req := &pb.CreateUserRequest{
		Name:    "John Doe",
		Email:   "john@example.com",
		Phone:   "1234567890",
		Address: "123 Street",
	}

	res, err := testUserServer.CreateUser(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.UserId != "1" {
		t.Fatalf("Expected user ID to be '1', got %v", res.UserId)
	}
}

func TestUpdateUser(t *testing.T) {
	userMock.ExpectExec("UPDATE users_create SET name=\\$1, email=\\$2, phone=\\$3, address=\\$4 WHERE id=\\$5").
		WithArgs("John Doe", "john@example.com", "1234567890", "123 Street", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	userMock.ExpectQuery("SELECT id, name, email, phone, address FROM users_create WHERE id=\\$1").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "address"}).AddRow("1", "John Doe", "john@example.com", "1234567890", "123 Street"))

	req := &pb.UpdateUserRequest{
		UserId:  "1",
		Name:    "John Doe",
		Email:   "john@example.com",
		Phone:   "1234567890",
		Address: "123 Street",
	}

	res, err := testUserServer.UpdateUser(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.User.UserId != "1" {
		t.Fatalf("Expected user ID to be '1', got %v", res.User.UserId)
	}
}

func TestDeleteUser(t *testing.T) {
	userMock.ExpectExec("DELETE FROM users_create WHERE id=\\$1").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.DeleteUserRequest{
		UserId: "1",
	}

	res, err := testUserServer.DeleteUser(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !res.Success {
		t.Fatalf("Expected success to be true")
	}
}

func TestListUsers(t *testing.T) {
	userMock.ExpectQuery("SELECT id, name, email, phone, address FROM users_create WHERE \\(name ILIKE \\$1 OR email ILIKE \\$1 OR phone ILIKE \\$1 OR address ILIKE \\$1\\) ORDER BY id LIMIT 10 OFFSET 0").
		WithArgs("%test%").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "address"}).
			AddRow("1", "John Doe", "john@example.com", "1234567890", "123 Street"))

	userMock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users_create WHERE \\(name ILIKE \\$1 OR email ILIKE \\$1 OR phone ILIKE \\$1 OR address ILIKE \\$1\\)").
		WithArgs("%test%").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	req := &pb.ListUsersRequest{
		Filter:   "test",
		Page:     1,
		PageSize: 10,
	}

	res, err := testUserServer.ListUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(res.Users) != 1 {
		t.Fatalf("Expected 1 user, got %v", len(res.Users))
	}
}
