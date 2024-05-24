package morse

import (
	"github.com/AyakuraYuki/morse/codes"
	"strings"
)

type Morse interface {
	Encoder
	Decoder
}

type Encoder interface {
	Encode(string) (string, error)
}

type Decoder interface {
	Decode(string) (string, error)
}

type morse struct{}

func New() Morse {
	return &morse{}
}

const (
	Space = ` `
	Slash = `/`
)

func (m morse) Encode(raw string) (string, error) {
	if raw == "" {
		return "", nil
	}

	data := strings.ToUpper(raw)
	parts := strings.Split(data, Space)
	wordAmount := len(parts)
	result := strings.Builder{}
	for cursor, part := range parts {
		length := len([]rune(part))
		for i, c := range part {
			s := string(c)
			result.WriteString(codes.AlphabetToMorse(s))
			if (i + 1) != length {
				result.WriteString(Space)
			}
		}
		if wordAmount != (cursor + 1) {
			result.WriteString(Space)
			result.WriteString(Slash)
			result.WriteString(Space)
		}
	}

	return result.String(), nil
}

func (m morse) Decode(encodedCodes string) (string, error) {
	if encodedCodes == "" {
		return "", nil
	}

	parts := strings.Split(encodedCodes, Space)
	result := strings.Builder{}
	for _, part := range parts {
		if part == Slash {
			result.WriteString(Space)
			continue
		}
		result.WriteString(codes.MorseToAlphabet(part))
	}
	return result.String(), nil
}
