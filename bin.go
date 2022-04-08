package bin

import (
	"bytes"
	"os"
	"unicode/utf8"
)

func probablyBinary24(b []byte) bool {
	zeroCount := bytes.Count(b, []byte{0})
	if zeroCount > len(b)/3 {
		// Suspiciously many zero bytes, more than a third of them.
		// Likely to be binary data.
		return true
	}
	return !utf8.ValidString(string(b))
}

func probablyBinary(first24, middle24, last24 []byte) bool {
	// 	fmt.Printf("first24 %v %s\n", first24, string(first24))
	// 	fmt.Printf("middle24 %v %s\n", middle24, string(middle24))
	// 	fmt.Printf("last24 %v %s\n", last24, string(last24))
	return probablyBinary24(first24) || probablyBinary24(middle24) || probablyBinary24(last24)
}

// BinaryFile tries to determine if the given filename is a binary file by reading the first, last
// and middle 24 bytes, counting the zero bytes and trying to convert the data to utf8.
func BinaryFile(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	const EndOfFile = 2
	const CurrentPosition = 1
	const StartOfFile = 0

	lastPositionMinus24, err := file.Seek(-24, EndOfFile)
	if err != nil {
		return false, err
	}

	last24 := make([]byte, 24)

	// Read 24 bytes
	last24count, err := file.Read(last24)
	if err != nil {
		return false, err
	}

	if last24count == 0 {
		// No data, decide it's not a binary
		return false, nil
	}

	// Are there enough bytes left to read the 24 last ones and the 24 center ones as well?
	if lastPositionMinus24 >= 48 {

		first24 := make([]byte, 24)
		first24count := 0
		middle24 := make([]byte, 24)
		middle24count := 0

		// Go to the start of the file
		_, err := file.Seek(0, StartOfFile)
		if err != nil {
			return false, err
		}

		// Read 24 bytes
		first24count, err = file.Read(first24)
		if err != nil {
			return false, err
		}

		centerPos := (lastPositionMinus24 + 24) / 2

		// Go to the center of the file
		_, err = file.Seek(centerPos, CurrentPosition)
		if err != nil {
			return false, err
		}

		// Read 24 bytes
		middle24count, err = file.Read(middle24)
		if err != nil {
			return false, err
		}

		if first24count == 0 && middle24count == 0 {
			// This should never happen
			return probablyBinary24(last24), nil
		}

		return probablyBinary(first24, middle24, last24), nil
	}

	return probablyBinary24(last24), nil

}
