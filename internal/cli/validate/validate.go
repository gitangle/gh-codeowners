package validate

import (
	"context"
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/repository"

	"github.com/gitcheasy/gh-codeowners/internal/client"
	"github.com/gitcheasy/gh-codeowners/internal/issues"
	"github.com/gitcheasy/gh-codeowners/internal/stylex"
)

var singleRepoNotSupportedMessage = fmt.Sprintf("Single repository lookup is not yet supported. %s", stylex.Underline("Please use '--all'."))

// Run executes the validate command.
func Run(ctx context.Context, opts Options) error {
	if !opts.All {
		msg := stylex.Border("Howdy! You discovered not yet implemented feature.\n\n", singleRepoNotSupportedMessage)
		fmt.Println(msg)
		os.Exit(1)
	}

	if opts.Owner == "" {
		current, err := repository.Current()
		if err != nil {
			return err
		}
		opts.Owner = current.Owner
	}

	cli, err := client.NewGraphQLClient()
	if err != nil {
		return err
	}

	finder := issues.NewFinder(cli, opts.Owner, opts.IgnoredRepos)
	foundIssues, err := finder.ListCodeownersIssues(ctx, issues.RepoOptions{
		IsFork:     false,
		IsArchived: false,
		Visibility: issues.RepoVisibilityPublic,
	})
	if err != nil {
		return err
	}

	if len(foundIssues.InvalidOwners) == 0 && len(foundIssues.MissingOwnersFiles) == 0 {
		fmt.Println(stylex.Border("No issues found. ", stylex.Underline("Have a nice day!")))
		return nil
	}

	printer, err := issues.NewPrinter()
	if err != nil {
		return err
	}
	err = printer.PrintMissingOwnersFile(foundIssues.MissingOwnersFiles, opts.Owner)
	if err != nil {
		return err
	}

	err = printer.PrintInvalidOwners(foundIssues.InvalidOwners, opts.Owner)
	if err != nil {
		return err
	}

	os.Exit(opts.IssuesExitCode)
	return nil
}
