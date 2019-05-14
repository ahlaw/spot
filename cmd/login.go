package cmd

import (
	"fmt"
	"log"
	"net/http"

	rndm "github.com/nmrshll/rndm-go"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login Spotify credentials",
	RunE:  login,
}

func login(cmd *cobra.Command, args []string) error {
	state := rndm.String(16)
	ch := make(chan *oauth2.Token)

	http.Handle("/callback", &AuthHandler{state: state, ch: ch, auth: auth})
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	tok := <-ch
	if err := storeToken(tok); err != nil {
		return err
	}
	return nil
}

type AuthHandler struct {
	state string
	ch    chan *oauth2.Token
	auth  spotify.Authenticator
}

func (ah *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(ah.state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != ah.state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, ah.state)
	}
	fmt.Fprintf(w, "Login Completed! Please return to your terminal.")
	ah.ch <- tok
}
