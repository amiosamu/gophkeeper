package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/auth/pb"
)

type MiddlewareInterceptor struct {
	*ServiceClient
}

func NewAuthMiddlewareInterceptor(client *ServiceClient) *MiddlewareInterceptor {
	return &MiddlewareInterceptor{client}
}

func (inter *MiddlewareInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		cont, err := inter.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(cont, req)
	}
}

func (inter *MiddlewareInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	if method == "/services.APIGatewayService/Register" ||
		method == "/services.APIGatewayService/Login" {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := strings.Split(values[0], "Bearer ")
	if len(token) < 2 {
		return ctx, status.Errorf(codes.Unauthenticated, "bad authrization token")
	}

	res, err := inter.Client.Verify(ctx, &pb.VerifyRequest{
		Token: token[1],
	})
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	ctx = context.WithValue(ctx, KeyPrincipalID, res.UserID)
	return ctx, nil
}
