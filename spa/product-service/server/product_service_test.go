package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	pb "product"
)

var testProductServer *server
var productMock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mck, _ := sqlmock.New()
	productMock = mck
	defer db.Close()
	testProductServer = &server{db: db}
	os.Exit(m.Run())
}

func TestGetBooking(t *testing.T) {
	productMock.ExpectQuery("SELECT id, user_id, spa_service, booking_time FROM bookings WHERE id=\\$1").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "spa_service", "booking_time"}).AddRow("1", "123", "massage", time.Now()))

	req := &pb.GetBookingRequest{
		BookingId: "1",
	}

	res, err := testProductServer.GetBooking(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.Booking.BookingId != "1" {
		t.Fatalf("Expected booking ID to be '1', got %v", res.Booking.BookingId)
	}
}

func TestCreateBooking(t *testing.T) {
	productMock.ExpectQuery("INSERT INTO bookings \\(user_id, spa_service, booking_time\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id").
		WithArgs("123", "massage", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	req := &pb.CreateBookingRequest{
		UserId:      "123",
		SpaService:  "massage",
		BookingTime: time.Now().Format(time.RFC3339),
	}

	res, err := testProductServer.CreateBooking(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.BookingId != "1" {
		t.Fatalf("Expected booking ID to be '1', got %v", res.BookingId)
	}
}

func TestUpdateBooking(t *testing.T) {
	productMock.ExpectExec("UPDATE bookings SET user_id=\\$1, spa_service=\\$2, booking_time=\\$3 WHERE id=\\$4").
		WithArgs("123", "massage", sqlmock.AnyArg(), "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	productMock.ExpectQuery("SELECT id, user_id, spa_service, booking_time FROM bookings WHERE id=\\$1").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "spa_service", "booking_time"}).AddRow("1", "123", "massage", time.Now()))

	req := &pb.UpdateBookingRequest{
		BookingId:   "1",
		UserId:      "123",
		SpaService:  "massage",
		BookingTime: time.Now().Format(time.RFC3339),
	}

	res, err := testProductServer.UpdateBooking(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.Booking.BookingId != "1" {
		t.Fatalf("Expected booking ID to be '1', got %v", res.Booking.BookingId)
	}
}

func TestDeleteBooking(t *testing.T) {
	productMock.ExpectExec("DELETE FROM bookings WHERE id=\\$1").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.DeleteBookingRequest{
		BookingId: "1",
	}

	res, err := testProductServer.DeleteBooking(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !res.Success {
		t.Fatalf("Expected success to be true")
	}
}

func TestListBookings(t *testing.T) {
	productMock.ExpectQuery("SELECT id, user_id, spa_service, booking_time FROM bookings WHERE \\(spa_service ILIKE \\$1\\) AND user_id = \\$2 ORDER BY id LIMIT 10 OFFSET 0").
		WithArgs("%test%", "123").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "spa_service", "booking_time"}).
			AddRow("1", "123", "massage", time.Now()))

	productMock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM bookings WHERE \\(spa_service ILIKE \\$1\\) AND user_id = \\$2").
		WithArgs("%test%", "123").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	req := &pb.ListBookingsRequest{
		Filter:   "test",
		UserId:   "123",
		Page:     1,
		PageSize: 10,
	}

	res, err := testProductServer.ListBookings(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(res.Bookings) != 1 {
		t.Fatalf("Expected 1 booking, got %v", len(res.Bookings))
	}
}
