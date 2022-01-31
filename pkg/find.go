package pkg

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func findFilesRecursivly(directory string, suffix string, currentList []string) ([]string, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return currentList, err
	}

	for _, f := range files {
		if f.IsDir() {
			currentList, err = findFilesRecursivly(filepath.FromSlash(directory+"/"+f.Name()), suffix, currentList)
			if err != nil {
				return nil, err
			}
		} else if strings.HasSuffix(strings.ToLower(f.Name()), suffix) {
			currentList = append(currentList, filepath.FromSlash(directory+"/"+f.Name()))
		}
	}
	return currentList, nil
}

func FindFiles(diretory string, suffix string) ([]string, error) {
	list := make([]string, 0)
	return findFilesRecursivly(diretory, suffix, list)
}
