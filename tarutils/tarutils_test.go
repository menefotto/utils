package tarutils

import "testing"

func TestTarExtract(t *testing.T) {
	datamap, err := TarExtractor("i3lock-2.8-1-x86_64.pkg.tar")
	if err != nil {
		t.Fatal(err)
	}

	i := 0
	for k, v := range datamap {
		t.Logf("File path/name is: %s size is: %d\n", k, len(v))
		i++
	}

	if i != 16 {
		t.Fatalf("Entry count should have been 16 instead is: %d\n", i)
	}
}

func TestTarNoExist(t *testing.T) {
	_, err := TarExtractor("nonesiste.tar")
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
