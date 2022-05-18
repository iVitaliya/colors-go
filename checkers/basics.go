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

	return open + ReplaceClose(index, str, close, replace, "", "", strings.Index(str, open)) + close
}

func FilterEmpty(open string, close string, replace string, at int) func(string) string {
	return func(str string) string {
		if str == "" || len(str) < 1 {
			return ""
		}

		return ClearBleed(at, str, open, close, replace)
	}
}

func Initiate(open int, close int, replace string) func(string) string {
	open_val := IntToString(open)
	close_val := IntToString(close)

	return func(s string) string {
		return FilterEmpty("\x1b["+open_val+"m", "\x1b["+close_val+"m", replace, strings.Index(s, open_val))(s)
	}
}
