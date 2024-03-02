package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	source := flag.String("s", "", "Source")
	destination := flag.String("o", "", "Destination")
	flag.Parse()
	fmt.Println("Source:", *source)
	fmt.Println("Destination:", *destination)
	if isFolderExist(*source) && isFolderExist(*destination) {
		files := getListFiles(*source)
		for _, file := range files {
			re := regexp.MustCompile(`.*\\|.*\/`)
			fileName := re.ReplaceAllString(file, "")
			dateFolder := *destination + "/" + getFileDate(file)
			destinationFile := dateFolder + "/" + fileName
			for isFileExist(destinationFile) {
				re := regexp.MustCompile(`(.+?)(\.[^.]*$|$)`)
				fileName = re.ReplaceAllString(fileName, "$1(1)$2")
				destinationFile = dateFolder + "/" + fileName
			}
			createFolder(dateFolder)
			moveFile(file, destinationFile)
		}
	} else {
		fmt.Println("error")
	}
}

func isFolderExist(folder string) bool {
	_, err := os.Stat(folder)
	return err == nil
}

func isFileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func getListFiles(pathSource string) []string {
	var files []string
	err := filepath.Walk(pathSource, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return files
}

func getFileDate(path string) string {
	file, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	return file.ModTime().Format("2006-01-02")
}

func moveFile(source string, destination string) {
	err := os.Rename(source, destination)
	if err != nil {
		fmt.Println(err)
	}
}

func createFolder(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
