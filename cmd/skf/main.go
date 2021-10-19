package main

import (
	"errors"
	"os"

	"github.com/manatriton/skf/pkg/api"
	"github.com/manatriton/skf/pkg/cmd"
	"github.com/manatriton/skf/pkg/conf"
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
	root.AddCommand(
		cmd.NewWorkspaceCommand(c),
		cmd.NewVariableCommand(c),
		cmd.NewRunCommand(c),
	)

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
	token := os.Getenv("SKF_TOKEN")
	if token != "" {
		c.Token = token
	}

	c.API = api.NewAPI(c.URL, c.Token)
}
