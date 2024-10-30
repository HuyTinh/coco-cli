package cmd

import (
	textInput "coco-cli/m/cmd/ui/text_input"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	tipMsgStyle    = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

func init() {
	rootCmd.AddCommand(createCmd)
}

type listOptions struct {
	options []string
}

type Options struct {
	ProjectName *textInput.Output
	ProjectType string
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "short description",
	Long:  ".",

	Run: func(cmd *cobra.Command, args []string) {
		options := Options{
			ProjectName: &textInput.Output{},
		}

		tprogram := tea.NewProgram(textInput.InitialTextInputModel(options.ProjectName, "What is the name of your project"))

		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}
	},
}
