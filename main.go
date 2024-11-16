package main

import (
	"Twilio-Sms-Sender/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	logrus.SetOutput(multiWriter)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetLevel(logrus.InfoLevel)

	server := gin.Default()

	router.SetupRouter(server)

	logrus.Info("Starting server on port 8000...")

	err = server.Run(":8000")

	if err != nil {
		return
	}
}
