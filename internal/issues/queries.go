package issues

import graphql "github.com/cli/shurcooL-graphql"

var _ RepoGetter = (*UserQuery)(nil)
var _ RepoGetter = (*OrgQuery)(nil)

// RepoGetter is the interface that allows to get the repositories.
type RepoGetter interface {
	GetRepositories() GQLRepositories
}

// UserQuery is the GraphQL query for getting the user's repositories.
type UserQuery struct {
	User WithRepositories `graphql:"user(login: $owner)"`
}

// GetRepositories implements RepoGetter.
func (q UserQuery) GetRepositories() GQLRepositories { return q.User.Repositories }

// OrgQuery is the GraphQL query for getting the organization's repositories.
type OrgQuery struct {
	Organization WithRepositories `graphql:"organization(login: $owner)"`
}

// GetRepositories implements RepoGetter.
func (q OrgQuery) GetRepositories() GQLRepositories { return q.Organization.Repositories }

// RepositoryVisibility is the visibility of the repository.
type RepositoryVisibility string

const (
	// RepoVisibilityPublic repositories are visible to everyone.
	RepoVisibilityPublic RepositoryVisibility = "PUBLIC"
	// RepoVisibilityPrivate repositories are only visible to organization members.
	RepoVisibilityPrivate RepositoryVisibility = "PRIVATE"
)

// RepoOptions are the options for the repository query.
type RepoOptions struct {
	IsFork     bool
	IsArchived bool
	Visibility RepositoryVisibility
}

type (
	// WithRepositories is the GraphQL query for getting the repositories.
	WithRepositories struct {
		Repositories GQLRepositories `graphql:"repositories(first: 100, visibility: $visibility, isFork: $isFork, isArchived: $isArchived)"`
	}

	// GQLRepositories represents a list of GitHub repositories with CODEOWNERS errors.
	GQLRepositories struct {
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
	}
)

// ReposQueryVars returns the variables for the repository query.
func ReposQueryVars(org string, filters RepoOptions) map[string]any {
	if filters.Visibility == "" {
		filters.Visibility = RepoVisibilityPublic
	}
	return map[string]any{
		"owner":      graphql.String(org),
		"isFork":     graphql.Boolean(filters.IsFork),
		"isArchived": graphql.Boolean(filters.IsArchived),
		"visibility": filters.Visibility,
	}
}
