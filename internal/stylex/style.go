package stylex

import "github.com/charmbracelet/lipgloss"

var (
	Border    = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("5")).Padding(1, 4).Align(lipgloss.Center).Render
	Underline = lipgloss.NewStyle().Underline(true).Render
)
