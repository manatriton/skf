package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"

	"github.com/manatriton/skf/pkg/api"
	"github.com/manatriton/skf/pkg/conf"
	"github.com/manatriton/skf/pkg/theme"
	"github.com/spf13/cobra"
)

func NewRunCommand(c *conf.Conf) *cobra.Command {
	runCmd := &cobra.Command{
		Use: "run",
	}

	runCmd.AddCommand(
		newRunPlanOutputCommand(c),
		newRunViewCommand(c),
	)
	return runCmd
}

func newRunViewCommand(c *conf.Conf) *cobra.Command {
	var id string
	runViewCmd := &cobra.Command{
		Use: "view",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runView(c, id)
		},
	}

	runViewCmd.Flags().StringVar(&id, "id", "", "id of the run")
	return runViewCmd
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
	if id == "" {
		fmt.Printf("More information is needed to search for a workspace!\n\n")
		prompt := &survey.Input{
			Message: "Run id",
		}
		if err := survey.AskOne(prompt, &id); err != nil {
			return err
		}
		fmt.Println()
	}

	value, err := c.API.Runs.GetPlanOutputById(id)
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

func runView(c *conf.Conf, id string) error {
	if id == "" {
		fmt.Printf("More information is needed to search for a workspace!\n\n")
		prompt := &survey.Input{
			Message: "Run id",
		}
		if err := survey.AskOne(prompt, &id); err != nil {
			return err
		}
		fmt.Println()
	}

	run, err := c.API.Runs.GetById(id)
	if err != nil {
		return err
	}

	if run == nil {
		fmt.Println("Run not found!")
		return nil
	}

	fmt.Printf("%s\n", run.ID)
	fmt.Printf("status: %s\n", run.Status)
	fmt.Printf("created at: %s\n", run.CreatedAt)
	return nil
}
