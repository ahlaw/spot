package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringP("type", "t", "track", "Search type")
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play Spotify",
	RunE:  play,
}

func play(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		typeFlag, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}

		var searchType spotify.SearchType
		switch typeFlag {
		case "album":
			searchType = spotify.SearchTypeAlbum
		case "artist":
			searchType = spotify.SearchTypeArtist
		case "playlist":
			searchType = spotify.SearchTypePlaylist
		default:
			searchType = spotify.SearchTypeTrack
		}

		opt, err := search(strings.Join(args, " "), searchType)
		if err != nil {
			return err
		}
		return client.PlayOpt(opt)
	}
	return client.Play()
}

func search(query string, searchType spotify.SearchType) (*spotify.PlayOptions, error) {
	results, err := client.Search(query, searchType)
	if err != nil {
		return nil, err
	}

	if results.Tracks != nil {
		opt := &spotify.PlayOptions{
			URIs: []spotify.URI{results.Tracks.Tracks[0].URI},
		}
		return opt, nil
	} else if results.Albums != nil {
		opt := &spotify.PlayOptions{
			PlaybackContext: &results.Albums.Albums[0].URI,
		}
		return opt, nil
	} else if results.Artists != nil {
		opt := &spotify.PlayOptions{
			PlaybackContext: &results.Artists.Artists[0].URI,
		}
		return opt, nil
	} else if results.Playlists != nil {
		opt := &spotify.PlayOptions{
			PlaybackContext: &results.Playlists.Playlists[0].URI,
		}
		return opt, nil
	}
	fmt.Println("Your search did not match any results")
	return nil, nil
}
