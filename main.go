package main

import (
	"context"

	"github.com/shashankbiet/rate-limiter/app/initializer"
)

func main() {
	ctx := context.Background()
	initializer.InitializerConfig()
	initializer.InitializeLogger()
	initializer.InitializeServer(ctx)
}
