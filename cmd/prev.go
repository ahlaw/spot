package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(prevCmd)
}

var prevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Play previous track",
	RunE:  prev,
}

func prev(cmd *cobra.Command, args []string) error {
	return client.Previous()
}
