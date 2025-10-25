package main

import (
	"github.com/rama-x2/andora-main/internal/logger"
	"github.com/rama-x2/andora-main/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.SetupLogger()
	err := server.StartServer()
	if err != nil {
		logrus.Fatalf("Failed start server: %v", err)
	}
}
