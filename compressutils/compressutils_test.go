package compressutils

import (
	"os"
	"testing"
)

func TestFileCompressionByte(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	dat := make([]byte, stat.Size())
	_, err = f.Read(dat)
	if err != nil {
		t.Fatal(err)
	}

	data, err := Compress(dat)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 772 {
		t.Fatal("Size must be equal to 772, instead was: ", len(data))
	}
}
func TestFileCompressionWrong(t *testing.T) {
	f, err := os.Open("malformed.txt.xz")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	dat := make([]byte, stat.Size())
	_, err = f.Read(dat)
	if err != nil {
		t.Fatal(err)
	}

	data, err := Compress(dat)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 772 {
		t.Fatal("Size must be equal to 772, instead was: ", len(data))
	}
}

func TestFileCompressionNoEsist(t *testing.T) {
	f, err := os.Open("nofilehere.txt")
	if err == nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err == nil {
		t.Fatal(err)
		return
	} else {
		return
	}

	dat := make([]byte, stat.Size())
	_, err = f.Read(dat)
	if err != nil {
		t.Fatal(err)
	}

	data, err := Compress(dat)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 772 {
		t.Fatal("Size must be equal to 772, instead was: ", len(data))
	}
}
func TestFileCompressionString(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	dat := make([]byte, stat.Size())
	_, err = f.Read(dat)
	if err != nil {
		t.Fatal(err)
	}

	data, err := Compress(string(dat))
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 772 {
		t.Fatal("Size must be equal to 772, instead was: ", len(data))
	}
}

func TestXzDecompressBytes(t *testing.T) {
	f, err := os.Open("test.txt.xz")
	if err != nil {
		t.Error(err)
	}

	info, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		t.Error(err)
	}

	b, err := Decompress(data)
	if err != nil {
		t.Log(err)
	}

	if len(b) != 1203 {
		t.Fatalf("Size should have been 1204 instead is: ", len(b))
	}
}

func TestXzDecompressBytesMalformed(t *testing.T) {
	f, err := os.Open("malformed.txt.xz")
	if err != nil {
		t.Error(err)
	}

	info, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		t.Error(err)
	}

	b, err := Decompress(data)
	if err == nil {
		t.Log(err)
	}

	if len(b) == 1203 {
		t.Fatalf("Size should have been 1204 instead is: ", len(b))
	}
}
func TestXzDecompressNoFile(t *testing.T) {
	f, err := os.Open("")
	if err == nil {
		t.Fatal(err)
	}

	info, err := f.Stat()
	if err == nil {
		t.Error(err)
		return
	} else {
		return
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		t.Error(err)
	}

	b, err := Decompress(data)
	if err != nil {
		t.Log(err)
	}

	if len(b) != 1203 {
		t.Fatalf("Size should have been 1204 instead is: ", len(b))
	}
}

func TestDecompressWrongInvalidInput(t *testing.T) {
	data := ""
	_, err := Decompress(data)
	if err == nil {
		t.Log(err)
	}

}
func TestXzDecompressString(t *testing.T) {
	f, err := os.Open("test.txt.xz")
	if err != nil {
		t.Error(err)
	}

	info, err := f.Stat()
	if err != nil {
		t.Error(err)
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		t.Error(err)
	}

	b, err := Decompress(string(data))
	if err != nil {
		t.Log(err)
	}

	if len(b) != 1203 {
		t.Fatalf("Size should have been 1204 instead is: ", len(b))
	}
}

func TestFileXzDecompress(t *testing.T) {
	data, err := FileDecompress("test.txt.xz")
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 1203 {
		t.Fatal("Decompression went wrong, size is: ", len(data))
	}
}

func TestFileXzCompress(t *testing.T) {
	err := FileCompress("test.txt", "test.xz")
	if err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat("test.xz")
	if err != nil {
		t.Fatal(err)
	}

	if info.Size() != int64(772) {
		t.Fatal("Compression went wrong, size missmatch: ", info.Size())
	}

	os.Remove("test.xz")
}

func TestFileXzCompressWrong(t *testing.T) {
	err := FileCompress("test.txt", "/var/cache/test.xz")
	if err == nil {
		t.Fatal(err)
	}

	_, err = os.Stat("/var/cache/test.xz")
	if err == nil {
		t.Fatal(err)
	}

}

func TestFileXzCompressWrongInput(t *testing.T) {
	err := FileCompress("boh", "test.xz")
	if err == nil {
		t.Fatal(err)
	}

	_, err = os.Stat("test.xz")
	if err == nil {
		t.Fatal(err)
	}

}
