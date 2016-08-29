//this package containes files and related utilities
// functions like DirList with returns a slice a names
// with the content in the directory given as parameter
// is it not recurisive, it will not walk all the
// directory tree.

package futils

import (
	"io/ioutil"
	"os"
)

func DirList(dirname string) ([]string, error) {
	filelist := make([]string, 0)

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return filelist, err
	}

	for _, file := range files {
		filelist = append(filelist, file.Name())
	}

	return filelist, nil
}

func WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, os.ModePerm)
}

func WriteFileOrDir(name, data []byte) error {
	if len(data) == 0 {
		return os.MkdirAll(name, os.ModePerm)
	}

	return WriteFile(name, data)
}
