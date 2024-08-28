package app

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/amiosamu/gophkeeper/auth-service/internal/config"
	"github.com/amiosamu/gophkeeper/auth-service/internal/db"
	"github.com/amiosamu/gophkeeper/auth-service/internal/pb"
	"github.com/amiosamu/gophkeeper/auth-service/internal/services"
	"github.com/amiosamu/gophkeeper/auth-service/internal/utils"
)

type App struct {
	cfg        *config.Config
	server     *services.AuthServer
	serverGRPC *grpc.Server
	jwt        *utils.JwtWraper
	postgres   db.Storage
}

func NewApp() *App {
	return &App{
		cfg: config.NewConfig(),
	}
}

func (app *App) Run() error {
	app.postgres = db.NewPostgres(app.cfg.DBPostgresURL)

	app.jwt = &utils.JwtWraper{
		SecretKey: app.cfg.JWTSecretKey,
		Issuer:    "Gophkeeper.Auth service",
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigint
		if app.serverGRPC != nil {
			app.serverGRPC.GracefulStop()
		}
	}()

	app.server = services.NewAuthServer(
		app.jwt,
		app.postgres,
	)

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

	pb.RegisterAuthServiceServer(
		app.serverGRPC,
		app.server,
	)

	log.Printf("Auth service start on %v", app.cfg.Port)

	err = app.serverGRPC.Serve(listen)
	if err == nil {
		log.Println("Auth service graceful shutdown")
	}
	return err
}
