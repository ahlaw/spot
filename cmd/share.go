package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.Flags().BoolP("track", "t", false, "Copy current track instead of context")
}

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Copy Open Spotify URL of current context",
	RunE:  share,
}

func share(cmd *cobra.Command, args []string) error {
	trackFlag, err := cmd.Flags().GetBool("track")
	if err != nil {
		return err
	}

	currPlayerStatus, err := client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}

	var url string
	contextType := currPlayerStatus.PlaybackContext.Type
	if trackFlag || contextType == "" {
		url = currPlayerStatus.Item.SimpleTrack.ExternalURLs["spotify"]
	} else {
		url = currPlayerStatus.PlaybackContext.ExternalURLs["spotify"]
	}
	if err := clipboard.WriteAll(url); err != nil {
		return err
	}
	fmt.Printf("Copied %s to clipboard\n", url)
	return nil
}
