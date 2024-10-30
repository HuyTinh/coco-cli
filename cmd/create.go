package cmd

import "github.com/spf13/cobra"

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "short description",
	Long:  ".",

	Run: func(cmd *cobra.Command, args []string) {

	},
}
