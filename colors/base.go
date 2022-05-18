package colors

import "github.com/iVitaliya/colors-go/checkers"

type IBaseColors struct {
	Black   func(string) string
	Red     func(string) string
	Green   func(string) string
	Yellow  func(string) string
	Blue    func(string) string
	Magenta func(string) string
	Cyan    func(string) string
	White   func(string) string
	Gray    func(string) string
}

func BaseColors() *IBaseColors {
	Colors := &IBaseColors{
		Black:   checkers.Initiate(30, 39, ""),
		Red:     checkers.Initiate(31, 39, ""),
		Green:   checkers.Initiate(32, 39, ""),
		Yellow:  checkers.Initiate(33, 39, ""),
		Blue:    checkers.Initiate(34, 39, ""),
		Magenta: checkers.Initiate(35, 39, ""),
		Cyan:    checkers.Initiate(36, 39, ""),
		White:   checkers.Initiate(37, 39, ""),
		Gray:    checkers.Initiate(90, 39, ""),
	}

	return Colors
}
