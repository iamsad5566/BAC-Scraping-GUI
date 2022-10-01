package main

import (
	"bac-scraper-gui/config"
	"bac-scraper-gui/gui"
)

func main() {
	// Load the env variables first
	envar := config.LoadYml()
	gui.OpenWindow(envar)
}
