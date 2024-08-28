package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/services/pb"
	"github.com/amiosamu/gophkeeper/client/internal/config"
)

type key uint64

const KeyPrincipalID key = iota

type ServiceClient struct {
	Client pb.APIGatewayServiceClient
}

func NewServiceClient() *ServiceClient {
	return &ServiceClient{}
}

func (s *ServiceClient) Dial(ctx context.Context, cfg *config.Config) pb.APIGatewayServiceClient {
	cc, err := grpc.DialContext(ctx, cfg.ServerAddres,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unary()))

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	s.Client = pb.NewAPIGatewayServiceClient(cc)

	return s.Client
}

func unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if method == "/services.APIGatewayService/Login" ||
			method != "/services.APIGatewayService/Register" {
			invoker(ctx, method, req, reply, cc)
		}
		return invoker(attachToken(ctx), method, req, reply, cc)

	}
}

func attachToken(ctx context.Context) context.Context {
	val, ok := ctx.Value(KeyPrincipalID).(string)
	if !ok {
		return ctx
	}
	md := metadata.Pairs()
	md.Append("authorization", fmt.Sprint("Bearer ", val))
	return metadata.NewOutgoingContext(ctx, md)
}
