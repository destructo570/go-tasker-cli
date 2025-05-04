package utils

import (
	"strings"
)

func GetArgumentByIndex(args []string, index int) *string {
	if len(args) > index {
		var trimmedArg = GetTrimmedString(args[0])
		return &trimmedArg
	} else {
		return nil
	}
}

func GetTrimmedString(arg string) string {
	trimmedArg := strings.Trim(arg, " ")
	return trimmedArg
}

func GetLowercaseString(arg string) string {
	lowerCasedStr := strings.ToLower(arg)
	return lowerCasedStr
}

func GetSanitizedString(arg string) string {
	result := GetLowercaseString(GetTrimmedString(arg))
	return result
}
