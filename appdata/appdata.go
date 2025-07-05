package appdata

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

func LoadApps() (map[string]string, error) {
	files, err := filepath.Glob("/usr/share/applications/*.desktop")
	if err != nil {
		return nil, err
	}

	apps := make(map[string]string, len(files))
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		name, exec, kind := parseDesktopFile(data)
		if kind == "Application" && name != "" && exec != "" {
			apps[name] = exec
		}
	}
	return apps, nil
}

func parseDesktopFile(data []byte) (name, exec, kind string) {
	lines := bytes.SplitSeq(data, []byte{'\n'})
	for raw := range lines {
		line := string(raw)

		if v, ok := strings.CutPrefix(line, "Name="); ok {
			name = v
		}
		if v, ok := strings.CutPrefix(line, "Exec="); ok {
			exec = cleanExecCommand(v)
		}
		if v, ok := strings.CutPrefix(line, "Type="); ok {
			kind = v
		}
		if name != "" && exec != "" && kind == "Application" {
			break
		}
	}
	return
}

func cleanExecCommand(exec string) string {
	fields := strings.Fields(exec)
	clean := make([]string, 0, len(fields))
	for _, f := range fields {
		if !strings.HasPrefix(f, "%") {
			clean = append(clean, f)
		}
	}
	return strings.Join(clean, " ")
}
