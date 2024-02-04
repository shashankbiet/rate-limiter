package grpcserver

import (
	"fmt"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/shashankbiet/rate-limiter/app/handler"
	"github.com/shashankbiet/rate-limiter/app/middleware"
	ratelimiter "github.com/shashankbiet/rate-limiter/app/middleware/rate-limiter"
	"github.com/shashankbiet/rate-limiter/pkg/config"
	"github.com/shashankbiet/rate-limiter/pkg/logger"
	productpb "github.com/shashankbiet/rate-limiter/proto/product"
	"google.golang.org/grpc"
)

func InitGrpcServer() (*grpc.Server, error) {
	productRateLimiter := ratelimiter.NewTokenBucket(1*time.Minute, 2) // tokens are added to the bucket every minute
	port := fmt.Sprintf(":%v", config.GetConfig().GrpcServer.Port)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.LogInterceptor,
			middleware.RateLimiterInterceptor(productRateLimiter),
		)),
	)

	productpb.RegisterProductServiceServer(server, handler.NewProductServer())

	go func() {
		if err := server.Serve(listen); err != nil {
			logger.Error("failed to start grpc server", "error", err)
		}
	}()

	return server, nil
}
