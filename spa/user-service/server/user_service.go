package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	pb "user-service"
)

type server struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user pb.User
	err := s.db.QueryRowContext(ctx, "SELECT id, name, email, phone, address FROM users_create WHERE id=$1", req.UserId).
		Scan(&user.UserId, &user.Name, &user.Email, &user.Phone, &user.Address)
	if err != nil {
		log.Error().Err(err).Str("user_id", req.UserId).Msg("Failed to get user")
		return nil, err
	}
	log.Info().Str("user_id", req.UserId).Msg("GetUser called")
	return &pb.GetUserResponse{User: &user}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var userID string
	err := s.db.QueryRowContext(ctx,
		"INSERT INTO users_create (name, email, phone, address) VALUES ($1, $2, $3, $4) RETURNING id",
		req.Name, req.Email, req.Phone, req.Address).Scan(&userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user")
		return nil, err
	}
	log.Info().Str("user_id", userID).Msg("CreateUser called")
	return &pb.CreateUserResponse{UserId: userID}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	_, err := s.db.ExecContext(ctx,
		"UPDATE users_create SET name=$1, email=$2, phone=$3, address=$4 WHERE id=$5",
		req.Name, req.Email, req.Phone, req.Address, req.UserId)
	if err != nil {
		log.Error().Err(err).Str("user_id", req.UserId).Msg("Failed to update user")
		return nil, err
	}

	var user pb.User
	err = s.db.QueryRowContext(ctx, "SELECT id, name, email, phone, address FROM users_create WHERE id=$1", req.UserId).
		Scan(&user.UserId, &user.Name, &user.Email, &user.Phone, &user.Address)
	if err != nil {
		log.Error().Err(err).Str("user_id", req.UserId).Msg("Failed to retrieve updated user")
		return nil, err
	}
	log.Info().Str("user_id", req.UserId).Msg("UpdateUser called")
	return &pb.UpdateUserResponse{User: &user}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := s.db.ExecContext(ctx, "DELETE FROM users_create WHERE id=$1", req.UserId)
	if err != nil {
		log.Error().Err(err).Str("user_id", req.UserId).Msg("Failed to delete user")
		return nil, err
	}
	log.Info().Str("user_id", req.UserId).Msg("DeleteUser called")
	return &pb.DeleteUserResponse{Success: true}, nil
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	query := "SELECT id, name, email, phone, address FROM users_create"
	var conditions []string
	var args []interface{}

	if req.Filter != "" {
		conditions = append(conditions, "(name ILIKE $1 OR email ILIKE $1 OR phone ILIKE $1 OR address ILIKE $1)")
		args = append(args, fmt.Sprintf("%%%s%%", req.Filter))
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
		log.Error().Err(err).Msg("Failed to list users_create")
		return nil, err
	}
	defer rows.Close()

	var users_create []*pb.User
	for rows.Next() {
		var user pb.User
		if err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Phone, &user.Address); err != nil {
			log.Error().Err(err).Msg("Failed to scan user row")
			return nil, err
		}
		users_create = append(users_create, &user)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over user rows")
		return nil, err
	}

	var totalCount int32
	countQuery := "SELECT COUNT(*) FROM users_create"
	if len(conditions) > 0 {
		countQuery += " WHERE " + strings.Join(conditions, " AND ")
	}
	err = s.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		log.Error().Err(err).Msg("Failed to count users_create")
		return nil, err
	}

	totalPages := (totalCount + pageSize - 1) / pageSize

	log.Info().Msg("ListUsers called")
	return &pb.ListUsersResponse{
		Users:       users_create,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}
