package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	source, destination := inputFlag()
	fmt.Println("Moving files from `", source, "` to `", destination, "`")
	if isFolderExist(source) && isFolderExist(destination) {
		files := getListFiles(source)
		for _, file := range files {
			re := regexp.MustCompile(`.*\\|.*\/`)
			fileName := re.ReplaceAllString(file, "")
			dateFolder := destination + "/" + getFileDate(file)
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
		fmt.Println("Folder not found!")
	}
}

func inputFlag() (string, string) {
	source := flag.String("s", "", "Source folder")
	destination := flag.String("d", "", "Destination folder")
	flag.Parse()
	if *source == "" || *destination == "" {
		*source, *destination = inputManual()
	}
	return *source, *destination
}

func inputManual() (string, string) {
	var source, destination string
	fmt.Print("Enter source folder: ")
	fmt.Scanln(&source)
	fmt.Print("Enter destination folder: ")
	fmt.Scanln(&destination)
	return source, destination
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
