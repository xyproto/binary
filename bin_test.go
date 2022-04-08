package bin

import (
	"testing"
)

func TestBinary1(t *testing.T) {
	isBinary, err := BinaryFile("testdata/exe1")
	if err != nil {
		t.Fatal(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/exe2")
	if err != nil {
		t.Fatal(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/conf1")
	if err != nil {
		t.Fatal(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/conf2")
	if err != nil {
		t.Fatal(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/fstab")
	if err != nil {
		t.Fatal(err)
	}
	if isBinary {
		t.FailNow()
	}

}
