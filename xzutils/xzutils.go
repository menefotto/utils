package xzutils

import (
	"bytes"
	"io"
	"os"

	xz "github.com/remyoudompheng/go-liblzma"
)

func Decompress(data interface{}) ([]byte, error) {
	var inbuffer *bytes.Buffer

	switch data.(type) {
	case string:
		inbuffer = bytes.NewBufferString(data.(string))
	case []byte:
		inbuffer = bytes.NewBuffer(data.([]byte))
	}

	decompressed, err := xz.NewReader(inbuffer)
	if err != nil {
		return nil, err
	}

	outbuffer := new(bytes.Buffer)
	_, err = io.Copy(outbuffer, decompressed)
	if err != nil {
		return nil, err
	}

	err = decompressed.Close()
	if err != nil {
		return nil, err
	}

	return outbuffer.Bytes(), nil
}

func FileDecompress(filename string) ([]byte, error) {
	data, err := openAndRead(filename)
	if err != nil {
		return nil, err
	}

	bytes, err := Decompress(data)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Compress(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	compressor, err := xz.NewWriter(&buffer, xz.LevelDefault)
	if err != nil {
		return nil, err
	}

	switch data.(type) {
	case string:
		if _, err := compressor.Write([]byte(data.(string))); err != nil {
			return nil, err
		}
	case []byte:
		if _, err := compressor.Write(data.([]byte)); err != nil {
			return nil, err
		}
	}

	err = compressor.Close()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func FileCompress(filein, fileout string) error {
	data, err := openAndRead(filein)
	if err != nil {
		return err
	}

	compressed, err := Compress(data)
	if err != nil {
		return err
	}

	fout, err := os.Create(fileout)
	if err != nil {
		return err
	}
	defer fout.Close()

	_, err = fout.Write(compressed)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func openAndRead(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, fstat.Size())
	_, err = f.Read(data)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return data, nil

}
