package services

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	commandpb "github.com/amiosamu/gophkeeper/api-gateway/pkg/command/pb"
)

type CommandServer struct {
	commandpb.UnimplementedCommandServiceServer
	comm commandpb.CommandServiceClient
}

var _ commandpb.CommandServiceServer = (*CommandServer)(nil)

func NewCommandServer(commandClient commandpb.CommandServiceClient) *CommandServer {
	return &CommandServer{
		comm: commandClient,
	}
}

func (s *CommandServer) AddCommand(ctx context.Context, in *commandpb.CommandRequest) (*commandpb.CommandResponse, error) {
	// тут отсылаем в command consumer либо через kafka либо по grpc
	// при отсылке в кафка формируем очередь от hash(userID)
	_, err := s.comm.AddCommand(ctx, &commandpb.CommandRequest{
		Id:     in.Id,
		Type:   in.Type,
		UserID: in.UserID,
		Data:   in.Data,
		Meta:   in.Meta,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "add command error: %v", err)
	}

	return &commandpb.CommandResponse{}, nil
}

func (s *CommandServer) ModifyCommand(ctx context.Context, in *commandpb.CommandRequest) (*commandpb.CommandResponse, error) {
	// тут отсылаем в command consumer либо через kafka либо по grpc
	// при отсылке в кафка формируем очередь от hash(userID)
	_, err := s.comm.ModifyCommand(ctx, &commandpb.CommandRequest{
		Id:     in.Id,
		Type:   in.Type,
		UserID: in.UserID,
		Data:   in.Data,
		Meta:   in.Meta,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "modify command error: %v", err)
	}

	return &commandpb.CommandResponse{}, nil
}
func (s *CommandServer) DeleteCommand(ctx context.Context, in *commandpb.CommandRequest) (*commandpb.CommandResponse, error) {
	// тут отсылаем в command consumer либо через kafka либо по grpc
	// при отсылке в кафка формируем очередь от hash(userID)
	_, err := s.comm.AddCommand(ctx, &commandpb.CommandRequest{
		Id:     in.Id,
		Type:   in.Type,
		UserID: in.UserID,
		Data:   in.Data,
		Meta:   in.Meta,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete command error: %v", err)
	}

	return &commandpb.CommandResponse{}, nil
}
