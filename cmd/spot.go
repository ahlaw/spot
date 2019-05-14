package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const (
	redirectURI = "http://localhost:8080/callback"
)

var (
	auth      spotify.Authenticator
	token     *oauth2.Token
	client    spotify.Client
	tokenPath string
)

var rootCmd = &cobra.Command{
	Use:               "spot",
	Short:             "Spot is a command-line interface for Spotify",
	PersistentPreRun:  preRootCmd,
	PersistentPostRun: postRootCmd,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func preRootCmd(cmd *cobra.Command, args []string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	tokenPath = filepath.Join(usr.HomeDir, ".spot")
	auth = spotify.NewAuthenticator(
		redirectURI,
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserReadRecentlyPlayed,
		spotify.ScopeUserTopRead,
	)

	_, ok := os.LookupEnv("SPOT_ID")
	if !ok {
		log.Fatal("App Client ID missing")
	}
	_, ok = os.LookupEnv("SPOT_SECRET")
	if !ok {
		log.Fatal("App Client Secret missing")
	}
	spotifyClientID := os.Getenv("SPOT_ID")
	spotifyClientSecret := os.Getenv("SPOT_SECRET")
	auth.SetAuthInfo(spotifyClientID, spotifyClientSecret)

	if cmd.Use == "login" {
		return
	}
	token, err = readToken()
	if err != nil {
		if os.IsNotExist(err) {
			if err := login(cmd, args); err != nil {
				log.Fatal(err)
			}

			token, err = readToken()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
	client = auth.NewClient(token)
}

func postRootCmd(cmd *cobra.Command, args []string) {
	if cmd.Use == "login" {
		return
	}
	currToken, err := client.Token()
	if err != nil {
		log.Fatal(err)
	}
	if currToken != token {
		if err := storeToken(currToken); err != nil {
			log.Fatal(err)
		}
	}
}

func storeToken(tok *oauth2.Token) error {
	f, err := os.OpenFile(tokenPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	return enc.Encode(tok)
}

func readToken() (*oauth2.Token, error) {
	content, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	var tok oauth2.Token
	if err = json.Unmarshal(content, &tok); err != nil {
		return nil, err
	}
	return &tok, nil
}
