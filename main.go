package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/cli/go-gh/v2/pkg/repository"

	"github.com/gitcheasy/gh-codeowners/internal/client"
	"github.com/gitcheasy/gh-codeowners/internal/issues"
)

func main() {
	org := flag.String("organization", "", "GitHub organization. Defaults to the GitHub organization of the repository in the current directory.")
	ignoredRepos := flag.String("ignore-repos", "", "Comma-separated list of repository names to ignore.")
	issuesExitCode := flag.Int("issues-exit-code", 3, "Exit code when issues were found (default 3)")

	flag.Parse()

	if *org == "" {
		current, err := repository.Current()
		exitOnError(err)
		org = &current.Owner
	}

	var ignoredReposList []string
	if *ignoredRepos != "" {
		ignoredReposList = strings.Split(*ignoredRepos, ",")
	}
	cli, err := client.NewGraphQLClient()
	exitOnError(err)

	invalidOwners, missingFiles, err := issues.ListCodeownersIssues(cli, *org, ignoredReposList)
	exitOnError(err)

	if len(missingFiles) == 0 && len(invalidOwners) == 0 {
		log.Println("No issues found.")
		os.Exit(0)
	}
	err = issues.PrintMissingOwnersFile(missingFiles, *org)
	exitOnError(err)

	err = issues.PrintInvalidOwners(invalidOwners, *org)
	exitOnError(err)
	os.Exit(*issuesExitCode)
}

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
