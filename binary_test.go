package binary

import (
	"io/ioutil"
	"testing"
)

func TestBinaryFile(t *testing.T) {
	isBinary, err := BinaryFile("testdata/exe1")
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/exe2")
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/conf1")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/conf2")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/fstab")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = BinaryFile("testdata/hai")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}
}

// binaryFile is a helper function for testing the BinaryData function
func binaryFile(filename string) (bool, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}
	return BinaryData(data), nil
}

func TestBinaryData(t *testing.T) {
	isBinary, err := binaryFile("testdata/exe1")
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, err = binaryFile("testdata/exe2")
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, err = binaryFile("testdata/conf1")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = binaryFile("testdata/conf2")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = binaryFile("testdata/fstab")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = binaryFile("testdata/hai")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

}
