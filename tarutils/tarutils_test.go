package tarutils

import "testing"

func TestTarExtract(t *testing.T) {
	datamap, err := FileExtractor("i3lock-2.8-1-x86_64.pkg.tar")
	if err != nil {
		t.Fatal(err)
	}

	i := 0
	for k, v := range datamap {
		t.Logf("File path/name is: %s size is: %d\n", k, len(v.Data))
		i++
	}

	if i != 16 {
		t.Fatalf("Entry count should have been 16 instead is: %d\n", i)
	}
}

func TestTarNoExist(t *testing.T) {
	_, err := FileExtractor("nonesiste.tar")
	if err == nil {
		t.Fatal("Testing error condition gone!")
	}
}

func TestIsTarFile(t *testing.T) {
	ok, err := IsTarFile("nontarfile.tar")
	if ok {
		t.Fatal(err)
	}
}
