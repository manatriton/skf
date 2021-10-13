package main

import (
	"errors"
	"os"
	"skf/pkg/api"
	"skf/pkg/cmd"
	"skf/pkg/conf"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func defaultConf() *conf.Conf {
	return &conf.Conf{
		URL: "http://127.0.0.1:4000/graphql",
	}
}

func main() {
	c := defaultConf()

	root := &cobra.Command{
		Use:   "skf",
		Short: "Skyform's CLI tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Welcome to skf!")
			return nil
		},
	}

	root.PersistentFlags().StringVarP(&c.Token, "token", "t", "", "API token to authenticate with")
	root.AddCommand(cmd.NewWorkspaceCommand(c))
	root.AddCommand(cmd.NewVariableCommand(c))

	root.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		initializeConf(c)

		if c.Token == "" {
			return errors.New("no access token specified")
		}

		return nil
	}

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initializeConf(c *conf.Conf) {
	c.API = api.NewAPI(c.URL)

	token := os.Getenv("SKF_TOKEN")
	if token != "" {
		c.Token = token
	}
}
