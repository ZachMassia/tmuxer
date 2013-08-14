package main

import (
	"fmt"
	"strconv"
)

const (
	heartFull  = "♥" // Charging / Fully Charged
	heartEmpty = "♡" // Discharging

	batPath   = "/sys/class/power_supply/BAT0/"
	batStatus = batPath + "status"
	batCap    = batPath + "energy_full"
	batNow    = batPath + "energy_now"
)

func init() { register("bat", batOutput) }

func batOutput() (string, error) {
	// Get battery status. Possible values:
	//       Charging, Discharging, Unknown
	// Unknown seems to occur when plugged in and fully charged.
	status, err := getVal(batStatus)
	if err != nil {
		return "", err
	}

	// Read battery stats and convert to ints
	var remaining, capacity float32
	if bNow, err := getVal(batNow); err == nil {
		if bNowI, err := strconv.Atoi(bNow); err != nil {
			return "", err
		} else {
			remaining = float32(bNowI)
		}
	}
	if bCap, err := getVal(batCap); err == nil {
		if bCapI, err := strconv.Atoi(bCap); err != nil {
			return "", err
		} else {
			capacity = float32(bCapI)
		}
	}

	// Calculate battery charged percentage.
	charge := remaining / capacity * 100

	// Use a full heart when charging or fully charged
	var heart string
	if status == "Charging" || status == "Unknown" {
		heart = heartFull
	} else {
		heart = heartEmpty
	}
	return fmt.Sprintf("%s %.1f", heart, charge), nil
}
