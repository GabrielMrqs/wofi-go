package main

import (
	"wifi-go/appdata"
	"wifi-go/ui"
	"os"
	"fmt"
)

func main() {
	apps, err := appdata.LoadApps()
  if err != nil {
  	fmt.Fprintln(os.Stderr, "Error loading applications:", err)
  	os.Exit(1)
  }

	ui.InitUI(apps)
}
