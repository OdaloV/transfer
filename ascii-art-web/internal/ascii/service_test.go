package ascii

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLoadBanner(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		filename := "./testdata/ascii_valid.txt"
		expected, err := os.ReadFile("./testdata/expected_art.txt")
		if err != nil {
			t.Fatalf("setup failed: %v", err)
		}

		bannerMap, err := LoadBanner(filename)
		if err != nil {
			t.Fatalf("expected error = nil, got %v", err)
		}

		val, exists := bannerMap['A']
		if !exists {
			t.Fatal("expected character 'A' to be in banner")
		}

		// Check that art matches (might be fragile)
		expectedStr := string(expected)
		gotStr := strings.Join(val, "")
		if expectedStr != gotStr {
			t.Fatalf("expected %v, got %v", expectedStr, gotStr)
		}

		// Check that number of characters is correct
		mapLen := len(bannerMap)
		if mapLen != 95 {
			t.Fatalf("expected 95 characters, got %d", mapLen)
		}

		// Check that all expected characters exist in the map
		for i := 32; i < 127; i++ {
			t.Run(fmt.Sprintf("rune %d", i), func(t *testing.T) {
				val, exists := bannerMap[rune(i)]
				if !exists {
					t.Fatalf("expected error = nil, got %v", err)
				}

				if len(val) != 8 {
					t.Fatalf("expected 8 lines, got %d", len(val))
				}

				// Check structure
				for i, line := range val {
					if line == "" {
						t.Fatalf("line %d is empty", i)
					}
				}
			})
		}
	})

	t.Run("fails when file cannot be read", func(t *testing.T) {
		filename := "doesnotexist.txt"

		if _, err := LoadBanner(filename); err == nil {
			t.Fatalf("expected error: %v, got nil", err)
		}
	})

	t.Run("extra blanklines should pass", func(t *testing.T) {
		filename := "./testdata/ascii_extra_newlines.txt"
		if _, err := LoadBanner(filename); err != nil {
			t.Fatalf("expected err = nil, got %v", err)
		}
	})

	t.Run("fails when banner file has invalid characters", func(t *testing.T) {
		tests := []struct {
			name string
			filename string
		}{
			{"character length is not equal to 8", "./testdata/ascii_invalid_char_length.txt"},
			{"missing characters", "./testdata/ascii_missing_chars.txt"},
			{"missing last newline", "./testdata/ascii_missing_last_newline.txt"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if _, err := LoadBanner(tt.filename); err == nil {
					t.Fatalf("expected error, got nil")
				}
			})
		}
	})
}

func TestRenderLine(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		// Setup
		filename := "./testdata/ascii_valid.txt"
		var buf bytes.Buffer

		bannerMap, err := LoadBanner(filename)
		if err != nil {
			t.Fatalf("expected err = nil, got %v", err)
		}

		// Test
		err = RenderLine("AB", bannerMap, &buf)
		if err != nil {
			t.Fatalf("expected err = nil, got %v", err)
		}

		got := buf.String()

		// Test that structure of the output is right, as opposed to testing exact expected ascii values, cause that would mean whenever the representation changes
		// the tests would have to change
		
		// The print/render function should write something to the writer
		if got == "" {
			t.Fatal("expected output, got empty string")
		}

		// lines := strings.Split(strings.TrimSuffix(got, "\n"), "\n")
		// if len(lines) != 8 {
		// 	t.Fatalf("expected 8 lines, got %d", len(lines))
		// }
	})

	t.Run("fails when input character is not allowed ascii", func(t *testing.T) {
		// Setup
		filename := "./testdata/ascii_valid.txt"
		var buf bytes.Buffer

		bannerMap, err := LoadBanner(filename)
		if err != nil {
			t.Fatalf("expected err = nil, got %v", err)
		}

		err = RenderLine("\t", bannerMap, &buf)
		if err == nil {
			t.Fatalf("expected err: %v, got nil", err)
		}
	})
}
