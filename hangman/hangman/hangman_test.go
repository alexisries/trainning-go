package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"
	hasLetter := lettersInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "a"
	hasLetter := lettersInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contains letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Error should be returned an invalid word=''")
	}
}

func TestGoodGuest(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)

}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("b")
	validState(t, "alreadyGuess", g.State)
}

func TestGameBadGuest(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("i")
	validState(t, "badGuess", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := New(3, "bob")
	bulkMakeAGuess(t, g, "bob")
	validState(t, "won", g.State)

}

func TestGameLost(t *testing.T) {
	g, _ := New(3, "bob")
	bulkMakeAGuess(t, g, "tuiopgkh")
	validState(t, "lost", g.State)

}

func bulkMakeAGuess(t *testing.T, g *Game, letters string) {
	for _, letter := range letters {
		g.MakeAGuess(string(letter))
	}
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v'. got=%v", expectedState, actualState)
		return false
	}
	return true
}
