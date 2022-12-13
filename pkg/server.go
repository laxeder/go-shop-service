package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/config"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

// @title API Golang
// @version 1.0
// @description API Padrão para star de projeto.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
type Server struct {
	instance    *fiber.App
	Config      *fiber.Config
	DateStart   time.Time
	StageStatus string
	Address     string
	Timeout     string
	Name        string
	Port        string
	Host        string
	Env         string
	TZ          string
}

func (s *Server) New() *fiber.App {

	s.Config = config.Server()
	s.DateStart = time.Now().UTC()

	var lock = &sync.Mutex{}
	if s.instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if s.instance == nil {
			s.instance = fiber.New(*s.Config)
		}
	}

	return s.instance
}

func (s *Server) Routes(routes func(*fiber.App)) *fiber.App {
	defer routes(s.instance)
	return s.instance
}

func (s *Server) Routines(routine func(*fiber.App)) *fiber.App {
	defer routine(s.instance)
	return s.instance
}

func (s *Server) Middlewares(middleware func(*fiber.App)) *fiber.App {
	defer middleware(s.instance)
	return s.instance
}

func (s *Server) Environments() *fiber.App {

	var log = logger.New()

	s.Host = os.Getenv("SERVER_HOST")
	s.Port = os.Getenv("SERVER_PORT")
	s.Timeout = os.Getenv("SERVER_READ_TIMEOUT")
	s.Address = fmt.Sprintf("%s:%s", s.Host, s.Port)

	s.Env = os.Getenv("ENVRONMENT")
	s.Name = os.Getenv("APP_NAME")
	s.TZ = os.Getenv("TZ")

	fmt.Println("#")
	fmt.Println("# $TZ: .............", s.TZ)
	fmt.Println("# $HOST: ...........", s.Host)
	fmt.Println("# $PORT: ...........", s.Port)
	fmt.Println("# $APP_NAME: .......", s.Name)
	fmt.Println("# $ENVIRONMENT: ....", s.Env)
	fmt.Println("# SERVER UP: .......", s.DateStart.Format("2006-01-02 15:04:05"), "UTC")
	fmt.Println("")

	log.Info().Msgf("[UP]: %v | [TZ]: %v | [HOST]: %v:%v | [ENVIRONMENT]: %v", s.DateStart.Format(time.RFC3339Nano), s.TZ, s.Host, s.Port, s.Env)

	return s.instance

}

func (s *Server) Listen() *fiber.App {

	var log = logger.New()

	s.Environments()
	err := s.instance.Listen(s.Address)
	if err != nil {
		log.Error().Err(err).Msgf("O Server foi derrubado - %v", s.DateStart.Format(time.RFC3339Nano))
	}
	return s.instance
}

func (s *Server) ListenGracefulShutdown() *fiber.App {

	var log = logger.New()

	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		err := s.instance.Shutdown()
		if err != nil {
			log.Error().Err(err).Msgf("O server não pode ser derrubado - %v", s.DateStart.Format(time.RFC3339Nano))
		}

		close(idleConnsClosed)
	}()

	s.Environments()
	err := s.instance.Listen(s.Address)
	if err != nil {
		log.Error().Err(err).Msgf("O Server foi derrubado - %v", s.DateStart.Format(time.RFC3339Nano))
	}

	<-idleConnsClosed
	return s.instance
}

func (s *Server) Start() *fiber.App {

	if s.Env == "dev" {
		return s.Listen()
	}

	return s.ListenGracefulShutdown()
}
