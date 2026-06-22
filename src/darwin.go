//go:build darwin

package drivedetector

import (
	"os"
	"path/filepath"
)

func Detect() ([]string, error) {
	var drives []string

	entries, err := os.ReadDir("/Volumes")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		path := filepath.Join("/Volumes", entry.Name())

		if _, err := os.Stat(path); err == nil {
			drives = append(drives, path)
		}
	}

	if _, err := os.Stat("/"); err == nil {
		drives = append([]string{"/"}, drives...)
	}

	return drives, nil
}
