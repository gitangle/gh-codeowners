package cmd

import (
	"github.com/spf13/cobra"

	"github.com/gitcheasy/gh-codeowners/internal/cli"
	"github.com/gitcheasy/gh-codeowners/internal/cli/validate"
	"github.com/gitcheasy/gh-codeowners/internal/heredoc"
)

// NewValidate returns a cobra.Command for logging into a Botkube Cloud.
func NewValidate() *cobra.Command {
	var opts validate.Options

	cmd := &cobra.Command{
		Use:   "validate [OPTIONS]",
		Short: "Validate CODEOWNERS files",
		Example: heredoc.WithCLIName(`
			# Validate CODEOWNERS across all repos for the GitHub owner taken from the current repository directory.
			<cli> validate --all
			
			# Validate CODEOWNERS for a specific owner across their repos 
			<cli> validate --owner "mszostok" --all
		`, cli.Name),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return validate.Run(cmd.Context(), opts)
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&opts.Owner, "owner", "", "GitHub organization or user. Defaults to the GitHub owner of the repository in the current directory.")
	flags.BoolVar(&opts.All, "all", false, "Whether to list issues for all repositories within a GitHub owner.")
	flags.StringSliceVar(&opts.IgnoredRepos, "ignore-repos", nil, "Comma-separated list of repository names to ignore.")
	flags.IntVar(&opts.IssuesExitCode, "issues-exit-code", 3, "Exit code when issues were found.")

	return cmd
}
