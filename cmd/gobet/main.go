package main

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/rellyson/gobet/pkg/amqp"
	"github.com/rellyson/gobet/pkg/http"
)

func main() {
	rootPath, _ := os.Getwd()
	godotenv.Load(path.Join(rootPath, "config", ".env"))
	amqpUrl := os.Getenv("RABBITMQ_URL")
	addr := os.Getenv("APP_ADDRESS")

	amqp.CreateConnection(amqpUrl)
	log.Fatal(http.CreateServer().Start(addr))
}
