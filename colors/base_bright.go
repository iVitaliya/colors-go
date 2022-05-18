package colors

import "github.com/iVitaliya/colors-go/checkers"

type IBrightColors struct {
	BrightBlack   func(string) string
	BrightRed     func(string) string
	BrightGreen   func(string) string
	BrightYellow  func(string) string
	BrightBlue    func(string) string
	BrightMagenta func(string) string
	BrightCyan    func(string) string
	BrightWhite   func(string) string
}

func BrightColors() *IBrightColors {
	Colors := &IBrightColors{
		BrightBlack:   checkers.Initiate(90, 39, ""),
		BrightRed:     checkers.Initiate(91, 39, ""),
		BrightGreen:   checkers.Initiate(92, 39, ""),
		BrightYellow:  checkers.Initiate(93, 39, ""),
		BrightBlue:    checkers.Initiate(94, 39, ""),
		BrightMagenta: checkers.Initiate(95, 39, ""),
		BrightCyan:    checkers.Initiate(96, 39, ""),
		BrightWhite:   checkers.Initiate(97, 39, ""),
	}

	return Colors
}
