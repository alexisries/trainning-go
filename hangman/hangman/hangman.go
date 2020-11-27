package hangman

import (
	"fmt"
	"strings"
)

type Game struct {
	State        string
	Letters      []string
	FoundLetters []string
	UseLetters   []string
	TurnsLeft    int
	MaxRetry     int
	retryCounter int
}

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("Word '%s' must be at least 3 characters. got=%v", word, len(word))
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UseLetters:   []string{},
		TurnsLeft:    turns,
	}
	return g, nil
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)

	if lettersInWord(guess, g.UseLetters) {
		g.State = "alreadyGuess"
	} else if lettersInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
	} else {
		g.State = "badGuess"
		g.LoseTurn(guess)
		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
	if hasWon(g.Letters, g.FoundLetters) {
		g.State = "won"
	}
}

func (g *Game) RevealLetter(guess string) {
	g.UseLetters = append(g.UseLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func (g *Game) LoseTurn(guess string) {
	g.TurnsLeft--
	g.UseLetters = append(g.UseLetters, guess)
}

func lettersInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
