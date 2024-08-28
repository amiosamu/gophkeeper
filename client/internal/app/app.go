package app

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	"github.com/amiosamu/gophkeeper/client/internal/client"
	"github.com/amiosamu/gophkeeper/client/internal/commands"
	"github.com/amiosamu/gophkeeper/client/internal/config"
)

type App struct {
	cfg *config.Config
}

func NewApp() *App {
	return &App{
		cfg: config.NewConfig(),
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serviceClient := client.NewServiceClient()
	serviceClient.Dial(ctx, a.cfg)

	versionCmd := commands.NewVersionCmd()
	registerCmd := commands.NewRegisterCmd(ctx, serviceClient)
	getCmd := commands.NewGetCmd(ctx, serviceClient)
	//addCmd := commands.NewAddCmd(ctx, serviceClient)

	var rootCmd = &cobra.Command{Use: "gophkeeper"}

	rootCmd.AddCommand(versionCmd.Command)
	rootCmd.AddCommand(registerCmd.Command)
	rootCmd.AddCommand(getCmd.Command)
	//rootCmd.AddCommand(addCmd.Command)

	err := rootCmd.Execute()

	ctx.Done()
	return err
}
