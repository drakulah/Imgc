package fs

import "os"

func WriteFile(path string, data []byte) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	_, err1 := file.Write(data)

	if err != nil {
		return err1
	}

	return nil
}

func WriteTextFile(path string, data string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	_, err1 := file.WriteString(data)

	if err != nil {
		return err1
	}

	return nil
}
