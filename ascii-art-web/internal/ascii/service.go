package ascii

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var ErrBannerNotFound = errors.New("banner not found")
var ErrBannerUnreadable = errors.New("banner unreadable")
var ErrInvalidInput = errors.New("invalid input")
var ErrInvalidBannerFormat = errors.New("invalid banner format")

func LoadBanner(path string) (map[rune][]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			return nil, ErrBannerNotFound
		case os.IsPermission(err):
			return nil, ErrBannerUnreadable
			default:
			return nil, ErrBannerUnreadable
		}
	}

	lines := strings.Split(string(content), "\n")

	for len(lines) > 0 && lines[0] == "" {
	lines = lines[1:]
}

	banner := make(map[rune][]string)

	char := rune(32)

	for i := 0; i+7 < len(lines) && char <= 126; i += 9 {
		banner[char] = lines[i : i+8]
		char++
	}

	if char != 127 {
		return nil, ErrInvalidBannerFormat
	}

	return banner, nil
}

func RenderLine(line string, banner map[rune][]string, w io.Writer) error {

	lines := strings.Split(line, "\n")

	for _, l := range lines {
		if l == "" {
			fmt.Fprintln(w)
			continue
		}
		
		output := make([]string, 8)
		for _, c := range l {
			art, ok := banner[c]
			if !ok {
				return ErrInvalidInput
			}
			for row := 0; row < 8; row++ {
				output[row] += art[row]
			}
		}
		for _, row := range output {
			fmt.Fprintln(w, row)
		}
	
	}

	return nil
}

func Generate(text, banner string) (string, error) {
	path := "banners/" + banner + ".txt"


	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\\n", "\n")

	if text == "" || text == "\n" {
		return text, nil
	}

	bannerMap, err := LoadBanner(path)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = RenderLine(text, bannerMap, &buf)
	if err != nil {
		return "", err
	}	
	
	return buf.String(), nil
}