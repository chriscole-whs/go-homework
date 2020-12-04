package connect4

import "errors"

const (
	RedPlayer   = "red"
	BlackPlayer = "black"
	DrawResult  = "draw"
)

var (
	ErrIllegalColumn = errors.New("illegal column")
	ErrFullColumn    = errors.New("full column")
	ErrGameOver      = errors.New("game over")
)

// Game models the data needed to play a game of connect4
type Game struct {
}

// NewGame initializes all data associated with a connect4 game
// This will returns a game with a clean game board (and any other metadata) as well as establish
// the current (first player to move) player as the red player -- necessary for the tests to run properly
func NewGame() *Game {
	return &Game{}
}

// MakeMove should take in a column number, validate that it's between 1 and 7,
// that the column indicated is not full (max 6 checkers per column), that the game isn't over,
// and then put a checker in that column for the current player, and then change the
// current player so that moves are alternated
func (g *Game) MakeMove(column int) error {
	return nil
}

// IsOver indicates if the game has been won or if it is a draw (no moves remaining)
func (g *Game) IsOver() bool {
	return false
}

// Result returns a single word, either "draw", "red", "black" or the empty string
// if there is not yet a result
func (g *Game) Result() string {
	return ""
}
