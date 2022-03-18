package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	devdDir  = ".gen_form_template"
	devdHome = "DEVD_HOME"
)

func HomeDir() (string, error) {
	if home := os.Getenv(devdHome); len(home) > 0 {
		return home, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return getHomeDir(usr.HomeDir)
}

func getHomeDir(userHome string) (string, error) {
	dir := filepath.Join(userHome, devdDir)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0o755); err != nil {
			return "", fmt.Errorf("mkdir failed: %w", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("check existence of %s dir: %w", dir, err)
	}

	return dir, nil
}
