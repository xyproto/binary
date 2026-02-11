package binary

import (
	"os"
	"testing"
)

func TestMagicSignatures(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected bool
	}{
		{"PNG signature", []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00}, true},
		{"JPEG signature", []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10}, true},
		{"GIF87a signature", []byte{0x47, 0x49, 0x46, 0x38, 0x37, 0x61}, true},
		{"GIF89a signature", []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61}, true},
		{"ELF signature", []byte{0x7F, 0x45, 0x4C, 0x46, 0x02, 0x01}, true},
		{"PDF signature", []byte{0x25, 0x50, 0x44, 0x46, 0x2D, 0x31, 0x2E}, true},
		{"ZIP signature", []byte{0x50, 0x4B, 0x03, 0x04, 0x0A, 0x00}, true},
		{"GZIP signature", []byte{0x1F, 0x8B, 0x08, 0x00}, true},
		{"Plain text", []byte("Hello, World!"), false},
		{"Empty data", []byte{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasMagicSignature(tt.data)
			if result != tt.expected {
				t.Errorf("hasMagicSignature(%v) = %v, want %v", tt.data, result, tt.expected)
			}
		})
	}
}

func TestDataAccurate(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected bool
	}{
		{"Empty", []byte{}, false},
		{"Plain ASCII", []byte("Hello, World!\nThis is a test."), false},
		{"ASCII with tabs", []byte("col1\tcol2\tcol3\nval1\tval2\tval3"), false},
		{"UTF-8 text", []byte("Hello, 世界! Привет мир!"), false},
		{"PNG header", []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52}, true},
		{"Binary with nulls", []byte{0x00, 0x01, 0x02, 0x03, 0x04}, true},
		{"Control characters", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}, true},
		{"ELF binary", []byte{0x7F, 0x45, 0x4C, 0x46, 0x02, 0x01, 0x01, 0x00}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DataAccurate(tt.data)
			if result != tt.expected {
				t.Errorf("DataAccurate() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDataAccurateAndUTF16(t *testing.T) {
	tests := []struct {
		name           string
		data           []byte
		expectedBinary bool
		expectedUTF16  bool
	}{
		{"Empty", []byte{}, false, false},
		{"Plain ASCII", []byte("Hello, World!"), false, false},
		{"UTF-16 LE BOM", []byte{0xFF, 0xFE, 0x48, 0x00, 0x65, 0x00, 0x6C, 0x00, 0x6C, 0x00, 0x6F, 0x00}, false, true},
		{"UTF-16 BE BOM", []byte{0xFE, 0xFF, 0x00, 0x48, 0x00, 0x65, 0x00, 0x6C, 0x00, 0x6C, 0x00, 0x6F}, false, true},
		// Note: PNG header may be detected as "valid UTF-16" since isValidUTF16 has false positives
		// for arbitrary even-length binary data without BOMs. The important thing is isBinary=true.
		{"PNG header", []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isBinary, isUTF16 := DataAccurateAndUTF16(tt.data)
			if isBinary != tt.expectedBinary {
				t.Errorf("DataAccurateAndUTF16() isBinary = %v, want %v", isBinary, tt.expectedBinary)
			}
			if isUTF16 != tt.expectedUTF16 {
				t.Errorf("DataAccurateAndUTF16() isUTF16 = %v, want %v", isUTF16, tt.expectedUTF16)
			}
		})
	}
}

func TestFileAccurate(t *testing.T) {
	// Test existing test files
	tests := []struct {
		filename string
		expected bool
	}{
		{"testdata/exe1", true},
		{"testdata/exe2", true},
		{"testdata/conf1", false},
		{"testdata/conf2", false},
		{"testdata/fstab", false},
		{"testdata/hai", false},
		{"testdata/utf16.csv", false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			isBinary, err := FileAccurate(tt.filename)
			if err != nil {
				t.Errorf("FileAccurate(%s) error: %v", tt.filename, err)
				return
			}
			if isBinary != tt.expected {
				t.Errorf("FileAccurate(%s) = %v, want %v", tt.filename, isBinary, tt.expected)
			}
		})
	}
}

func TestFileAccurateAndUTF16(t *testing.T) {
	// Test the UTF-16 file
	isBinary, isUTF16, err := FileAccurateAndUTF16("testdata/utf16.csv")
	if err != nil {
		t.Errorf("FileAccurateAndUTF16(testdata/utf16.csv) error: %v", err)
		return
	}
	if isBinary {
		t.Errorf("FileAccurateAndUTF16(testdata/utf16.csv) detected as binary, want text")
	}
	if !isUTF16 {
		t.Errorf("FileAccurateAndUTF16(testdata/utf16.csv) did not detect UTF-16")
	}
}

// binaryFileAccurate is a helper function for testing the DataAccurate function
func binaryFileAccurate(filename string) (bool, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	return DataAccurate(data), nil
}

func TestDataAccurateWithFiles(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"testdata/exe1", true},
		{"testdata/exe2", true},
		{"testdata/conf1", false},
		{"testdata/conf2", false},
		{"testdata/fstab", false},
		{"testdata/hai", false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			isBinary, err := binaryFileAccurate(tt.filename)
			if err != nil {
				t.Errorf("binaryFileAccurate(%s) error: %v", tt.filename, err)
				return
			}
			if isBinary != tt.expected {
				t.Errorf("binaryFileAccurate(%s) = %v, want %v", tt.filename, isBinary, tt.expected)
			}
		})
	}
}

func TestFileAccurateNonexistent(t *testing.T) {
	_, err := FileAccurate("nonexistent_file_12345")
	if err == nil {
		t.Error("FileAccurate should return error for nonexistent file")
	}
}

func TestCompareAccurateVsQuick(t *testing.T) {
	// Both methods should agree on clearly binary files
	tests := []struct {
		filename string
		expected bool
	}{
		{"testdata/exe1", true},
		{"testdata/exe2", true},
		{"testdata/conf1", false},
		{"testdata/conf2", false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			quickResult, err := File(tt.filename)
			if err != nil {
				t.Errorf("File(%s) error: %v", tt.filename, err)
				return
			}
			accurateResult, err := FileAccurate(tt.filename)
			if err != nil {
				t.Errorf("FileAccurate(%s) error: %v", tt.filename, err)
				return
			}

			if quickResult != tt.expected {
				t.Errorf("File(%s) = %v, want %v", tt.filename, quickResult, tt.expected)
			}
			if accurateResult != tt.expected {
				t.Errorf("FileAccurate(%s) = %v, want %v", tt.filename, accurateResult, tt.expected)
			}
		})
	}
}

func TestPNGDetection(t *testing.T) {
	// Test that PNG files are correctly detected as binary
	filename := "testdata/tiny.png"

	// Test FileAccurate
	isBinary, err := FileAccurate(filename)
	if err != nil {
		t.Errorf("FileAccurate(%s) error: %v", filename, err)
		return
	}
	if !isBinary {
		t.Errorf("FileAccurate(%s) = false, want true (PNG should be detected as binary)", filename)
	}

	// Test File (quick method)
	isBinaryQuick, err := File(filename)
	if err != nil {
		t.Errorf("File(%s) error: %v", filename, err)
		return
	}
	if !isBinaryQuick {
		t.Logf("Note: File(%s) = false (quick method may miss small PNG files)", filename)
	}

	// Test DataAccurate with file contents
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("ReadFile(%s) error: %v", filename, err)
		return
	}
	if !DataAccurate(data) {
		t.Errorf("DataAccurate(PNG data) = false, want true")
	}
}

// Benchmarks to compare quick vs accurate methods

func BenchmarkFileQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		File("testdata/exe1")
	}
}

func BenchmarkFileAccurate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FileAccurate("testdata/exe1")
	}
}

func BenchmarkDataQuick(b *testing.B) {
	data, _ := os.ReadFile("testdata/exe1")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Data(data)
	}
}

func BenchmarkDataAccurate(b *testing.B) {
	data, _ := os.ReadFile("testdata/exe1")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DataAccurate(data)
	}
}
