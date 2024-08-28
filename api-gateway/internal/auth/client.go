package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/amiosamu/gophkeeper/api-gateway/internal/config"
	"github.com/amiosamu/gophkeeper/api-gateway/pkg/auth/pb"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

type key uint64

const (
	KeyPrincipalID key = iota
)

func NewAuthServiceClient(ctx context.Context, cfg *config.Config) *ServiceClient {
	return &ServiceClient{
		Client: InitAuthServiceClient(ctx, cfg),
	}
}

func InitAuthServiceClient(ctx context.Context, cfg *config.Config) pb.AuthServiceClient {
	cc, err := grpc.DialContext(ctx, cfg.AuthServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
