package handler

import (
	"context"
	"fmt"
	"time"

	productpb "github.com/shashankbiet/rate-limiter/proto/product"
	"google.golang.org/grpc/status"
)

type productServer struct{}

func NewProductServer() *productServer {
	return &productServer{}
}

func (*productServer) GetProduct(ctx context.Context, req *productpb.ProductRequest) (*productpb.ProductResponse, error) {
	if req.Id <= 0 {
		return nil, status.Error(400, "bad request!")
	}
	time.Sleep(100 * time.Millisecond)
	return &productpb.ProductResponse{
		Product: &productpb.Product{
			Name: fmt.Sprintf("Product: %v", req.Id),
		},
	}, nil
}
