package fs

import (
	"log"
	"os"
	"path/filepath"
)

func BinaryFilePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panic(err)
	}
	return dir
}
