package colors

import "github.com/iVitaliya/colors-go/checkers"

type IFormatColors struct {
	Reset         func(string) string
	Bold          func(string) string
	Dim           func(string) string
	Italic        func(string) string
	Underline     func(string) string
	Inverse       func(string) string
	Hidden        func(string) string
	StrikeThrough func(string) string
}

func FormatColors() *IFormatColors {
	Colors := &IFormatColors{
		Reset:         checkers.Initiate(0, 0, ""),
		Bold:          checkers.Initiate(1, 22, "\x1b[22m\x1b[1m"),
		Dim:           checkers.Initiate(2, 22, "\x1b[22m\x1b[2m"),
		Italic:        checkers.Initiate(3, 23, ""),
		Underline:     checkers.Initiate(4, 24, ""),
		Inverse:       checkers.Initiate(7, 27, ""),
		Hidden:        checkers.Initiate(8, 28, ""),
		StrikeThrough: checkers.Initiate(9, 29, ""),
	}

	return Colors
}
