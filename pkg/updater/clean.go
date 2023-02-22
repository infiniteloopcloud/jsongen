package updater

import (
	"os"
	"strings"
)

func Clean(dest string) error {
	dir, err := os.ReadDir(dest)
	if err != nil {
		return err
	}
	for _, entry := range dir {
		switch {
		case strings.Contains(entry.Name(), ".go"):
			if err := os.RemoveAll(dest + "/" + entry.Name()); err != nil {
				return err
			}
		case strings.Contains(entry.Name(), "testdata"):
			if err := os.RemoveAll(dest + "/" + entry.Name()); err != nil {
				return err
			}
		}
	}
	return nil
}
