package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const directory = "."
const pathSeparator = "/"
const fileExtension = ".md"
const filePermissions = 0644

const whitespace = " "
const tagPrefix = "[["
const tagSuffix = "]]"

func main() {

	fileNames, err := getFileNames(directory, fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	tags := getTagsFromFileNames(fileNames, fileExtension)
	for _, value := range tags {
		fmt.Printf(value)
	}

	for _, fileName := range fileNames {

		fileContent, err := ioutil.ReadFile(directory + pathSeparator + fileName)
		if err != nil {
			log.Fatal("unable to read file " + fileName)
		}

		updatedFileContent := updateTagsInFile(tags, fileName, strings.TrimSpace(string(fileContent)))

		err = ioutil.WriteFile(directory+pathSeparator+fileName, []byte(updatedFileContent), filePermissions)
		if err != nil {
			log.Fatal("unable to write file " + fileName)
		}

	}
}

func updateTagsInFile(tags []string, fileName string, fileContent string) string {
	parts := strings.Split(fileContent, whitespace)
	for _, tag := range tags {
		for index, part := range parts {
			log.Println("part = " + part + " tag = " + tag)
			if part == tag && !strings.Contains(fileName, tag) {
				parts[index] = tagPrefix + part + tagSuffix
			}
		}
	}
	return strings.Join(parts, whitespace)
}

func getTagsFromFileNames(fileNames []string, fileExtension string) []string {
	var tokens []string
	for _, fileName := range fileNames {
		tokens = append(tokens, strings.Replace(fileName, fileExtension, "", -1))
	}
	return tokens
}

func getFileNames(directory string, fileExtension string) ([]string, error) {
	var fileNames []string

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), fileExtension) {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, nil
}
