package themes


const Reset = "\033[0m"

type ColorTheme struct {
	WrongColor   string
	AcceptColor string
	InWordColor  string
	BoardColor    string
}

type BorderSymbols struct {
	TopRight     string
	TopLeft      string
	BottomRight  string
	BottomLeft   string
	Horizontal   string
	Vertical     string
}


type Theme struct {
	Colors  ColorTheme
	Symbols BorderSymbols
}

func SetTheme(themeName string) *Theme {
	var theme *Theme
	switch themeName {
	case "oceanbreeze":
		theme = &Theme{
			Colors: ColorTheme{
				WrongColor:   "\033[1;34m",  // Deep Blue
				AcceptColor: "\033[1;36m",  // Cyan
				InWordColor:  "\033[1;94m",  // Light Blue
				BoardColor:    "\033[1;37m",  // White
			},
			Symbols: BorderSymbols{
				TopRight:    "╮",
				TopLeft:     "╭",
				BottomRight: "╯",
				BottomLeft:  "╰",
				Horizontal:  "─",
				Vertical:    "│",
			},
		}
	case "desertsunset":
		theme = &Theme{
			Colors: ColorTheme{
				WrongColor:   "\033[1;31m",         // Red
				AcceptColor: "\033[1;38;5;214m",   // Orange
				InWordColor:  "\033[1;33m",         // Yellow
				BoardColor:    "\033[1;38;5;94m",    // Brown
			},
			Symbols: BorderSymbols{
				TopRight:    "▞",
				TopLeft:     "▚",
				BottomRight: "▚",
				BottomLeft:  "▞",
				Horizontal:  "━",
				Vertical:    "┃",
			},
		}
	case "neonnight":
		theme = &Theme{
			Colors: ColorTheme{
				WrongColor:   "\033[1;35m",  // Purple
				AcceptColor: "\033[1;92m",  // Lime
				InWordColor:  "\033[1;96m",  // Cyan
				BoardColor:    "\033[1;95m",  // Magenta
			},
			Symbols: BorderSymbols{
				TopRight:    "╖",
				TopLeft:     "╓",
				BottomRight: "╜",
				BottomLeft:  "╙",
				Horizontal:  "═",
				Vertical:    "║",
			},
		}
	case "forestmeadow":
		theme = &Theme{
			Colors: ColorTheme{
				WrongColor:   "\033[1;32m",  // Dark Green
				AcceptColor: "\033[1;92m",  // Light Green
				InWordColor:  "\033[1;33m",  // Yellow
				BoardColor:    "\033[1;93m",  // Light Yellow
			},
			Symbols: BorderSymbols{
				TopRight:    "╮",
				TopLeft:     "╭",
				BottomRight: "╯",
				BottomLeft:  "╰",
				Horizontal:  "─",
				Vertical:    "│",
			},
		}
	case "cyberpunkglow":
		theme = &Theme{
			Colors: ColorTheme{
				WrongColor:   "\033[1;35m",  // Magenta
				AcceptColor: "\033[1;96m",  // Bright Cyan
				InWordColor:  "\033[1;34m",  // Blue
				BoardColor:    "\033[1;33m",  // Yellow
			},
			Symbols: BorderSymbols{
				TopRight:    "╗",
				TopLeft:     "╔",
				BottomRight: "╝",
				BottomLeft:  "╚",
				Horizontal:  "═",
				Vertical:    "║",
			},
		}
	case "monochromepop":
		theme = &Theme{
			Colors: ColorTheme{
				WrongColor:   "\033[1;37m",  // White
				AcceptColor: "\033[1;90m",  // Gray
				InWordColor:  "\033[1;30m",  // Black
				BoardColor:    "\033[1;97m",  // Light Gray
			},
			Symbols: BorderSymbols{
				TopRight:    "»",
				TopLeft:     "«",
				BottomRight: "«",
				BottomLeft:  "»",
				Horizontal:  "─",
				Vertical:    "│",
			},
		}
	case "impossible":
			theme = &Theme{
				Colors: ColorTheme{
					WrongColor:   "\033[1;37m",  // Red
					AcceptColor: "\033[1;37m",  // Green
					InWordColor:  "\033[1;37m",  // Magenta
					BoardColor:    "\033[1;37m", // Baby blue
				},
				Symbols: BorderSymbols{
					TopRight:    "┐",
					TopLeft:     "┌",
					BottomRight: "┘",
					BottomLeft:  "└",
					Horizontal:  "─",
					Vertical:    "│",
				},
			}

	default:
		// If the theme name does not match, return the default theme
		theme = &Theme{
			Colors: ColorTheme{
				WrongColor:   "\033[1;31m",  // Red
				AcceptColor: "\033[1;32m",  // Green
				InWordColor:  "\033[1;35m",  // Magenta
				BoardColor:    "\033[38;5;153m", // Baby blue
			},
			Symbols: BorderSymbols{
				TopRight:    "┐",
				TopLeft:     "┌",
				BottomRight: "┘",
				BottomLeft:  "└",
				Horizontal:  "─",
				Vertical:    "│",
			},
		}
	}

	return theme
}