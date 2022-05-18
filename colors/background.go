package colors

import "github.com/iVitaliya/colors-go/checkers"

type IBgColors struct {
	BgBlack   func(string) string
	BgRed     func(string) string
	BgGreen   func(string) string
	BgYellow  func(string) string
	BgBlue    func(string) string
	BgMagenta func(string) string
	BgCyan    func(string) string
	BgWhite   func(string) string
}

func BgColors() *IBgColors {
	Colors := &IBgColors{
		BgBlack:   checkers.Initiate(40, 49, ""),
		BgRed:     checkers.Initiate(41, 49, ""),
		BgGreen:   checkers.Initiate(42, 49, ""),
		BgYellow:  checkers.Initiate(43, 49, ""),
		BgBlue:    checkers.Initiate(44, 49, ""),
		BgMagenta: checkers.Initiate(45, 49, ""),
		BgCyan:    checkers.Initiate(46, 49, ""),
		BgWhite:   checkers.Initiate(47, 49, ""),
	}

	return Colors
}
