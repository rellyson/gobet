package http

import (
	"crypto/tls"
	"log"
	"net"
	"os"
	"path"

	_ "github.com/rellyson/gobet/api/openapi"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/rellyson/gobet/pkg/http/handlers"
	"github.com/rellyson/gobet/pkg/http/middlewares"
)

// @title           GoBet
// @version         1.0.0
// @description     GoBet - A sports Bet platform.

// @host      127.0.0.1:3000
// @BasePath  /api
// @accept json
// @produce json
type Server struct {
	app *fiber.App
}

func CreateServer() *Server {
	s := new(Server)

	s.app = fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})

	s.app.Use(recover.New())

	s.app.Use(requestid.New())

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	s.app.Use(logger.New(logger.Config{
		Format: "[${locals:requestid}] ${ip}:${port} ${status} - ${method} ${path}\n",
	}))

	s.app.Get("/api/metrics", monitor.New(monitor.Config{
		APIOnly: true,
	}))

	s.app.Get("/api/healthcheck", middlewares.HealtcheckMiddleware)

	s.app.Get("/api/docs/*", swagger.New(swagger.Config{ // custom
		DeepLinking:            false,
		DocExpansion:           "none",
		TryItOutEnabled:        true,
		RequestSnippetsEnabled: true,
		WithCredentials:        true,
	}))

	return s
}

func (s *Server) Start(addr string) error {
	workDir, _ := os.Getwd()
	ln, _ := net.Listen("tcp4", addr)

	cert, err := tls.LoadX509KeyPair(
		path.Join(workDir, "config", "certs", "certificate.crt"),
		path.Join(workDir, "config", "certs", "cert-key.key"),
	)

	if err != nil {
		log.Fatalf("There was an error loading TLS certificate: %s", err)
	}

	ln = tls.NewListener(ln, &tls.Config{
		Certificates: []tls.Certificate{cert},
	})

	return s.app.Listener(ln)
}
