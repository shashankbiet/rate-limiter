package middleware

import (
	"context"

	"github.com/shashankbiet/rate-limiter/app/constants"
	ratelimiter "github.com/shashankbiet/rate-limiter/app/middleware/rate-limiter"
	"github.com/shashankbiet/rate-limiter/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Perform actions before the RPC call.
	logger.Info("GRPC Request", "method", info.FullMethod, "body", req)

	// Call the next interceptor/handler in the chain.
	resp, err := handler(ctx, req)

	// Perform actions after the RPC call.

	return resp, err
}

func RateLimiterInterceptor(limiter ratelimiter.IRateLimiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		switch info.FullMethod {
		case constants.PRODUCT_METHOD:
			if limiter.Allow() {
				return handler(ctx, req)
			} else {
				return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded!")
			}
		}

		// Default case: allow the request
		return handler(ctx, req)
	}
}
