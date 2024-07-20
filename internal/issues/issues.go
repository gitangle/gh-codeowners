package issues

import (
	"fmt"
	"slices"

	graphql "github.com/cli/shurcooL-graphql"
)

// GraphQLQuery is the interface that allows to execute GraphQL queries against the GitHub GraphQL server.
type GraphQLQuery interface {
	Query(name string, q any, variables map[string]any) error
}

func ListCodeownersIssues(client GraphQLQuery, organization string, ignoredRepos []string) (map[string][]string, []string, error) {
	var query struct {
		Organization struct {
			Repositories struct {
				Nodes []struct {
					Name       string
					URL        string
					Codeowners *struct {
						Errors []struct {
							Path       string
							Source     string
							Kind       string
							Line       int
							Message    string
							Suggestion string
						}
					} `graphql:"codeowners"`
				}
			} `graphql:"repositories(first: 100, visibility: PUBLIC)"`
		} `graphql:"organization(login: $owner)"`
	}

	variables := map[string]interface{}{
		"owner": graphql.String(organization),
	}
	err := client.Query("OrgCodeownersIssues", &query, variables)
	if err != nil {
		return nil, nil, fmt.Errorf("while executing GraphQL query: %w", err)
	}

	var (
		missing []string
		issues  = map[string][]string{}
	)
	for _, repo := range query.Organization.Repositories.Nodes {
		if slices.Contains(ignoredRepos, repo.Name) {
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

	return issues, missing, nil
}
