package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cli/go-gh/v2/pkg/repository"

	"github.com/gitcheasy/gh-codeowners/internal/client"
	"github.com/gitcheasy/gh-codeowners/internal/issues"
	"github.com/gitcheasy/gh-codeowners/internal/stylex"
)

var singleRepoNotSupportedMessage = fmt.Sprintf("Single repository lookup is not yet supported. %s", stylex.Underline("Please use '--all'."))

func main() {
	org := flag.String("owner", "", "GitHub organization or user. Defaults to the GitHub owner of the repository in the current directory.")
	all := flag.Bool("all", false, "Whether to list issues for all repositories within a GitHub owner.")
	ignoredRepos := flag.String("ignore-repos", "", "Comma-separated list of repository names to ignore.")
	issuesExitCode := flag.Int("issues-exit-code", 3, "Exit code when issues were found (default 3)")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		exitOnError(errors.New("No command specified"))
	}
	switch args[0] {
	case "validate":
	default:
		exitOnError(errors.New("Unknown command"))
	}

	if !*all {
		msg := stylex.Border(fmt.Sprintf("Howdy! You discovered not yet implemented feature.\n\n%s", singleRepoNotSupportedMessage))
		fmt.Println(msg)
		os.Exit(1)
	}
	if *org == "" {
		current, err := repository.Current()
		exitOnError(err)
		org = &current.Owner
	}

	cli, err := client.NewGraphQLClient()
	exitOnError(err)

	finder := issues.NewFinder(*ignoredRepos)
	foundIssues, err := finder.ListCodeownersIssues(cli, *org, issues.RepoOptions{
		IsFork:     false,
		IsArchived: false,
		Visibility: issues.RepoVisibilityPublic,
	})
	exitOnError(err)

	if len(foundIssues.InvalidOwners) == 0 && len(foundIssues.MissingOwnersFiles) == 0 {
		log.Println("No issues found. Have a nice day!")
		os.Exit(0)
	}

	printer, err := issues.NewPrinter()
	exitOnError(err)
	err = printer.PrintMissingOwnersFile(foundIssues.MissingOwnersFiles, *org)
	exitOnError(err)

	err = printer.PrintInvalidOwners(foundIssues.InvalidOwners, *org)
	exitOnError(err)

	os.Exit(*issuesExitCode)
}

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
