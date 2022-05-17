package checkers

import (
	"fmt"
	"runtime"
)

func StringElementExists(element string, array []string) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

func DetectOS() string {
	os := runtime.GOOS
	var osName string

	switch os {
	case "windows":
		osName = "windows"
	case "darwin":
		osName = "mac"
	case "linux":
		osName = "linux"
	default:
		osName = fmt.Sprintf("%s", os)
	}

	return osName
}

// Converts provided integer to a string
func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}
