package stringutil

import "strings"

func IsNullOrEmpty(str string) bool {
	return len(str) == 0 || strings.Trim(str, " ") == ""
}
