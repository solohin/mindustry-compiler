package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/solohin/mindustry-compiler/pkg/mlog"
)

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.IsDir() {
		log.Println("Compiling directory " + path)
		err := compileDir(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	if !fileInfo.IsDir() {
		err := compileFile(path)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func compileFile(filename string) error {
	log.Println("Compiling file " + filename)

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	newText, err := mlog.Compile(string(fileBytes))
	if err != nil {
		return err
	}

	newFileName, err := getNewFilename(filename)
	if err != nil {
		return err
	}

	err = os.WriteFile(newFileName, []byte(newText), 0644)
	if err != nil {
		return err
	}

	return nil
}

func compileDir(dirname string) error {
	log.Println("Looking for all *.mlog files in directory " + dirname)

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".mlog") {
			continue
		}

		err := compileFile(filepath.Join(dirname, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func getNewFilename(filename string) (string, error) {
	dirName := filepath.Join(filepath.Dir(filename), "compiled")

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return "", err
	}

	return filepath.Join(dirName, filepath.Base(filename)), nil
}
