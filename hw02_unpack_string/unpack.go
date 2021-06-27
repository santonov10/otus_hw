package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	escapeRune       = '\\'
	ErrInvalidString = errors.New("invalid string")
)

func Unpack(str string) (string, error) {
	if !isValidString(str) {
		return "", ErrInvalidString
	}

	var strBuilder strings.Builder

	strSlice := []rune(str)
	lastRuneIndex := len(strSlice) - 1

	if len(strSlice) <= 1 {
		return str, nil
	}

	skipNextIteration := false
	for i, currentRune := range strSlice {
		if skipNextIteration {
			skipNextIteration = false
			continue
		}

		var prevRune rune
		var nextRune rune

		if i > 0 {
			prevRune = strSlice[i-1]
		}

		if i < lastRuneIndex {
			nextRune = strSlice[i+1]
		}

		if i < lastRuneIndex {
			if isEscapeRune(currentRune) && !isEscapeRune(prevRune) {
				continue
			}

			writeRuneTimes := 1
			if unicode.IsDigit(nextRune) {
				writeRuneTimes, _ = strconv.Atoi(string(nextRune))
				skipNextIteration = true
			}
			strBuilder.WriteString(strings.Repeat(string(currentRune), writeRuneTimes))
			if isEscapeRune(currentRune) {
				strSlice[i] = 0
			}
		} else {
			strBuilder.WriteRune(currentRune)
		}
	}

	return strBuilder.String(), nil
}

func isEscapeRune(r rune) bool {
	return r == escapeRune
}

func isValidString(str string) bool {
	strSlice := []rune(str)
	if len(strSlice) > 0 {
		firstRune := strSlice[0]
		if unicode.IsDigit(firstRune) {
			return false
		}
		match, _ := regexp.MatchString(`[^\\][\d]{2,}`, str)
		if match {
			return false
		}

		strBuf := strings.ReplaceAll(str, `\\`, `\ `)
		match, _ = regexp.MatchString(`\\n`, strBuf)
		if match {
			return false
		}
	}
	return true
}
