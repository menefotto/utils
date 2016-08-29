package tarutils

import (
	"archive/tar"
	"errors"
	"io"
	"os"
)

var ErrNotTarFile = errors.New("This is not a valid tar archive")

func IsTarFile(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	ar := tar.NewReader(file)
	h, err := ar.Next()
	if h == nil {
		return false, ErrNotTarFile
	}
	return true, nil
}

func TarExtractor(filename string) (map[string][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ar := tar.NewReader(file)
	data := make(map[string][]byte, 0)

	for {
		header, err := ar.Next()
		if err == io.EOF {
			break
		}
		if err != io.EOF && err != nil {
			return nil, err
		}

		content := make([]byte, header.Size)
		_, err = ar.Read(content)
		if err != io.EOF && err != nil {
			return nil, err
		}

		data[header.Name] = content
	}

	return data, nil

}
