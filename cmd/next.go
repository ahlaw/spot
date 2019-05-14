package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nextCmd)
}

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Play next track",
	RunE:  next,
}

func next(cmd *cobra.Command, args []string) error {
	return client.Next()
}
