package issues

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// Finder finds GitHub issues for invalid CODEOWNERS files.
type Finder struct {
	ignoredRepos []string
}

// NewFinder creates a new Finder.
func NewFinder(ignoredRepos string) *Finder {
	return &Finder{ignoredRepos: strings.Split(ignoredRepos, ",")}
}

// GraphQLQuery is the interface that allows to execute GraphQL queries against the GitHub GraphQL server.
type GraphQLQuery interface {
	Query(name string, q any, variables map[string]any) error
}

// ListCodeownersIssuesOutput is the output of the ListCodeownersIssues function.
type ListCodeownersIssuesOutput struct {
	MissingOwnersFiles []string
	InvalidOwners      map[string][]string
}

// ListCodeownersIssues lists GitHub issues for invalid CODEOWNERS files.
func (f *Finder) ListCodeownersIssues(client GraphQLQuery, org string, filters RepoOptions) (*ListCodeownersIssuesOutput, error) {
	var errs []error
	for _, q := range []RepoGetter{&UserQuery{}, &OrgQuery{}} {
		err := client.Query("CodeownersIssues", q, ReposQueryVars(org, filters))
		if err != nil {
			errs = append(errs, fmt.Errorf("while executing GraphQL query: %w", err))
			continue
		}
		return f.collectIssues(q.GetRepositories()), nil
	}

	return nil, errors.Join(errs...)
}

func (f *Finder) collectIssues(repos GQLRepositories) *ListCodeownersIssuesOutput {
	var (
		missing []string
		issues  = map[string][]string{}
	)
	for _, repo := range repos.Nodes {
		if slices.Contains(f.ignoredRepos, repo.Name) {
			continue
		}

		if repo.Codeowners == nil {
			missing = append(missing, repo.Name)
			continue
		}

		if len(repo.Codeowners.Errors) == 0 {
			continue
		}
		issues[repo.Name] = []string{}
		for _, codeowners := range repo.Codeowners.Errors {
			issues[repo.Name] = append(issues[repo.Name], codeowners.Message)
		}
	}

	return &ListCodeownersIssuesOutput{
		MissingOwnersFiles: missing,
		InvalidOwners:      issues,
	}
}
