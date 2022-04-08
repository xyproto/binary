package bin

import (
	"bytes"
	"os"
	"unicode/utf8"
)

func probablyBinaryData(b []byte) bool {
	//fmt.Printf("%v\n", b)
	zeroCount := bytes.Count(b, []byte{0})
	if zeroCount > len(b)/3 {
		// Suspiciously many zero bytes; more than a third of them.
		return true
	}
	return !utf8.ValidString(string(b))
}

// BinaryFile tries to determine if the given filename is a binary file by reading the first, last
// and middle 24 bytes, counting the zero bytes and trying to convert the data to utf8.
func BinaryFile(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Go to the end of the file, minus 24 bytes
	fileLength, err := file.Seek(-24, os.SEEK_END)
	if err != nil || fileLength < 24 {
		//fmt.Println(filename, "could not seek to -24 and/or the file is too short")

		// Go to the start of the file, ignore errors
		_, err = file.Seek(0, os.SEEK_SET)

		// Read up to 24 bytes
		fileBytes := make([]byte, 24)
		n, err := file.Read(fileBytes)
		if err != nil {
			// Could not read the file
			return false, err
		} else if n == 0 {
			// The file is too short, decide it's a text file
			return false, nil
		}
		fileBytes = fileBytes[:n]

		// Check if it's likely to be a binary file, based on the few available bytes
		return probablyBinaryData(fileBytes), nil
	}

	last24 := make([]byte, 24)

	// Read 24 bytes
	last24count, err := file.Read(last24)
	if err != nil {
		return false, err
	}
	// Shorten the byte slice
	//last24 = last24[:last24count]

	// fmt.Printf("last24 %v %s\n", last24, string(last24))
	if last24count > 0 && probablyBinaryData(last24) {
		return true, nil
	}

	if fileLength-24 >= 24 {
		first24 := make([]byte, 24)
		first24count := 0

		// Go to the start of the file
		if _, err := file.Seek(0, os.SEEK_SET); err != nil {
			// Could not go to the start of the file (!)
			return false, err
		}

		// Read 24 bytes
		first24count, err = file.Read(first24)
		if err != nil {
			return false, err
		}
		// Shorten the byte slice
		//first24 = first24[:first24count]

		//fmt.Printf("first24 %v %s\n", first24, string(first24))
		if first24count > 0 && probablyBinaryData(first24) {
			return true, nil
		}
	}

	if fileLength-24 >= 48 {

		middle24 := make([]byte, 24)
		middle24count := 0

		middlePos := fileLength / 2

		// Go to the middle of the file, relative to the start. Ignore errors.
		_, _ = file.Seek(middlePos, os.SEEK_SET)

		// Read 24 bytes from where MIGHT be the middle of the file
		middle24count, err = file.Read(middle24)
		if err != nil {
			return false, err
		}
		// Shorten the byte slice
		//middle24 = middle24[:middle24count]

		// fmt.Printf("middle24 %v %s\n", middle24, string(middle24))
		return middle24count > 0 && probablyBinaryData(middle24), nil
	}

	// If it was a binary file, it should have been catched by one of the returns above
	return false, nil
}
