package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/amiosamu/gophkeeper/api-gateway/pkg/services/pb"
	"github.com/amiosamu/gophkeeper/client/internal/client"
	"github.com/amiosamu/gophkeeper/client/internal/models"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add gophkeeper stored user data",
		Long:  `This command can be used to add user stored data to gophkeeper`,
	}
)

func init() {
	addCmd.PersistentFlags().StringVarP(&nameValue, "user", "u", "", "user name")
	addCmd.MarkPersistentFlagRequired("user")

	addCmd.PersistentFlags().StringVarP(&passwordValue, "password", "p", "", "user password")
	addCmd.MarkPersistentFlagRequired("password")

	addCmd.PersistentFlags().StringVarP(&keyValue, "key", "k", "", "secret key")
	addCmd.MarkPersistentFlagRequired("key")
}

type AddCmd struct {
	Command *cobra.Command
}

func NewAddCmd(ctx context.Context, cli *client.ServiceClient) *AddCmd {
	var id int64
	var meta string

	// card data
	var cardowner string
	var cardnumber string
	var carddate string
	var cardcvv string
	// card

	cardTypeCmd := &cobra.Command{
		Use:   "card",
		Short: "card type of stored data",
		Long:  `This command can be used to add card type of stored data`,
		PreRun: func(cmd *cobra.Command, args []string) {
		},

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

			card := &models.Card{
				Owner:  cardowner,
				Number: cardnumber,
				Date:   carddate,
				CVV:    cardcvv,
			}
			rec := models.NewRecord(
				id,
				models.Add,
				models.Tcard,
				[]byte(keyValue),
				card,
				meta,
			)

			ares, err := cli.Client.AddCommand(context.WithValue(ctx, client.KeyPrincipalID, res.Token),
				&pb.CommandRequest{
					Type: pb.MessageType_CARD,
					Id:   rec.Id,
					Data: rec.Value,
					Meta: rec.Meta,
				})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(rec)
			fmt.Println(ares)

			fmt.Println("add card type executed")
		},
	}

	cardTypeCmd.Flags().StringVarP(&cardowner, "cardowner", "", "", "card owner")
	cardTypeCmd.MarkFlagRequired("cardowner")

	cardTypeCmd.Flags().StringVarP(&cardnumber, "cardnumber", "", "", "card number")
	cardTypeCmd.MarkFlagRequired("cardnumber")

	cardTypeCmd.Flags().StringVarP(&carddate, "carddate", "", "", "card expires date")
	cardTypeCmd.MarkFlagRequired("carddate")

	cardTypeCmd.Flags().StringVarP(&cardcvv, "cardcvv", "", "", "card cvv code")
	cardTypeCmd.MarkFlagRequired("cardcvv")

	cardTypeCmd.Flags().StringVarP(&meta, "meta", "", "", "meta information")

	// var login string
	// var pass string

	// userPasswordTypeCmd := &cobra.Command{
	// 	Use:   "cred",
	// 	Short: "credential type of stored data",
	// 	Long:  `This command can be used to add credential type to stored data`,

	// 	PreRun: func(cmd *cobra.Command, args []string) {
	// 		res, err = cli.Client.Login(ctx, &pb.LoginRequest{
	// 			UserName: name,
	// 			Password: password,
	// 		})
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		if res == nil {
	// 			log.Fatal("something went wrong")
	// 		}
	// 		if res.Token == "" {
	// 			log.Fatal("empty token")
	// 		}
	// 	},

	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		cred := &models.LoginPass{
	// 			Login: login,
	// 			Pass:  pass,
	// 		}
	// 		rec := models.NewRecord(
	// 			id,
	// 			models.Add,
	// 			models.Tcard,
	// 			[]byte(key),
	// 			cred,
	// 			meta,
	// 		)

	// 		ares, err := cli.Client.AddCommand(context.WithValue(ctx, client.KeyPrincipalID, res.Token),
	// 			&pb.CommandRequest{
	// 				Type: pb.MessageType_LOGIN_PASSWORD,
	// 				Id:   rec.Id,
	// 				Data: rec.Value,
	// 				Meta: rec.Meta,
	// 			})
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		fmt.Println(rec)
	// 		fmt.Println(ares)

	// 		fmt.Println("add login/pass type executed")
	// 	},
	// }

	// userPasswordTypeCmd.Flags().StringVarP(&key, "key", "", "", "secret key")
	// userPasswordTypeCmd.MarkFlagRequired("key")

	// userPasswordTypeCmd.Flags().StringVarP(&login, "login", "", "", "stored login")
	// userPasswordTypeCmd.MarkFlagRequired("login")

	// userPasswordTypeCmd.Flags().StringVarP(&pass, "pass", "", "", "stored password")
	// userPasswordTypeCmd.MarkFlagRequired("pass")

	//addCmd.AddCommand(cardTypeCmd)
	//	addCmd.AddCommand(userPasswordTypeCmd)

	addCmd.AddCommand(cardTypeCmd)

	return &AddCmd{
		Command: addCmd,
	}
}
