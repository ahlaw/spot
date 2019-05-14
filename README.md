# Spot

[![](https://img.shields.io/github/license/ahlaw/spot.svg)](LICENSE.md)


`spot` is a CLI app for Spotify powered by [Spotify's Web API](https://developer.spotify.com/documentation/web-api/).

It aims to make certain actions more efficient than on the official Spotify desktop application. 

## Requirements

You first need to create a [new app](https://developer.spotify.com/dashboard/applications).

Store its Client ID in `SPOT_ID` and Client Secret in
`SPOT_SECRET` environment variables.

## Installation

```
$ go get -u github.com/ahlaw/spot
```

## Usage

Note that the Spotify app must be open since it is required to talk to the Spotify API.

```
$ spot --help
Spot is a command-line interface for Spotify

Usage:
  spot [command]

Available Commands:
  help        Help about any command
  history     Show recently played tracks
  info        Show Spotify info
  login       Login Spotify credentials
  next        Play next track
  pause       Pause Spotify
  play        Play Spotify
  prev        Play previous track
  share       Copy Open Spotify URL of current track
  top         Show top played artists by user
  vol         Show or set current volume

Flags:
  -h, --help   help for spot

Use "spot [command] --help" for more information about a command.
```

## Built With

* [cobra](https://github.com/spf13/cobra/) - Go CLI framework
* [zmb3/spotify](https://github.com/zmb3/spotify/) - Go wrapper for Spotify Web API

## License

[MIT License](LICENSE.md)

## Acknowledgments

* Inspired (heavily stolen) from [spotctl](https://github.com/jingweno/spotctl/), particularly the method for saving access tokens
