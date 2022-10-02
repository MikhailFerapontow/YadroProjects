package simpletest

import "strings"

func isBad(str string) bool {
	return strings.Contains(str, "bad")
}

func IsGood(str string) bool {
	return strings.Contains(str, "good")
}
