package main

import (
	"context"
	"os"
	"testing"

	pb "auth"
	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/bcrypt"
)

var testServer *server
var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mck, _ := sqlmock.New()
	mock = mck
	defer db.Close()
	testServer = &server{db: db}
	os.Exit(m.Run())
}

func TestLogin(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	mock.ExpectQuery("SELECT id, password FROM users WHERE username=\\$1").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).AddRow("1", string(hashedPassword)))

	req := &pb.LoginRequest{
		Username: "testuser",
		Password: "password",
	}

	res, err := testServer.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.UserId != "1" {
		t.Fatalf("Expected user ID to be '1', got %v", res.UserId)
	}
}

func TestRegister(t *testing.T) {
	mock.ExpectQuery("INSERT INTO users \\(username, email, password\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id").
		WithArgs("newuser", "newuser@example.com", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	req := &pb.RegisterRequest{
		Username: "newuser",
		Email:    "newuser@example.com",
		Password: "password",
	}

	res, err := testServer.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.UserId != "1" {
		t.Fatalf("Expected user ID to be '1', got %v", res.UserId)
	}
}

func TestValidateToken(t *testing.T) {
	token, _ := generateJWTToken("1")

	req := &pb.TokenRequest{
		Token: token,
	}

	res, err := testServer.ValidateToken(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !res.Valid {
		t.Fatalf("Expected token to be valid")
	}

	if res.UserId != "1" {
		t.Fatalf("Expected user ID to be '1', got %v", res.UserId)
	}
}

func TestRefreshToken(t *testing.T) {
	token, _ := generateJWTToken("1")

	req := &pb.TokenRequest{
		Token: token,
	}

	res, err := testServer.RefreshToken(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !res.Valid {
		t.Fatalf("Expected token to be valid")
	}

	if res.UserId != "1" {
		t.Fatalf("Expected user ID to be '1', got %v", res.UserId)
	}

	if res.NewToken == "" {
		t.Fatalf("Expected new token, got empty string")
	}
}
