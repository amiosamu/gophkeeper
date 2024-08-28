package services

import (
	"context"

	querypb "github.com/amiosamu/gophkeeper/api-gateway/pkg/query/pb"
	servicespb "github.com/amiosamu/gophkeeper/api-gateway/pkg/services/pb"
	"github.com/amiosamu/gophkeeper/query-service/internal/db"
	"github.com/amiosamu/gophkeeper/query-service/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QueryServer struct {
	querypb.UnimplementedQueryServiceServer
	storage db.Storage
}

var _ querypb.QueryServiceServer = (*QueryServer)(nil)

func NewQueryServer(storage db.Storage) *QueryServer {
	return &QueryServer{storage: storage}
}

func (s *QueryServer) Query(ctx context.Context, in *querypb.QueryRequest) (*servicespb.QueryResponseArray, error) {
	var record *models.Record
	if in.Type == querypb.MessageType_ANY {
		record = &models.Record{UserId: in.UserID}
	} else {
		record = &models.Record{UserId: in.UserID, MessageType: byte(in.Type)}
	}

	rows, err := s.storage.GetRecord(record)

	if err != nil {
		return &servicespb.QueryResponseArray{
			Count: 0,
			Items: nil,
		}, status.Errorf(codes.Internal, err.Error())
	}

	count := len(*rows)

	items := make([]*servicespb.QueryResponseArray_QueryResponse, 0, count)

	for _, v := range *rows {
		items = append(items, &servicespb.QueryResponseArray_QueryResponse{
			Id:   v.Id,
			Type: servicespb.MessageType(v.MessageType),
			Data: v.Data,
			Meta: v.Meta,
		})
	}

	return &servicespb.QueryResponseArray{
		Count: int64(count),
		Items: items,
	}, nil
}
