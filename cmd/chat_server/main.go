package main

import (
	config2 "github.com/ewik2k21/grpc_simple_chat/config"
	"github.com/ewik2k21/grpc_simple_chat/internal/api/grpc/chat_server"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	config2.LoadEnv()
	server := chat_server.NewServer(os.Getenv(config2.GRPCPort1))

	if err := server.Run(); err != nil {
		logrus.Fatalf("failed run server: %s", err)
	}
}
