package util

import (
	"os"

	gap "github.com/muesli/go-app-paths"
)

func GetApplicationDataDir() (string, error) {
	scope := gap.NewScope(gap.User, "tilcli")
	dirs, err := scope.DataDirs()
	dir := dirs[0]
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0o770); err != nil {
		return "", err
	}
	return dir, nil
}
