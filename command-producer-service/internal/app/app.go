package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/command/pb"
	"github.com/amiosamu/gophkeeper/command-producer-service/internal/client"
	"github.com/amiosamu/gophkeeper/command-producer-service/internal/config"
	"github.com/amiosamu/gophkeeper/command-producer-service/internal/services"
)

type App struct {
	cfg        *config.Config
	server     *services.CommandServer
	serverGRPC *grpc.Server
}

func NewApp() *App {
	return &App{
		cfg: config.NewConfig(),
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigint
		if a.serverGRPC != nil {
			a.serverGRPC.GracefulStop()
		}
	}()

	if a.cfg.KafkaURL == "" {
		cli := client.NewCommandServiceClient(ctx, a.cfg)
		a.server = services.NewCommandServer(cli.Client)
	}

	err := a.startGRPC()
	close(sigint)
	return err
}

func (a *App) startGRPC() error {
	listen, err := net.Listen("tcp", a.cfg.Port)

	if err != nil {
		return err
	}

	a.serverGRPC = grpc.NewServer()

	pb.RegisterCommandServiceServer(
		a.serverGRPC,
		a.server,
	)

	log.Printf("Query service start on %v", a.cfg.Port)

	err = a.serverGRPC.Serve(listen)
	if err == nil {
		log.Println("Query service graceful shutdown")
	}
	return err
}
