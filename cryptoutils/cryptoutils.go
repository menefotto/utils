package miscutils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

func FileSha256Sum(filename string) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return "", nil
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", nil

	}

	return fmt.Sprintf("%x", sha256.Sum256(content)), nil

}

func FileMd5Sum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", md5.Sum(content)), nil

}
