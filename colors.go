package colors

import (
	"os"

	"github.com/iVitaliya/colors-go/checkers"
	CLRS "github.com/iVitaliya/colors-go/colors"
	"github.com/mattn/go-isatty"
)

var isDisabled = os.Getenv("NO_COLOR") != "" || checkers.StringElementExists("--no-color", os.Args)
var isForced = os.Getenv("FORCED") != "" || checkers.StringElementExists("--color", os.Args)
var isWindows = checkers.DetectOS() == "windows"
var isCompatibleTerminal = isatty.IsTerminal(os.Stdout.Fd()) && os.Getenv("TERM") != "dumb"
var isCI = os.Getenv("CI") != "" && (os.Getenv("GITHUB_ACTIONS") != "" || os.Getenv("GITLAB_CI") != "" || os.Getenv("CIRCLECI") != "")
var isColorSupported = !isDisabled && (isForced || isWindows || isCompatibleTerminal || isCI)

func Reset(str string) string {
	if isColorSupported {
		return CLRS.FormatColors().Reset(str)
	} else {
		panic("Color not supported :(")
	}
}

func Bold(str string) string {
	if isColorSupported {
		return CLRS.FormatColors().Bold(str)
	} else {
		panic("Color not supported :(")
	}
}

func Dim(str string) string {
	if isColorSupported {
		return CLRS.FormatColors().Dim(str)
	} else {
		panic("Color not supported :(")
	}
}

func Italic(str string) string {
	if isColorSupported {
		return CLRS.FormatColors().Italic(str)
	} else {
		panic("Color not supported :(")
	}
}

func Underline(str string) string {
	if isColorSupported {
		return CLRS.FormatColors().Underline(str)
	} else {
		panic("Color not supported :(")
	}
}

func Inverse(str string) string {
	if isColorSupported {
		return CLRS.FormatColors().Inverse(str)
	} else {
		panic("Color not supported :(")
	}
}

func Hidden(str string) string {
	if isColorSupported {
		return CLRS.FormatColors().Hidden(str)
	} else {
		panic("Color not supported :(")
	}
}

func StrikeThrough(str string) string {
	if isColorSupported {
		return CLRS.FormatColors().StrikeThrough(str)
	} else {
		panic("Color not supported :(")
	}
}

func Black(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().Black(str)
	} else {
		panic("Color not supported :(")
	}
}

func Red(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().Red(str)
	} else {
		panic("Color not supported :(")
	}
}

func Green(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().Green(str)
	} else {
		panic("Color not supported :(")
	}
}

func Yellow(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().Yellow(str)
	} else {
		panic("Color not supported :(")
	}
}

func Blue(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().Blue(str)
	} else {
		panic("Color not supported :(")
	}
}

func Magenta(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().Magenta(str)
	} else {
		panic("Color not supported :(")
	}
}

func Cyan(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().Cyan(str)
	} else {
		panic("Color not supported :(")
	}
}

func White(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().White(str)
	} else {
		panic("Color not supported :(")
	}
}

func Gray(str string) string {
	if isColorSupported {
		return CLRS.BaseColors().Gray(str)
	} else {
		panic("Color not supported :(")
	}
}

func BrightBlack(str string) string {
	if isColorSupported {
		return CLRS.BrightColors().BrightBlack(str)
	} else {
		panic("Color not supported :(")
	}
}

func BrightRed(str string) string {
	if isColorSupported {
		return CLRS.BrightColors().BrightRed(str)
	} else {
		panic("Color not supported :(")
	}
}

func BrightGreen(str string) string {
	if isColorSupported {
		return CLRS.BrightColors().BrightGreen(str)
	} else {
		panic("Color not supported :(")
	}
}

func BrightYellow(str string) string {
	if isColorSupported {
		return CLRS.BrightColors().BrightYellow(str)
	} else {
		panic("Color not supported :(")
	}
}

func BrightBlue(str string) string {
	if isColorSupported {
		return CLRS.BrightColors().BrightBlue(str)
	} else {
		panic("Color not supported :(")
	}
}

func BrightMagenta(str string) string {
	if isColorSupported {
		return CLRS.BrightColors().BrightMagenta(str)
	} else {
		panic("Color not supported :(")
	}
}

func BrightCyan(str string) string {
	if isColorSupported {
		return CLRS.BrightColors().BrightCyan(str)
	} else {
		panic("Color not supported :(")
	}
}

func BrightWhite(str string) string {
	if isColorSupported {
		return CLRS.BrightColors().BrightWhite(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBlack(str string) string {
	if isColorSupported {
		return CLRS.BackgroundColors().BgBlack(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgRed(str string) string {
	if isColorSupported {
		return CLRS.BackgroundColors().BgRed(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgGreen(str string) string {
	if isColorSupported {
		return CLRS.BackgroundColors().BgGreen(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgYellow(str string) string {
	if isColorSupported {
		return CLRS.BackgroundColors().BgYellow(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBlue(str string) string {
	if isColorSupported {
		return CLRS.BackgroundColors().BgBlue(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgMagenta(str string) string {
	if isColorSupported {
		return CLRS.BackgroundColors().BgMagenta(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgCyan(str string) string {
	if isColorSupported {
		return CLRS.BackgroundColors().BgCyan(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgWhite(str string) string {
	if isColorSupported {
		return CLRS.BackgroundColors().BgWhite(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBrightBlack(str string) string {
	if isColorSupported {
		return CLRS.BrightBackgroundColors().BgBlackBright(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBrightRed(str string) string {
	if isColorSupported {
		return CLRS.BrightBackgroundColors().BgRedBright(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBrightGreen(str string) string {
	if isColorSupported {
		return CLRS.BrightBackgroundColors().BgGreenBright(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBrightYellow(str string) string {
	if isColorSupported {
		return CLRS.BrightBackgroundColors().BgYellowBright(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBrightBlue(str string) string {
	if isColorSupported {
		return CLRS.BrightBackgroundColors().BgBlueBright(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBrightMagenta(str string) string {
	if isColorSupported {
		return CLRS.BrightBackgroundColors().BgMagentaBright(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBrightCyan(str string) string {
	if isColorSupported {
		return CLRS.BrightBackgroundColors().BgCyanBright(str)
	} else {
		panic("Color not supported :(")
	}
}

func BgBrightWhite(str string) string {
	if isColorSupported {
		return CLRS.BrightBackgroundColors().BgWhiteBright(str)
	} else {
		panic("Color not supported :(")
	}
}
