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

// Printer prints issues in Markdown format.
type Printer struct {
	termWidth int
	terminal  term.Term
	renderer  *glamour.TermRenderer
}

// NewPrinter creates a new Printer.
func NewPrinter() (*Printer, error) {
	terminal := term.FromEnv()
	termWidth, _, _ := terminal.Size()

	renderer, err := glamour.NewTermRenderer(glamour.WithWordWrap(termWidth), glamour.WithEnvironmentConfig())
	if err != nil {
		return nil, fmt.Errorf("while creating glamour renderer: %w", err)
	}

	return &Printer{
		terminal:  terminal,
		termWidth: termWidth,
		renderer:  renderer,
	}, nil
}

// PrintInvalidOwners prints missing owners file in the Markdown format.
func (p *Printer) PrintInvalidOwners(issues map[string][]string, organization string) error {
	render := func(s ...string) string {
		return strings.Join(s, " ")
	}
	if p.terminal.IsTerminalOutput() {
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
		fmt.Fprintln(&out, "")
	}

	if out.String() == "" {
		return nil
	}

	rawMd := fmt.Sprintf("# Invalid CODEOWNERS \n%s", out.String())

	md, err := p.renderer.Render(rawMd)
	if err != nil {
		return fmt.Errorf("while rendering markdown: %w", err)
	}

	md = strings.ReplaceAll(md, "^", render("^"))
	md = issueIntro.ReplaceAllString(md, fmt.Sprintf("%s $2", render("Unknown owner on $1:")))

	fmt.Fprintln(p.terminal.Out(), md)

	return nil
}

// PrintMissingOwnersFile prints missing owners file in the Markdown format.
func (p *Printer) PrintMissingOwnersFile(missingFiles []string, org string) error {
	if len(missingFiles) == 0 {
		return nil
	}

	var out bytes.Buffer
	for _, name := range missingFiles {
		fmt.Fprintf(&out, "- [%s](https://github.com/%s/%s)\n", name, org, name)
	}

	rawMd := fmt.Sprintf("# Missing CODEOWNERS files \n%s", out.String())

	md, err := p.renderer.Render(rawMd)
	if err != nil {
		return fmt.Errorf("while rendering markdown: %w", err)
	}

	fmt.Fprintln(p.terminal.Out(), md)

	return nil
}
