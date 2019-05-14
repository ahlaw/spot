package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(volCmd)
}

var volCmd = &cobra.Command{
	Use:   "vol [up|down|amount]",
	Short: "Show or set current volume",
	RunE:  vol,
}

func vol(cmd *cobra.Command, args []string) error {
	state, err := client.PlayerState()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		fmt.Printf("Current volume is %d.\n", state.Device.Volume)
		return nil
	}

	var newVolume int
	switch reqVolume := args[0]; reqVolume {
	case "up":
		if len(args) == 2 {
			upVolume := args[1]
			volumeIncrease, err := strconv.Atoi(upVolume)
			if err != nil {
				return err
			}
			newVolume = state.Device.Volume + volumeIncrease
		} else {
			newVolume = state.Device.Volume + 5
		}
	case "down":
		if len(args) == 2 {
			downVolume := args[1]
			volumeDecrease, err := strconv.Atoi(downVolume)
			if err != nil {
				return err
			}
			newVolume = state.Device.Volume - volumeDecrease
		} else {
			newVolume = state.Device.Volume - 5
		}
	default:
		newVolume, err = strconv.Atoi(reqVolume)
		if err != nil {
			return err
		}
	}

	if newVolume < 0 {
		newVolume = 0
	} else if newVolume > 100 {
		newVolume = 100
	}

	return client.Volume(newVolume)
}
