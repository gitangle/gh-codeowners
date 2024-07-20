package issues

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh/v2/pkg/term"
)

var issueIntro = regexp.MustCompile(`Unknown owner on (.*):(.*)`)

func PrintInvalidOwners(issues map[string][]string, organization string) error {
	terminal := term.FromEnv()
	termWidth, _, _ := terminal.Size()
	render := func(s ...string) string {
		return fmt.Sprint(strings.Join(s, " "))
	}

	if terminal.IsTerminalOutput() {
		render = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render
	}

	var out bytes.Buffer
	for repo, issue := range issues {
		fmt.Fprintf(&out, "## [%s](https://github.com/%s/%s)\n", repo, organization, repo)
		for _, i := range issue {
			fmt.Fprintln(&out, "```")
			fmt.Fprintln(&out, i)
			fmt.Fprintln(&out, "```")
		}
		fmt.Fprintln(&out, "----")
		fmt.Fprintln(&out, "")
	}

	if out.String() == "" {
		return nil
	}

	r, err := glamour.NewTermRenderer(glamour.WithWordWrap(termWidth), glamour.WithEnvironmentConfig())
	if err != nil {
		return fmt.Errorf("while creating glamour renderer: %w", err)
	}

	rawMd := fmt.Sprintf("# Invalid CODEOWNERS \n%s", out.String())

	md, err := r.Render(rawMd)
	if err != nil {
		return fmt.Errorf("while rendering markdown: %w", err)
	}

	md = strings.ReplaceAll(md, "^", render("^"))
	md = issueIntro.ReplaceAllString(md, fmt.Sprintf("%s $2", render("Unknown owner on $1:")))

	fmt.Fprintln(terminal.Out(), md)

	return nil
}

func PrintMissingOwnersFile(missingFiles []string, org string) error {
	if len(missingFiles) == 0 {
		return nil
	}

	terminal := term.FromEnv()
	termWidth, _, _ := terminal.Size()

	var out bytes.Buffer
	for _, name := range missingFiles {
		fmt.Fprintf(&out, "- [%s](https://github.com/%s/%s)\n", name, org, name)

	}

	r, err := glamour.NewTermRenderer(glamour.WithWordWrap(termWidth), glamour.WithEnvironmentConfig())
	if err != nil {
		return fmt.Errorf("while creating glamour renderer: %w", err)
	}

	rawMd := fmt.Sprintf("# Missing CODEOWNERS files \n%s", out.String())

	md, err := r.Render(rawMd)
	if err != nil {
		return fmt.Errorf("while rendering markdown: %w", err)
	}

	fmt.Fprintln(terminal.Out(), md)

	return nil
}
