package commands

import (
	"github.com/spf13/cobra"

	"github.com/argoproj/argo/cmd/argo/commands/cron"

	"github.com/argoproj/argo/cmd/argo/commands/archive"
	"github.com/argoproj/argo/cmd/argo/commands/client"
	"github.com/argoproj/argo/cmd/argo/commands/template"
	"github.com/argoproj/argo/util/cmd"
)

const (
	// CLIName is the name of the CLI
	CLIName = "argo"
)

// NewCommand returns a new instance of an argo command
func NewCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   CLIName,
		Short: "argo is the command line interface to Argo",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewCompletionCommand())
	command.AddCommand(NewDeleteCommand())
	command.AddCommand(NewGetCommand())
	command.AddCommand(NewLintCommand())
	command.AddCommand(NewListCommand())
	command.AddCommand(NewLogsCommand())
	command.AddCommand(NewResubmitCommand())
	command.AddCommand(NewResumeCommand())
	command.AddCommand(NewRetryCommand())
	command.AddCommand(NewServerCommand())
	command.AddCommand(NewSubmitCommand())
	command.AddCommand(NewSuspendCommand())
	command.AddCommand(NewTokenCommand())
	command.AddCommand(NewWatchCommand())
	command.AddCommand(NewWatchCommand())
	command.AddCommand(NewTerminateCommand())
	command.AddCommand(archive.NewArchiveCommand())
	command.AddCommand(cmd.NewVersionCmd(CLIName))
	command.AddCommand(template.NewTemplateCommand())
	command.AddCommand(cron.NewCronWorkflowCommand())
	client.AddKubectlFlagsToCmd(command)
	client.AddArgoServerFlagsToCmd(command)
	return command
}
