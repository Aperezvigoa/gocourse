package intermediate

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()

	// Set log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// Loggin example
	log.Info("This is info message.")
	log.Warn("This is warning message.")
	log.Error("This is error message.")

	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User logged in.")
}
