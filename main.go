package main

import (
	"wifi-go/appdata"
	"wifi-go/ui"
)

func main() {
	apps := appdata.LoadApps()
	ui.InitUI(apps)
}
