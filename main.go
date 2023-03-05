package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"imgc/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	argsLen := len(os.Args)

	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		fmt.Println("Couldn't access current working directory!")
		return
	}

	if argsLen != 2 {
		fmt.Println("Invalid arguments!")
		return
	}

	dirPath := filepath.Join(cwd, os.Args[1])

	if !fs.ExistsDir(dirPath) {
		fmt.Printf("Directory `%s` not found!\n", dirPath)
		return
	}

	i := 1
	outDirPath := fmt.Sprintf("%s (%d)", dirPath, i)

	for {
		if !fs.ExistsDir(outDirPath) {
			break
		}
		outDirPath = fmt.Sprintf("%s (%d)", dirPath, i)
		i++
	}

	fmt.Printf("Creating output directory `%s`!\n", outDirPath)

	if createDirErr := fs.EnsureDir(outDirPath); createDirErr != nil {
		fmt.Printf("Couldn't create output directory `%s`!\n", outDirPath)
		return
	}

	fmt.Println("Entering the main loop!")
	for _, item := range fs.ReadDir(dirPath) {
		if item.IsDir {
			continue
		}

		fileExtension := strings.ToLower(filepath.Ext(item.Name))

		if !regexp.MustCompile(`^\.(png|jpeg|jpg)$`).MatchString(fileExtension) {
			continue
		}

		filePath := path.Join(dirPath, "./"+item.Name)
		outFilePath := path.Join(outDirPath, "./"+item.Name)
		file, fileErr := os.Open(filePath)

		if fileErr != nil {
			fmt.Printf("Couldn't open file `%s`\n", filePath)
			continue
		}

		defer file.Close()

		img, _, imgErr := image.Decode(file)
		if imgErr != nil {
			fmt.Printf("Couldn't decode image `%s`\n", filePath)
			continue
		}

		outputFile, outputFileErr := os.Create(outFilePath)
		if outputFileErr != nil {
			fmt.Printf("Couldn't create output file `%s`\n", outFilePath)
			continue
		}
		defer outputFile.Close()

		var encodeErr error

		if fileExtension == ".png" {
			encodeErr = png.Encode(outputFile, img)
		} else if fileExtension == ".jpg" || fileExtension == ".jpeg" {
			encodeErr = jpeg.Encode(outputFile, img, &jpeg.Options{
				Quality: 80,
			})
		}

		if encodeErr != nil {
			fmt.Printf("Couldn't encode image `%s`\n", filePath)
			continue
		}

		if stat, statErr := file.Stat(); statErr == nil {
			statErr = os.Chtimes(outFilePath, stat.ModTime(), stat.ModTime())
			if statErr != nil {
				fmt.Printf("Couldn't modify modification date of compressed file `%s`\n", outFilePath)
			}
		}
	}
	fmt.Println("Done compressing!")
}
