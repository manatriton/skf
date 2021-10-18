package cmd

import (
	"fmt"

	"github.com/manatriton/skf/pkg/api"

	"github.com/fatih/color"
	"github.com/manatriton/skf/pkg/conf"
	"github.com/spf13/cobra"
)

type variableCreateOpts struct {
	workspaceId string
	key         string
	value       string
	sensitive   bool
}

type variableUpdateOpts struct {
	id        string
	key       string
	value     string
	sensitive bool
}

func NewVariableCommand(c *conf.Conf) *cobra.Command {
	variableCmd := &cobra.Command{
		Use: "var",
	}

	variableCmd.AddCommand(
		newVariableGetCommand(c),
		newVariableCreateCommand(c),
		newVariableDeleteCommand(c),
		newVariableUpdateCommand(c),
	)
	return variableCmd
}

func newVariableGetCommand(c *conf.Conf) *cobra.Command {
	var id string
	variableGetCmd := &cobra.Command{
		Use: "view",
		RunE: func(cmd *cobra.Command, args []string) error {
			return variableGet(c, id)
		},
	}

	variableGetCmd.Flags().StringVar(&id, "id", "", "id of the variable")
	return variableGetCmd
}

func newVariableDeleteCommand(c *conf.Conf) *cobra.Command {
	var id string
	variableDeleteCmd := &cobra.Command{
		Use: "delete",
		RunE: func(cmd *cobra.Command, args []string) error {
			return variableDelete(c, id)
		},
	}

	variableDeleteCmd.Flags().StringVar(&id, "id", "", "id of the variable")
	return variableDeleteCmd
}

func newVariableCreateCommand(c *conf.Conf) *cobra.Command {
	opts := &variableCreateOpts{}
	variableCreateCmd := &cobra.Command{
		Use: "create",
		RunE: func(cmd *cobra.Command, args []string) error {
			return variableCreate(c, opts)
		},
	}

	flags := variableCreateCmd.Flags()
	flags.StringVar(&opts.workspaceId, "workspace-id", "", "id of the workspace")
	flags.StringVar(&opts.key, "key", "", "the key for the variable")
	flags.StringVar(&opts.value, "value", "", "the value for the variable")
	flags.BoolVar(&opts.sensitive, "sensitive", false, "indicates if the value is sensitive")
	return variableCreateCmd
}

func newVariableUpdateCommand(c *conf.Conf) *cobra.Command {
	opts := &variableUpdateOpts{}
	variableUpdateCmd := &cobra.Command{
		Use: "update",
		RunE: func(cmd *cobra.Command, args []string) error {
			return variableUpdate(c, opts)
		},
	}

	flags := variableUpdateCmd.Flags()
	flags.StringVar(&opts.id, "id", "", "id of the variable")
	flags.StringVar(&opts.key, "key", "", "the new key for the variable")
	flags.StringVar(&opts.value, "value", "", "the new value for the variable")
	flags.BoolVar(&opts.sensitive, "sensitive", false, "indicates if the value is sensitive")

	return variableUpdateCmd
}

func variableGet(c *conf.Conf, id string) error {
	v, err := c.API.Variables.GetVariable(id)
	if err != nil {
		return err
	}

	white := color.New(color.FgWhite, color.Bold).SprintFunc()

	fmt.Printf("%s (belongs to workspace %s)\n", white(v.ID), white(v.WorkspaceID))

	value := v.Value
	if v.Sensitive {
		value = "<value hidden>"
	}

	fmt.Printf("key: %s\nvalue: %s\n\n", v.Key, value)
	return nil
}

func variableCreate(c *conf.Conf, opts *variableCreateOpts) error {
	// TODO: Prompt for missing inputs
	_, err := c.API.Variables.Create(opts.workspaceId, opts.key, opts.value, opts.sensitive)
	if err != nil {
		return err
	}

	fmt.Println("Variable created successfully!")
	return nil
}

func variableDelete(c *conf.Conf, id string) error {
	_, err := c.API.Variables.DeleteById(id)
	if err != nil {
		return err
	}

	fmt.Printf("Variable %s was deleted (if it existed)\n", id)
	return nil
}

func variableUpdate(c *conf.Conf, opts *variableUpdateOpts) error {
	v, err := c.API.Variables.UpdateById(opts.id, opts.key, opts.value, opts.sensitive)
	if err != nil {
		return err
	}

	fmt.Printf("Variable %s updated successfully\n", v.ID)
	fmt.Printf("key: %s\nvalue: %s\n\n", v.Key, value(v))
	return nil
}

func value(v *api.Variable) string {
	if v.Sensitive {
		return "<value hidden>"
	}

	return string(v.Value)
}
