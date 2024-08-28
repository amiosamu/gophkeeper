package client

import (
	"context"
	"log"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/command/pb"
	"github.com/amiosamu/gophkeeper/command-producer-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.CommandServiceClient
}

func InitCommandServiceClient(ctx context.Context, cfg *config.Config) pb.CommandServiceClient {
	cc, err := grpc.DialContext(ctx, cfg.CommandConsumerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect: ", err)
	}

	return pb.NewCommandServiceClient(cc)
}

func NewCommandServiceClient(ctx context.Context, cfg *config.Config) *ServiceClient {
	return &ServiceClient{
		Client: InitCommandServiceClient(ctx, cfg),
	}
}
