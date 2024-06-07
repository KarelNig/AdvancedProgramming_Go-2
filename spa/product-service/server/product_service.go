package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	pb "product"
)

type server struct {
	pb.UnimplementedProductServiceServer
	db *sql.DB
}

func (s *server) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	var booking pb.Booking
	err := s.db.QueryRowContext(ctx, "SELECT id, user_id, spa_service, booking_time FROM bookings WHERE id=$1", req.BookingId).
		Scan(&booking.BookingId, &booking.UserId, &booking.SpaService, &booking.BookingTime)
	if err != nil {
		log.Error().Err(err).Str("booking_id", req.BookingId).Msg("Failed to get booking")
		return nil, err
	}
	log.Info().Str("booking_id", req.BookingId).Msg("GetBooking called")
	return &pb.GetBookingResponse{Booking: &booking}, nil
}

func (s *server) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	var bookingID string
	bookingTime, err := time.Parse(time.RFC3339, req.BookingTime)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse booking time")
		return nil, err
	}

	err = s.db.QueryRowContext(ctx,
		"INSERT INTO bookings (user_id, spa_service, booking_time) VALUES ($1, $2, $3) RETURNING id",
		req.UserId, req.SpaService, bookingTime).Scan(&bookingID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create booking")
		return nil, err
	}
	log.Info().Str("booking_id", bookingID).Msg("CreateBooking called")
	return &pb.CreateBookingResponse{BookingId: bookingID}, nil
}

func (s *server) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	bookingTime, err := time.Parse(time.RFC3339, req.BookingTime)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse booking time")
		return nil, err
	}

	_, err = s.db.ExecContext(ctx,
		"UPDATE bookings SET user_id=$1, spa_service=$2, booking_time=$3 WHERE id=$4",
		req.UserId, req.SpaService, bookingTime, req.BookingId)
	if err != nil {
		log.Error().Err(err).Str("booking_id", req.BookingId).Msg("Failed to update booking")
		return nil, err
	}

	var booking pb.Booking
	err = s.db.QueryRowContext(ctx, "SELECT id, user_id, spa_service, booking_time FROM bookings WHERE id=$1", req.BookingId).
		Scan(&booking.BookingId, &booking.UserId, &booking.SpaService, &booking.BookingTime)
	if err != nil {
		log.Error().Err(err).Str("booking_id", req.BookingId).Msg("Failed to retrieve updated booking")
		return nil, err
	}
	log.Info().Str("booking_id", req.BookingId).Msg("UpdateBooking called")
	return &pb.UpdateBookingResponse{Booking: &booking}, nil
}

func (s *server) DeleteBooking(ctx context.Context, req *pb.DeleteBookingRequest) (*pb.DeleteBookingResponse, error) {
	_, err := s.db.ExecContext(ctx, "DELETE FROM bookings WHERE id=$1", req.BookingId)
	if err != nil {
		log.Error().Err(err).Str("booking_id", req.BookingId).Msg("Failed to delete booking")
		return nil, err
	}
	log.Info().Str("booking_id", req.BookingId).Msg("DeleteBooking called")
	return &pb.DeleteBookingResponse{Success: true}, nil
}

func (s *server) ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	query := "SELECT id, user_id, spa_service, booking_time FROM bookings"
	var conditions []string
	var args []interface{}

	if req.Filter != "" {
		conditions = append(conditions, "(spa_service ILIKE $1)")
		args = append(args, fmt.Sprintf("%%%s%%", req.Filter))
	}

	if req.UserId != "" {
		conditions = append(conditions, "user_id = $2")
		args = append(args, req.UserId)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if req.Sort != "" {
		query += fmt.Sprintf(" ORDER BY %s", req.Sort)
	} else {
		query += " ORDER BY id"
	}

	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list bookings")
		return nil, err
	}
	defer rows.Close()

	var bookings []*pb.Booking
	for rows.Next() {
		var booking pb.Booking
		if err := rows.Scan(&booking.BookingId, &booking.UserId, &booking.SpaService, &booking.BookingTime); err != nil {
			log.Error().Err(err).Msg("Failed to scan booking row")
			return nil, err
		}
		bookings = append(bookings, &booking)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over booking rows")
		return nil, err
	}

	var totalCount int32
	countQuery := "SELECT COUNT(*) FROM bookings"
	if len(conditions) > 0 {
		countQuery += " WHERE " + strings.Join(conditions, " AND ")
	}
	err = s.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		log.Error().Err(err).Msg("Failed to count bookings")
		return nil, err
	}

	totalPages := (totalCount + pageSize - 1) / pageSize

	log.Info().Msg("ListBookings called")
	return &pb.ListBookingsResponse{
		Bookings:    bookings,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}
