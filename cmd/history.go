package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().StringP("number", "n", "10", "Number of tracks to show")
}

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show recently played tracks",
	RunE:  history,
}

func history(cmd *cobra.Command, args []string) error {
	numFlag, err := cmd.Flags().GetString("number")
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(numFlag)
	if err != nil {
		return err
	}

	if limit < 1 {
		limit = 1
	} else if limit > 50 {
		limit = 50
	}

	opt := &spotify.RecentlyPlayedOptions{
		Limit: limit,
	}
	recentItems, err := client.PlayerRecentlyPlayedOpt(opt)
	if err != nil {
		return err
	}

	for i, item := range recentItems {
		artists := make([]string, len(item.Track.Artists))
		for i, a := range item.Track.Artists {
			artists[i] = a.Name
		}
		fmt.Printf("%d. %s - %s\n", i+1, strings.Join(artists, ", "), item.Track.Name)
	}
	return nil
}
