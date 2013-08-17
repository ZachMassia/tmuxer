package main

import (
	"fmt"
	mpris "github.com/lann/mpris2-go"
)

func init() { register("nowplaying", nowPlaying) }

func nowPlaying() (string, error) {
	// Try to find a player
	player, err := getPlayer()
	if err != nil {
		return "", err
	}

	// Fetch metadata.
	md, err := player.Metadata()
	if err != nil {
		return "", fmt.Errorf("Could not get metadata: %s", err)
	}

	// Make sure title and artist not empty.
	title := md.Title()
	if title == "" {
		return "", fmt.Errorf("No title")
	}
	artists := md.Artists()
	if len(artists) == 0 {
		return "", fmt.Errorf("No artists")
	}

	return fmt.Sprintf("♫ %s - %s", title, artists[0]), nil
}

// getPlayer opens a connection and returns the first available object
// implementing the MediaPlayer2 interface.
func getPlayer() (p *mpris.MediaPlayer, err error) {
	// Connect to DBus.
	conn, err := mpris.Connect()
	if err != nil {
		return nil, fmt.Errorf("Could not connect to DBus: %s", err)
	}
	// Try and get a MediaPlayer2 object
	p, err = conn.GetAnyMediaPlayer()
	if err != nil {
		return nil, fmt.Errorf("Could not get a player: %s", err)
	}
	return p, nil
}
