package checkers

import (
	"strings"
)

func ReplaceClose(index int, str string, close string, replace string, head string, tail string, next int) string {
	if next < 0 {
		return head + tail
	}

	return head + ReplaceClose(next, tail, close, replace, head+replace, tail, strings.Index(tail, close))
}

func ClearBleed(index int, str string, open string, close string, replace string) string {
	if index < 0 {
		return open + str + close
	}

	return open + ReplaceClose(index, str, close, replace, "", "", strings.Index(str, open))
}

func FilterEmpty(open string, close string, replace string, at int) func(string) string {
	return func(str string) string {
		if str == "" || len(str) < 1 {
			return ""
		}

		return ClearBleed(at, str, open, close, replace)
	}
}

func Initiate(open string, close string, replace string) func(string) string {
	return func(s string) string {
		return FilterEmpty(open, close, replace, strings.Index(s, open))(s)
	}
}
