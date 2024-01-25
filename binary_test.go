package binary

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestFile(t *testing.T) {
	filename := "testdata/exe1"
	isBinary, err := File(filename)
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		log.Println("not detected as binary: " + filename)
		t.FailNow()
	}

	filename = "testdata/exe2"
	isBinary, err = File(filename)
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		log.Println("not detected as binary: " + filename)
		t.FailNow()
	}

	filename = "testdata/conf1"
	isBinary, err = File(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}

	filename = "testdata/conf2"
	isBinary, err = File(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}

	filename = "testdata/fstab"
	isBinary, err = File(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}

	filename = "testdata/hai"
	isBinary, err = File(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}

	filename = "testdata/utf16.csv"
	isBinary, err = File(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as UTF16 text: " + filename)
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
