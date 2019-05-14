package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shareCmd)
}

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Copy Open Spotify URL of current track",
	RunE:  share,
}

func share(cmd *cobra.Command, args []string) error {
	currPlayerStatus, err := client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}
	url := currPlayerStatus.Item.SimpleTrack.ExternalURLs["spotify"]
	if err := clipboard.WriteAll(url); err != nil {
		return err
	}
	fmt.Printf("Copied %s to clipboard\n", url)
	return nil
}
