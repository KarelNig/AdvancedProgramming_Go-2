package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	pb "news"
)

type server struct {
	pb.UnimplementedNewsServiceServer
	db *sql.DB
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func (s *server) GetNews(ctx context.Context, req *pb.GetNewsRequest) (*pb.GetNewsResponse, error) {
	var news pb.News
	err := s.db.QueryRowContext(ctx, "SELECT id, title, content FROM news WHERE id=$1", req.NewsId).
		Scan(&news.NewsId, &news.Title, &news.Content)
	if err != nil {
		log.Error().Err(err).Str("news_id", req.NewsId).Msg("Failed to get news")
		return nil, err
	}
	log.Info().Str("news_id", req.NewsId).Msg("GetNews called")
	return &pb.GetNewsResponse{News: &news}, nil
}

func (s *server) CreateNews(ctx context.Context, req *pb.CreateNewsRequest) (*pb.CreateNewsResponse, error) {
	var newsID string
	err := s.db.QueryRowContext(ctx,
		"INSERT INTO news (title, content) VALUES ($1, $2) RETURNING id",
		req.Title, req.Content).Scan(&newsID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create news")
		return nil, err
	}
	log.Info().Str("news_id", newsID).Msg("CreateNews called")
	return &pb.CreateNewsResponse{NewsId: newsID}, nil
}

func (s *server) UpdateNews(ctx context.Context, req *pb.UpdateNewsRequest) (*pb.UpdateNewsResponse, error) {
	_, err := s.db.ExecContext(ctx,
		"UPDATE news SET title=$1, content=$2 WHERE id=$3",
		req.Title, req.Content, req.NewsId)
	if err != nil {
		log.Error().Err(err).Str("news_id", req.NewsId).Msg("Failed to update news")
		return nil, err
	}

	var news pb.News
	err = s.db.QueryRowContext(ctx, "SELECT id, title, content FROM news WHERE id=$1", req.NewsId).
		Scan(&news.NewsId, &news.Title, &news.Content)
	if err != nil {
		log.Error().Err(err).Str("news_id", req.NewsId).Msg("Failed to retrieve updated news")
		return nil, err
	}
	log.Info().Str("news_id", req.NewsId).Msg("UpdateNews called")
	return &pb.UpdateNewsResponse{News: &news}, nil
}

func (s *server) DeleteNews(ctx context.Context, req *pb.DeleteNewsRequest) (*pb.DeleteNewsResponse, error) {
	_, err := s.db.ExecContext(ctx, "DELETE FROM news WHERE id=$1", req.NewsId)
	if err != nil {
		log.Error().Err(err).Str("news_id", req.NewsId).Msg("Failed to delete news")
		return nil, err
	}
	log.Info().Str("news_id", req.NewsId).Msg("DeleteNews called")
	return &pb.DeleteNewsResponse{Success: true}, nil
}

func (s *server) ListNews(ctx context.Context, req *pb.ListNewsRequest) (*pb.ListNewsResponse, error) {
	query := "SELECT id, title, content FROM news"
	var conditions []string
	var args []interface{}

	// Filtering
	if req.Filter != "" {
		conditions = append(conditions, "(title ILIKE $1 OR content ILIKE $1)")
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
		log.Error().Err(err).Msg("Failed to list news")
		return nil, err
	}
	defer rows.Close()

	var newsList []*pb.News
	for rows.Next() {
		var news pb.News
		if err := rows.Scan(&news.NewsId, &news.Title, &news.Content); err != nil {
			log.Error().Err(err).Msg("Failed to scan news row")
			return nil, err
		}
		newsList = append(newsList, &news)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over news rows")
		return nil, err
	}

	var totalCount int32
	countQuery := "SELECT COUNT(*) FROM news"
	if len(conditions) > 0 {
		countQuery += " WHERE " + strings.Join(conditions, " AND ")
	}
	err = s.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		log.Error().Err(err).Msg("Failed to count news rows")
		return nil, err
	}

	totalPages := (totalCount + pageSize - 1) / pageSize

	log.Info().Msg("ListNews called")
	return &pb.ListNewsResponse{
		NewsList:    newsList,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}

func (s *server) SendNewsToUsers(ctx context.Context, req *pb.SendNewsToUsersRequest) (*pb.SendNewsToUsersResponse, error) {
	var title, content string
	err := s.db.QueryRowContext(ctx, "SELECT title, content FROM news WHERE id=$1", req.NewsId).
		Scan(&title, &content)
	if err != nil {
		log.Error().Err(err).Str("news_id", req.NewsId).Msg("Failed to get news for sending to users")
		return nil, err
	}

	rows, err := s.db.Query("SELECT email FROM users")
	if err != nil {
		log.Error().Err(err).Msg("Failed to list user emails")
		return nil, err
	}
	defer rows.Close()

	var emails []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			log.Error().Err(err).Msg("Failed to scan email row")
			return nil, err
		}
		emails = append(emails, email)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over email rows")
		return nil, err
	}

	for _, email := range emails {
		fmt.Printf("Sending news to %s: %s - %s\n", email, title, content)
		log.Info().Str("email", email).Msg("Sending news to user")
	}

	log.Info().Str("news_id", req.NewsId).Msg("SendNewsToUsers called")
	return &pb.SendNewsToUsersResponse{Success: true}, nil
}
