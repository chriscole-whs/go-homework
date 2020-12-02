package connect4

type Game struct {

}

func NewGame() *Game {

}

// MakeMove should take in a column number, validate that it between 1 and 7,
// that the column indicated is not full (max 6 checkers per column), and then
// put a checker in that column for the current player, and then change the
// current player so that moves are alternated
func (g *Game) MakeMove(column int) error {

}

// IsOver indicates if the game has been won or if it is a draw (no moves remaining)
func (g *Game) IsOver() bool {

}

// Result should a string describing which player won, or if it was a draw
func (g *Game) Result() string {

}
