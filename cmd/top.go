package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

func init() {
	rootCmd.AddCommand(topCmd)
	topCmd.Flags().StringP("number", "n", "10", "Number of artists to show")
	topCmd.Flags().StringP("timerange", "t", "short", "[length] timerange considered: short, medium or long")
}

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Show top played artists by user",
	RunE:  top,
}

func top(cmd *cobra.Command, args []string) error {
	numFlag, err := cmd.Flags().GetString("number")
	timerange, err := cmd.Flags().GetString("timerange")
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

	opt := &spotify.Options{
		Limit:     &limit,
		Timerange: &timerange,
	}
	topArtists, err := client.CurrentUsersTopArtistsOpt(opt)
	if err != nil {
		return err
	}

	for i, artist := range topArtists.Artists {
		fmt.Printf("%d. %s\n", i+1, artist.SimpleArtist.Name)
	}
	return nil
}
