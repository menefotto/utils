package compressutils

import (
	"bytes"
	"io"
	"os"

	"github.com/golang/snappy"
	xz "github.com/remyoudompheng/go-liblzma"
	"github.com/sonic/lib/errors"
)

func XzDecompress(data interface{}) ([]byte, error) {
	var inbuffer *bytes.Buffer

	switch data.(type) {
	case string:
		inbuffer = bytes.NewBufferString(data.(string))
	case []byte:
		inbuffer = bytes.NewBuffer(data.([]byte))
	}

	decompressed, err := xz.NewReader(inbuffer)
	if err != nil {
		return nil, errors.Wrap(err)()
	}

	outbuffer := new(bytes.Buffer)
	_, err = io.Copy(outbuffer, decompressed)
	if err != nil {
		return nil, errors.Wrap(err)()
	}

	err = decompressed.Close()
	if err != nil {
		return nil, errors.Wrap(err)()
	}

	return outbuffer.Bytes(), nil
}

func XzFileDecompress(filename string) ([]byte, error) {
	data, err := openAndRead(filename)
	if err != nil {
		return nil, err
	}

	bytes, err := XzDecompress(data)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func XzCompress(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	compressor, err := xz.NewWriter(&buffer, xz.LevelDefault)
	if err != nil {
		return nil, errors.Wrap(err)()
	}

	switch data.(type) {
	case string:
		if _, err := compressor.Write([]byte(data.(string))); err != nil {
			return nil, errors.Wrap(err)()
		}
	case []byte:
		if _, err := compressor.Write(data.([]byte)); err != nil {
			return nil, errors.Wrap(err)()
		}
	}

	err = compressor.Close()
	if err != nil {
		return nil, errors.Wrap(err)()
	}

	return buffer.Bytes(), nil
}

func XzFileCompress(filein, fileout string) error {
	data, err := openAndRead(filein)
	if err != nil {
		return err
	}

	compressed, err := XzCompress(data)
	if err != nil {
		return err
	}

	fout, err := os.Create(fileout)
	if err != nil {
		return errors.Wrap(err)()
	}
	defer fout.Close()

	_, err = fout.Write(compressed)
	if err != nil && err != io.EOF {
		return errors.Wrap(err)()
	}

	return nil
}

func openAndRead(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err)()
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		return nil, errors.Wrap(err)()
	}

	data := make([]byte, fstat.Size())
	_, err = f.Read(data)
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err)()
	}

	return data, nil

}

func SnappyCompress(src []byte) []byte {
	return snappy.Encode([]byte(""), src)
}

func SnappyDecompress(src []byte) ([]byte, error) {
	blob, err := snappy.Decode([]byte(""), src)
	if err != nil {
		return nil, errors.Wrap(err)()
	}
	return blob, nil
}
