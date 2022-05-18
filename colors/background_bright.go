package colors

import "github.com/iVitaliya/colors-go/checkers"

type IBrightBackgroundColors struct {
	BgBlackBright   func(string) string
	BgRedBright     func(string) string
	BgGreenBright   func(string) string
	BgYellowBright  func(string) string
	BgBlueBright    func(string) string
	BgMagentaBright func(string) string
	BgCyanBright    func(string) string
	BgWhiteBright   func(string) string
}

func BrightBackgroundColors() *IBrightBackgroundColors {
	Colors := &IBrightBackgroundColors{
		BgBlackBright:   checkers.Initiate(100, 49, ""),
		BgRedBright:     checkers.Initiate(101, 49, ""),
		BgGreenBright:   checkers.Initiate(102, 49, ""),
		BgYellowBright:  checkers.Initiate(103, 49, ""),
		BgBlueBright:    checkers.Initiate(104, 49, ""),
		BgMagentaBright: checkers.Initiate(105, 49, ""),
		BgCyanBright:    checkers.Initiate(106, 49, ""),
		BgWhiteBright:   checkers.Initiate(107, 49, ""),
	}

	return Colors
}
