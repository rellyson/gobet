package main

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/rellyson/gobet/pkg/http"
)

func main() {
	workDir, _ := os.Getwd()
	godotenv.Load(path.Join(workDir, "config", ".env"))
	addr := os.Getenv("APP_ADDRESS")

	log.Fatal(http.CreateServer().Start(addr))
}
