package binary

import (
	"io/ioutil"
	"testing"
)

func TestFile(t *testing.T) {
	isBinary, err := File("testdata/exe1")
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, err = File("testdata/exe2")
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, err = File("testdata/conf1")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = File("testdata/conf2")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = File("testdata/fstab")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, err = File("testdata/hai")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}
}

// binaryFile is a helper function for testing the Data function
func binaryFile(filename string) (bool, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}
	return Data(data), nil
}

func TestData(t *testing.T) {
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
