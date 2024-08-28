package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/services/pb"
	"github.com/amiosamu/gophkeeper/client/internal/client"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get gophkeeper stored user data",
		Long:  `This command can be used to get user stored data at gophkeeper`,
	}
)

func init() {
	getCmd.PersistentFlags().StringVarP(&nameValue, "user", "u", "", "user name")
	getCmd.MarkPersistentFlagRequired("user")

	getCmd.PersistentFlags().StringVarP(&passwordValue, "password", "p", "", "user password")
	getCmd.MarkPersistentFlagRequired("password")
}

type GetCmd struct {
	Command *cobra.Command
}

func NewGetCmd(ctx context.Context, cli *client.ServiceClient) *GetCmd {
	anyType := &cobra.Command{
		Use:   "any",
		Short: "any type of stored data",
		Long:  `This command can be used to get any type of stored data`,
		Run: func(cmd *cobra.Command, args []string) {
			res, err := cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: nameValue,
				Password: passwordValue,
			})

			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}

			if res.Token == "" {
				log.Fatal("empty token")
			}

			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_ANY,
			})

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(qres)
		},
	}

	binType := &cobra.Command{
		Use:   "bin",
		Short: "bin type of stored data",
		Long:  `This command can be used to get binnary type of stored data`,
		Run: func(cmd *cobra.Command, args []string) {
			res, err := cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: nameValue,
				Password: passwordValue,
			})

			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}

			if res.Token == "" {
				log.Fatal("empty token")
			}

			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_BINNARY,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(qres)
		},
	}

	cardType := &cobra.Command{
		Use:   "card",
		Short: "card type of stored data",
		Long:  `This command can be used to get card type of stored data`,

		Run: func(cmd *cobra.Command, args []string) {
			res, err := cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: nameValue,
				Password: passwordValue,
			})

			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}

			if res.Token == "" {
				log.Fatal("empty token")
			}

			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_CARD,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(qres)
		},
	}

	credType := &cobra.Command{
		Use:   "cred",
		Short: "login pasword type of stored data",
		Long:  `This command can be used to get login password type of stored data`,
		Run: func(cmd *cobra.Command, args []string) {
			res, err := cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: nameValue,
				Password: passwordValue,
			})

			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}

			if res.Token == "" {
				log.Fatal("empty token")
			}

			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_CARD,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(qres)
		},
	}

	textType := &cobra.Command{
		Use:   "text",
		Short: "text type of stored data",
		Long:  `This command can be used to get text type of stored data`,

		Run: func(cmd *cobra.Command, args []string) {
			res, err := cli.Client.Login(ctx, &pb.LoginRequest{
				UserName: nameValue,
				Password: passwordValue,
			})

			if err != nil {
				log.Fatal(err)
			}

			if res == nil {
				log.Fatal("something went wrong")
			}

			if res.Token == "" {
				log.Fatal("empty token")
			}

			qres, err := cli.Client.Query(context.WithValue(ctx, client.KeyPrincipalID, res.Token), &pb.QueryRequest{
				Type: pb.MessageType_CARD,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(qres)
		},
	}

	getCmd.AddCommand(anyType)
	getCmd.AddCommand(binType)
	getCmd.AddCommand(cardType)
	getCmd.AddCommand(credType)
	getCmd.AddCommand(textType)

	return &GetCmd{
		Command: getCmd,
	}
}
