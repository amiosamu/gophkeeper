package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type VersionCmd struct {
	Command *cobra.Command
}

func NewVersionCmd() *VersionCmd {
	return &VersionCmd{
		Command: &cobra.Command{
			Use:   "version",
			Short: "Print version number of gophkeeper",
			Long:  `This command can be used get the version number of gophkeeper`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("gophkeeper v0.0.1")
			},
		},
	}
}
