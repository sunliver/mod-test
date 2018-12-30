package modinner

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "web_server",
	Short: "web_server is lambdacal's rdp project's webserver",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infoln("can not go to definition")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
