package appdata

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// limpa placeholders do Exec como %U, %f, %F, %i etc
func cleanExecCommand(exec string) string {
	fields := strings.Fields(exec)
	clean := []string{}

	for _, field := range fields {
		if !strings.HasPrefix(field, "%") {
			clean = append(clean, field)
		}
	}
	return strings.Join(clean, " ")
}

func LoadApps() map[string]string {
	files, err := filepath.Glob("/usr/share/applications/*.desktop")
	check(err)
	apps := make(map[string]string)
	for _, file := range files {
		func() {
			fileContent, err := os.Open(file)
			check(err)
			defer fileContent.Close()

			var name, exec, kind string
			scanner := bufio.NewScanner(fileContent)
			for scanner.Scan() {
				line := scanner.Text()

				if after, ok := strings.CutPrefix(line, "Name="); ok {
					name = after
				}

				if after, ok := strings.CutPrefix(line, "Exec="); ok {
					exec = cleanExecCommand(after)
				}

				if after, ok := strings.CutPrefix(line, "Type="); ok {
					kind = after
				}

				if name != "" && exec != "" && kind == "Application" {
					apps[name] = exec
					break
				}
			}
		}()
	}
	return apps
}

