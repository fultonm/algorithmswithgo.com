package module01

import (
	"fmt"
	"log"
	"unicode/utf8"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	return ReverseR(word, "")
}

func ReverseR(word string, rev string) string {
	if len(word) == 0 {
		return rev
	}

	return ReverseR(GetAllButFistRune(word), GetNthRune(word, 0)+rev)
}

func GetAllButFistRune(word string) string {
	_, i := utf8.DecodeRuneInString(word)
	return word[i:]
}

func GetNthRune(word string, n int) string {
	if n > len(word) {
		log.Fatal("The word of length " + fmt.Sprint(len(word)) + " is not long enough to get the " + fmt.Sprint(n) + "th rune.")
	}
	for i, iRune := range word {
		if i < n {
			continue
		} else if i == n {
			return string(iRune)
		}
	}
	return ""
}
