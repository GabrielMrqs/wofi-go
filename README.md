# wofi-go

**wofi-go** is a lightweight graphical launcher written in Go using GTK (gotk3), inspired by the `wofi` menu. It scans `.desktop` files in `/usr/share/applications` and displays a list of available applications to launch.

## Features

- GTK3-based GUI
- Parses all valid `.desktop` files
- Lightweight and dependency-free (beyond Go and GTK)

## Prerequisites

To build and run this project, you need:

- Go 1.20 or newer
- GTK development packages installed

On Fedora/Nobara:

```bash
sudo dnf install gtk3-devel glib2-devel cairo-devel pango-devel gdk-pixbuf2-devel
