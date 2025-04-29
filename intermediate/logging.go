package intermediate

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message")

	log.SetPrefix("INFO:")
	log.Println("This is a info message")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message")

	// infoLogger.Println("This is a info message")
	// warnLogger.Println("This is a warning message")
	// errorLogger.Println("This is a error message")

	file, err := os.OpenFile("logging.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message")
	infoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger.Println("This is a info message")
	warnLogger := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger.Println("This is a warning message")
	errorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger.Println("This is a error message")
}
