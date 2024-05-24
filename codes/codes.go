package codes

import (
	"slices"
	"strings"
)

var alphabetToMorse = map[string]string{
	// alpha
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",

	// numeric
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",

	// punctuation
	`.`: ".-.-.-",
	`,`: "--..--",
	`?`: "..--..",
	`'`: ".----.",
	`!`: "-.-.--",
	`/`: "-..-.",
	`(`: "-.--.",
	`)`: "-.--.-",
	`&`: ".-...",
	`:`: "---...",
	`;`: "-.-.-.",
	`=`: "-...-",
	`+`: ".-.-.",
	`-`: "-....-",
	`_`: "..--.-",
	`"`: ".-..-.",
	`$`: "...-..-",
	`@`: ".--.-.",

	// non-latin
	`Ä`: ".-.-",
	`Ą`: ".-.-",
	`Æ`: ".-.-",
	`À`: ".--.-",
	`Å`: ".--.-",
	`Ç`: "-.-..",
	`Š`: "----",
	`Ð`: "..--.",
	`Ł`: ".-..-",
	`Ę`: "..-..",
	`Ĝ`: "--.-.",
	`Ĵ`: ".---.",
	`Ñ`: "--.--",
	`Ø`: "---.",
	`Ŝ`: "...-.",
	`Þ`: ".--..",
	`Ŭ`: "..--",
}

var morseToAlphabet map[string]string

func init() {
	morseToAlphabet = make(map[string]string)

	grouped := make(map[string][]string)
	for letter, morse := range alphabetToMorse {
		grouped[morse] = append(grouped[morse], letter)
	}

	for morse, letters := range grouped {
		if len(letters) > 1 {
			slices.SortStableFunc(letters, func(a, b string) int {
				return strings.Compare(a, b)
			})
		}
		morseToAlphabet[morse] = letters[0]
	}
}

// AlphabetToMorse converts alphabet to morse code
func AlphabetToMorse(alphabet string) string {
	return alphabetToMorse[alphabet]
}

// MorseToAlphabet converts morse code to alphabet
func MorseToAlphabet(morse string) string {
	return morseToAlphabet[morse]
}
