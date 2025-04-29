package intermediate

import (
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Failed to create zap logger")
	}

	defer logger.Sync()

	logger.Info("This is an info message")

	logger.Info("User logged in", zap.String("username", "John"), zap.String("method", "GET"))
}
