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
				WrongColor:   "\033[1;91m",
				AcceptColor:  "\033[1;96m",
				InWordColor:  "\033[1;94m",
				BoardColor:   "\033[1;97m",
			},
			Symbols: BorderSymbols{
				TopRight:    "░",
				TopLeft:     "░",
				BottomRight: "░",
				BottomLeft:  "░",
				Horizontal:  "░",
				Vertical:    "░",
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
		case "diddy":
	theme = &Theme{
		Colors: ColorTheme{
			WrongColor:   "\033[1;95m",  // Vibrant Pink (playful and tropical)
			AcceptColor:  "\033[1;92m",  // Bright Green (palm leaves or tropical plants)
			InWordColor:  "\033[1;93m",  // Bright Yellow (sunshine, bananas)
			BoardColor:   "\033[1;94m",  // Light Blue (ocean sky color)
		},
		Symbols: BorderSymbols{
			TopRight:    "✿",
			TopLeft:     "✿",
			BottomRight: "✿",
			BottomLeft:  "✿",
			Horizontal:  "»",
			Vertical:    "♥",
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
				WrongColor:   "\033[1;31m",  
				AcceptColor: "\033[1;92m",  
				InWordColor:  "\033[1;33m",  
				BoardColor:   "\033[0;32m", 
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
				WrongColor:   "\033[1;95m",  
				AcceptColor: "\033[1;92m",  
				InWordColor:  "\033[1;36m",  // Blue
				BoardColor:    "\033[1;93m",  // Yellow
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
				AcceptColor: "\033[1;93m",  // Gray
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