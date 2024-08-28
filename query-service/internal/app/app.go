package app

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/query/pb"
	"github.com/amiosamu/gophkeeper/query-service/internal/config"
	"github.com/amiosamu/gophkeeper/query-service/internal/db"
	"github.com/amiosamu/gophkeeper/query-service/internal/services"
)

type App struct {
	cfg        *config.Config
	server     *services.QueryServer
	serverGRPC *grpc.Server
	storage    db.Storage
}

func NewApp() *App {
	return &App{
		cfg: config.NewConfig(),
	}
}

func (app *App) Run() error {
	app.storage = db.NewPostgres(app.cfg.DBPostgresURL)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigint
		if app.serverGRPC != nil {
			app.serverGRPC.GracefulStop()
		}
	}()

	app.server = services.NewQueryServer(app.storage)

	err := app.startGRPC()
	close(sigint)
	return err
}

func (app *App) startGRPC() error {
	listen, err := net.Listen("tcp", app.cfg.Port)

	if err != nil {
		return err
	}

	app.serverGRPC = grpc.NewServer()

	pb.RegisterQueryServiceServer(
		app.serverGRPC,
		app.server,
	)

	log.Printf("Query service start on %v", app.cfg.Port)

	err = app.serverGRPC.Serve(listen)
	if err == nil {
		log.Println("Query service graceful shutdown")
	}
	return err
}
