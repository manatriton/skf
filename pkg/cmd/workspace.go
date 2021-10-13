package cmd

import (
	"fmt"
	"skf/pkg/conf"
	"skf/pkg/theme"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type workspaceCreateOpts struct {
	name string
}

type workspaceViewOpts struct {
	name string
}

func NewWorkspaceCommand(c *conf.Conf) *cobra.Command {
	workspaceCmd := &cobra.Command{
		Use: "workspace",
	}

	workspaceCmd.AddCommand(newWorkspaceListCommand(c))
	workspaceCmd.AddCommand(newWorkspaceCreateCommand(c))
	workspaceCmd.AddCommand(newWorkspaceViewCommand(c))

	return workspaceCmd
}

func newWorkspaceListCommand(c *conf.Conf) *cobra.Command {
	workspaceListCommand := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("listing workspaces")
			return nil
		},
	}

	return workspaceListCommand
}

func workspaceList(c *conf.Conf, opts *workspaceViewOpts) error {
	return nil
}

func newWorkspaceViewCommand(c *conf.Conf) *cobra.Command {
	opts := &workspaceViewOpts{}
	workspaceViewCommand := &cobra.Command{
		Use: "view",
		RunE: func(cmd *cobra.Command, args []string) error {
			return workspaceView(c, opts)
		},
	}

	workspaceViewCommand.Flags().StringVarP(&opts.name, "name", "n", "", "name of the workspace")
	return workspaceViewCommand
}

func workspaceView(c *conf.Conf, opts *workspaceViewOpts) error {
	if err := workspaceViewPromptOptions(opts); err != nil {
		return err
	}

	workspace, err := c.API.Workspaces.WorkspaceByName(opts.name)
	if err != nil {
		return err
	}

	if workspace == nil {
		fmt.Printf("No workspace found :(\n\n")
	} else {
		color.New(color.FgWhite, color.Bold).Printf("%s\n", workspace.Name)
		fmt.Printf("<no description found>\n\n")
	}

	return nil
}

func workspaceViewPromptOptions(opts *workspaceViewOpts) error {
	if opts.name == "" {
		fmt.Printf("More information is needed to search for a workspace!\n\n")
		prompt := &survey.Input{
			Message: "Name",
		}
		if err := survey.AskOne(prompt, &opts.name); err != nil {
			return err
		}
		fmt.Println()
	}

	return nil
}

func newWorkspaceCreateCommand(c *conf.Conf) *cobra.Command {
	opts := &workspaceCreateOpts{}

	workspaceCreateCommand := &cobra.Command{
		Use: "create",
		RunE: func(cmd *cobra.Command, args []string) error {
			return workspaceCreate(c, opts)
		},
	}

	workspaceCreateCommand.Flags().StringVarP(&opts.name, "name", "n", "", "name of the workspace")
	return workspaceCreateCommand
}

func workspaceCreatePromptOptions(opts *workspaceCreateOpts) error {
	if opts.name == "" {
		prompt := &survey.Input{
			Message: "Name",
		}
		if err := survey.AskOne(prompt, &opts.name); err != nil {
			return err
		}
	}

	return nil
}

func workspaceCreate(c *conf.Conf, opts *workspaceCreateOpts) error {
	fmt.Println("\nCreating a new workspace\n")

	if err := workspaceCreatePromptOptions(opts); err != nil {
		return err
	}

	_, err := c.API.Workspaces.CreateWorkspace(opts.name)
	if err != nil {
		return err
	}

	theme.Title.Println("Workspace created successfully!")
	fmt.Println()
	return nil
}
