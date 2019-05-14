package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show Spotify info",
	RunE:  info,
}

func info(cmd *cobra.Command, args []string) error {
	currPlayerStatus, err := client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}

	track := currPlayerStatus.Item
	artists := make([]string, len(track.SimpleTrack.Artists))
	for i, a := range track.SimpleTrack.Artists {
		artists[i] = a.Name
	}

	fmt.Printf("Artist: %s\n", strings.Join(artists, ", "))
	fmt.Printf("Album: %s\n", track.Album.Name)
	fmt.Printf("Track: %s\n", track.Name)

	return nil
}
