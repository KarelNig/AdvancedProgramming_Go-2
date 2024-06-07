package main

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "auth"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_spa")

type server struct {
	pb.UnimplementedAuthServiceServer
	db *sql.DB
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var userID, hashedPassword string
	err := s.db.QueryRowContext(ctx, "SELECT id, password FROM users WHERE username=$1", req.Username).Scan(&userID, &hashedPassword)
	if err != nil {
		log.Error().Err(err).Str("username", req.Username).Msg("Failed to query user")
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		log.Error().Err(err).Msg("Invalid password")
		return nil, err
	}

	token, err := generateJWTToken(userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate token")
		return nil, err
	}

	refreshToken, err := generateJWTToken(userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate refresh token")
		return nil, err
	}

	log.Info().Str("user_id", userID).Msg("User logged in")
	return &pb.LoginResponse{Token: token, RefreshToken: refreshToken, UserId: userID}, nil
}

func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Failed to hash password")
		return nil, err
	}

	var userID string
	err = s.db.QueryRowContext(ctx, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", req.Username, req.Email, hashedPassword).Scan(&userID)
	if err != nil {
		log.Error().Err(err).Str("username", req.Username).Msg("Failed to insert user")
		return nil, err
	}

	log.Info().Str("user_id", userID).Msg("User registered")
	return &pb.RegisterResponse{UserId: userID}, nil
}

func (s *server) ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	claims, err := validateJWTToken(req.Token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate token")
		return &pb.TokenResponse{Valid: false}, err
	}

	log.Info().Str("user_id", claims.UserID).Msg("Token validated")
	return &pb.TokenResponse{Valid: true, UserId: claims.UserID}, nil
}

func (s *server) RefreshToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	claims, err := validateJWTToken(req.Token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate token for refresh")
		return &pb.TokenResponse{Valid: false}, err
	}

	newToken, err := generateJWTToken(claims.UserID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate new token")
		return nil, err
	}

	log.Info().Str("user_id", claims.UserID).Msg("Token refreshed")
	return &pb.TokenResponse{Valid: true, UserId: claims.UserID, NewToken: newToken}, nil
}

func generateJWTToken(userID string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Error().Err(err).Msg("Failed to sign token")
		return "", err
	}
	return tokenString, nil
}

func validateJWTToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to parse token")
		return nil, err
	}

	if !token.Valid {
		log.Error().Msg("Invalid token")
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
