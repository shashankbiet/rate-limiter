package httpserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shashankbiet/rate-limiter/app/constants"
	"github.com/shashankbiet/rate-limiter/app/handler"
	"github.com/shashankbiet/rate-limiter/app/middleware"
	ratelimiter "github.com/shashankbiet/rate-limiter/app/middleware/rate-limiter"
	"github.com/shashankbiet/rate-limiter/pkg/config"
	"github.com/shashankbiet/rate-limiter/pkg/logger"
)

func InitHttpServer() (*http.Server, error) {
	userApiRateLimiter := ratelimiter.NewTokenBucket(30*time.Second, 4) // tokens are added to the bucket every 30 seconds (or twice per minute)
	userHandler := handler.NewUserHandler()
	router := mux.NewRouter()
	router.HandleFunc(constants.USER_PATH, func(w http.ResponseWriter, r *http.Request) {
		if userApiRateLimiter.Allow() {
			userHandler.GetUserDetails(w, r)
		} else {
			http.Error(w, "rate limit exceeded!", http.StatusTooManyRequests)
		}
	}).Methods("GET")
	http.Handle("/", router)
	router.Use(middleware.HttpRequestLogger)

	port := fmt.Sprintf(":%v", config.GetConfig().HttpServer.Port)
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to listen", "error", err)
		}
	}()

	return server, nil
}
