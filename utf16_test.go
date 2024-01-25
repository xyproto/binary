package binary

import (
	"log"
	"os"
	"testing"
)

func TestFileAndUTF16(t *testing.T) {
	filename := "testdata/exe1"
	isBinary, _, err := FileAndUTF16(filename)
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		log.Println("not detected as binary: " + filename)
		t.FailNow()
	}
	filename = "testdata/exe2"
	isBinary, _, err = FileAndUTF16(filename)
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		log.Println("not detected as binary: " + filename)
		t.FailNow()
	}
	filename = "testdata/conf1"
	isBinary, _, err = FileAndUTF16(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}
	filename = "testdata/conf2"
	isBinary, _, err = FileAndUTF16(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}
	filename = "testdata/fstab"
	isBinary, _, err = FileAndUTF16(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}
	filename = "testdata/hai"
	isBinary, _, err = FileAndUTF16(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}
	filename = "testdata/utf16.csv"
	var isUtf16 bool
	isBinary, isUtf16, err = FileAndUTF16(filename)
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		log.Println("not detected as text: " + filename)
		t.FailNow()
	}
	if !isUtf16 {
		log.Println("not detected as UTF-16: " + filename)
		t.FailNow()
	}

}

// binaryFileAndUTF16 is a helper function for testing the Data function
func binaryFileAndUTF16(filename string) (bool, bool, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false, false, err
	}
	isBinary, isUtf16 := DataAndUTF16(data)
	return isBinary, isUtf16, nil
}

func TestDataAndUTF16(t *testing.T) {
	isBinary, _, err := binaryFileAndUTF16("testdata/exe1")
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, _, err = binaryFileAndUTF16("testdata/exe2")
	if err != nil {
		t.Error(err)
	}
	if !isBinary {
		t.FailNow()
	}

	isBinary, _, err = binaryFileAndUTF16("testdata/conf1")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, _, err = binaryFileAndUTF16("testdata/conf2")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, _, err = binaryFileAndUTF16("testdata/fstab")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	isBinary, _, err = binaryFileAndUTF16("testdata/hai")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}

	var isUtf16 bool
	isBinary, isUtf16, err = binaryFileAndUTF16("testdata/utf16.csv")
	if err != nil {
		t.Error(err)
	}
	if isBinary {
		t.FailNow()
	}
	if !isUtf16 {
		t.FailNow()
	}
}
