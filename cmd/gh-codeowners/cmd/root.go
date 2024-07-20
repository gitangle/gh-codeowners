package cmd

import (
	"github.com/spf13/cobra"
	"go.szostok.io/version/extension"

	"github.com/gitangle/gh-codeowners/internal/cli"
	"github.com/gitangle/gh-codeowners/internal/heredoc"
)

const (
	orgName  = "gitangle"
	repoName = "gh-codeowners"
)

// NewRoot returns a root cobra.Command for the whole Codeowners CLI.
func NewRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   cli.Name,
		Short: "Codeowners CLI",
		Long: heredoc.WithCLIName(`
        <cli> - Codeowners CLI

        A utility that simplifies managing and validating CODEOWNERS files.

        Quick Start:

            $ <cli> validate                              # Validate CODEOWNERS files

            `, cli.Name),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.AddCommand(
		NewValidate(),
		NewDocs(),
		extension.NewVersionCobraCmd(
			extension.WithUpgradeNotice(orgName, repoName),
		),
	)

	return rootCmd
}
