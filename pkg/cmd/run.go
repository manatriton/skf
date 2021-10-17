package cmd

import (
	"fmt"
	"skf/pkg/api"
	"skf/pkg/conf"
	"skf/pkg/theme"

	"github.com/spf13/cobra"
)

func NewRunCommand(c *conf.Conf) *cobra.Command {
	runCmd := &cobra.Command{
		Use: "run",
	}

	runCmd.AddCommand(
		newRunPlanOutputCommand(c),
	)
	return runCmd
}

func newRunPlanOutputCommand(c *conf.Conf) *cobra.Command {
	var id string
	runPlanOutputCommand := &cobra.Command{
		Use: "plan-output",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPlanOutput(c, id)
		},
	}

	runPlanOutputCommand.Flags().StringVar(&id, "id", "", "id of the run")
	return runPlanOutputCommand
}

func runPlanOutput(c *conf.Conf, id string) error {
	value, err := c.API.Runs.GetPlanOutput(id)
	if err == api.ErrNotExist {
		fmt.Printf("Run %s does not exist!\n", id)
		return nil
	} else if err != nil {
		return err
	}

	if value == "" {
		fmt.Printf("Run %s has not yet been planned.\n", id)
	}

	theme.Title.Printf("Plan output for run %s\n\n", id)
	fmt.Print(value)
	fmt.Println()
	return nil
}
