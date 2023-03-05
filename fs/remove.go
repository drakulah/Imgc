package fs

import "os"

func RemoveFile(path string) error {
	if ExistsFile(path) {
		return os.Remove(path)
	}
	return nil
}

func RemoveDir(path string) error {
	if ExistsDir(path) {
		return os.Remove(path)
	}
	return nil
}

func RemoveDirRecursively(path string) error {
	if ExistsDir(path) {
		return os.RemoveAll(path)
	}
	return nil
}
