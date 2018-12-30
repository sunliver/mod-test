package main

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"mod-test/modinner"
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

func main() {
	fmt.Println("mod-test")

	// slow
	logrus.Infoln("can go to definition, but fallback on old godef")
	cron.New()
	viper.Get("hello")

	// can not go to definition
	rootCmd.Execute()

	modinner.Execute()

}
