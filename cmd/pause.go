package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pauseCmd)
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause Spotify",
	RunE:  pause,
}

func pause(cmd *cobra.Command, args []string) error {
	return client.Pause()
}
