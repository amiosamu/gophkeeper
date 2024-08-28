package services

import (
	"context"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/command/pb"
	"github.com/amiosamu/gophkeeper/command-consumer-service/internal/db"
	"github.com/amiosamu/gophkeeper/command-consumer-service/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommandServer struct {
	pb.UnimplementedCommandServiceServer
	storage db.Storage
}

var _ pb.CommandServiceServer = (*CommandServer)(nil)

func NewCommandServer(storage db.Storage) *CommandServer {
	return &CommandServer{
		storage: storage,
	}
}

func (s *CommandServer) AddCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error) {
	err := s.storage.AddRecord(&models.Record{
		UserId:      in.UserID,
		MessageType: byte(in.Type),
		Data:        in.Data,
		Meta:        in.Meta,
	})

	if err != nil {
		return &pb.CommandResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CommandResponse{}, nil
}

func (s *CommandServer) ModifyCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error) {
	err := s.storage.ModifyRecord(&models.Record{
		Id:          in.Id,
		UserId:      in.UserID,
		MessageType: byte(in.Type),
		Data:        in.Data,
		Meta:        in.Meta,
	})

	if err != nil {
		return &pb.CommandResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.CommandResponse{}, nil
}

func (s *CommandServer) DeleteCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error) {
	err := s.storage.DeleteRecord(&models.Record{
		Id: in.Id,
	})

	if err != nil {
		return &pb.CommandResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.CommandResponse{}, nil
}
