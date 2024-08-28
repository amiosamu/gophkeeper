package app

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/command/pb"
	"github.com/amiosamu/gophkeeper/command-consumer-service/internal/config"
	"github.com/amiosamu/gophkeeper/command-consumer-service/internal/db"
	"github.com/amiosamu/gophkeeper/command-consumer-service/internal/services"
)

type App struct {
	cfg        *config.Config
	server     *services.CommandServer
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

	app.server = services.NewCommandServer(app.storage)

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

	pb.RegisterCommandServiceServer(
		app.serverGRPC,
		app.server,
	)

	log.Printf("Command consumer service start on %v", app.cfg.Port)

	err = app.serverGRPC.Serve(listen)
	if err == nil {
		log.Println("Command consumer service graceful shutdown")
	}
	return err
}
