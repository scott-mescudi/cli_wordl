# Wordl Game

A terminal-based word-guessing game written in Go. The goal of the game is to guess a randomly selected word, and with each guess, get feedback on which letters are correct, incorrect, or in the wrong position.

## Features

- Customizable themes and letter counts.
- Configurable settings stored in a `config.toml` file.
- Interactive terminal interface for guessing words.
- Supports various word lengths (2, 3, 5, 7, and 10 letters).
- Visual feedback with colored output for correct, incorrect, and misplaced letters.
  
## Requirements

- Go 1.16 or higher
- Terminal or command-line interface

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/scott-mescudi/cli_wordl.git
    cd cli_wordl
    ```

2. **Install dependencies**:

    This project requires `github.com/BurntSushi/toml` for TOML file parsing, which is included as part of the `go.mod` file.

    ```bash
    go mod tidy
    ```

3. **Set up word lists**:

    The game requires word lists for different letter lengths. Make sure the following word list files are available in your project directory under `wordlists/`:
    
    - `wordlists/2_letter_words.txt`
    - `wordlists/3_letter_words.txt`
    - `wordlists/5_letter_words.txt`
    - `wordlists/7_letter_words.txt`
    - `wordlists/10_letter_words.txt`

4. **Build and run the game**:

    ```bash
    go build -o wordl main.go
    ./wordl
    ```

## Configuration

The game uses a `config.toml` file to store user preferences for themes and word lengths. If this file does not exist, it will be created with the default settings.

### Example `config.toml`:

```toml
theme = "wordl"
letter_count = 5
```

You can modify the settings by running the game and choosing the `:settings` option or directly editing the `config.toml` file.

## Game Controls

- Type a word of the correct length to guess.
- Feedback is provided as:
  - `▧ = Correct letter in the correct position`
  - `▧ = Correct letter, but in the wrong position`
  - `▧ = Incorrect letter`
- Type `:q` to quit the game.
- Type `:settings` to modify the theme and word length.

### Themes Available:

- `wordl` (default)
- `oceanbreeze`
- `desertsunset`
- `neonnight`
- `forestmeadow`
- `cyberpunkglow`
- `monochromepop`
- `impossible`

### Word Length Options:

- 2 letters
- 3 letters
- 5 letters (default)
- 7 letters
- 10 letters

## How It Works

1. The game will load a random word from the appropriate word list (based on the configured word length).
2. The user will have several attempts to guess the word.
3. After each guess, feedback is provided:
   - Correct letters are highlighted in the right color.
   - Incorrect letters are marked in another color.
   - Misplaced letters are marked in yet another color.
4. Once the correct word is guessed, the game congratulates the player and ends.

## File Structure

```
wordl/
├── wordlists/
│   ├── 2_letter_words.txt
│   ├── 3_letter_words.txt
│   ├── 5_letter_words.txt
│   ├── 7_letter_words.txt
│   └── 10_letter_words.txt
├── config.toml
└── main.go
```

## Customizing Themes

Themes are defined in the `wordl/internal/themes` package and can be customized by modifying the theme colors and symbols. The available theme attributes include:

- `Colors.AcceptColor`, `Colors.InWordColor`, `Colors.WrongColor`, `Colors.BoardColor`
- `Symbols.Horizontal`, `Symbols.Vertical`, `Symbols.TopLeft`, `Symbols.TopRight`, `Symbols.BottomLeft`, `Symbols.BottomRight`

You can add more themes or modify the existing ones by changing the theme files.
