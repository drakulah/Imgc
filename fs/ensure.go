package fs

import (
	"os"
	"path/filepath"
)

func EnsureDir(path string) error {
	if ExistsDir(path) {
		return nil
	}

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func EnsureFile(path string) error {
	if ExistsFile(path) {
		return nil
	}
	dir := filepath.Dir(path)
	dirErr := EnsureDir(dir)
	if dirErr != nil {
		return dirErr
	}
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}
