package initializer

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpcserver "github.com/shashankbiet/rate-limiter/app/server/grpc"
	httpserver "github.com/shashankbiet/rate-limiter/app/server/http"
	"github.com/shashankbiet/rate-limiter/pkg/config"
	"github.com/shashankbiet/rate-limiter/pkg/logger"
	"google.golang.org/grpc"
)

func InitializerConfig() {
	config := config.GetConfig()
	fmt.Printf("config:%+v\n", config)
}

func InitializeLogger() {
	logger.InitLogger()
	logger.Info("logger setup done")
}

func InitializeServer(ctx context.Context) {
	grpcServer, err := grpcserver.InitGrpcServer()
	if err != nil {
		logger.Error("error in starting grpc server", "error", err)
	}

	httpServer, err := httpserver.InitHttpServer()
	if err != nil {
		logger.Error("error in starting http server", "error", err)
	}

	gracefulShutdown(grpcServer, httpServer)
}

func gracefulShutdown(grpcServer *grpc.Server, httpServer *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	logger.Info("graceful shut down server")
	grpcServer.GracefulStop()

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logger.Error("error in shutting down http server", "error", err)
	}
}
