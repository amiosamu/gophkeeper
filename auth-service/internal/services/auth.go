package services

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/amiosamu/gophkeeper/auth-service/internal/db"
	"github.com/amiosamu/gophkeeper/auth-service/internal/models"
	"github.com/amiosamu/gophkeeper/auth-service/internal/pb"
	"github.com/amiosamu/gophkeeper/auth-service/internal/utils"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	jwt     *utils.JwtWraper
	storage db.Storage
}

func NewAuthServer(jwt *utils.JwtWraper, storage db.Storage) *AuthServer {
	return &AuthServer{
		jwt:     jwt,
		storage: storage,
	}
}

func (s *AuthServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := s.storage.AddUser(&models.User{
		Name:     in.UserName,
		Password: in.Password,
	})
	if err != nil {
		if errors.Is(err, db.ErrorUserAlreadyExist) {
			return &pb.RegisterResponse{
				Token:  "",
				UserID: 0,
			}, status.Errorf(codes.AlreadyExists, db.ErrorUserAlreadyExist.Error())
		}
		return &pb.RegisterResponse{
			Token:  "",
			UserID: 0,
		}, status.Errorf(codes.Internal, err.Error())
	}

	token, err := s.jwt.GenerateToken(user)

	if err != nil {
		return &pb.RegisterResponse{
			Token:  "",
			UserID: 0,
		}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.RegisterResponse{
		Token:  token,
		UserID: user.Id,
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.storage.GetUser(&models.User{Name: in.UserName})

	if err != nil {
		return &pb.LoginResponse{
			Token:  "",
			UserID: 0,
		}, status.Errorf(codes.NotFound, err.Error())
	}

	match := utils.CheckPasswordHash(in.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Token:  "",
			UserID: 0,
		}, status.Errorf(codes.NotFound, db.ErrorUserNotFound.Error())
	}

	token, err := s.jwt.GenerateToken(user)

	if err != nil {
		return &pb.LoginResponse{
			Token:  "",
			UserID: 0,
		}, status.Errorf(codes.Internal, "generate token error:", err.Error())
	}

	return &pb.LoginResponse{
		Token:  token,
		UserID: user.Id,
	}, nil
}

func (s *AuthServer) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	claims, err := s.jwt.ValidateToken(in.Token)

	if err != nil {
		return &pb.VerifyResponse{}, status.Errorf(codes.Internal, "bad request", err.Error())
	}

	return &pb.VerifyResponse{
		UserID: claims.IdUser,
	}, nil

}
