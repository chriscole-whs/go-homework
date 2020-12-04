package main

import (
	"fmt"
	"github.com/scags9876/go-homework/internal/connect4"
)

// Just used for outputting text to the terminal
const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
	colorCyan  = "\033[36m"
	indent     = "  "
)

func main() {
	defer fmt.Println(colorReset)
	testGame()
}

func testGame() {
	game := connect4.NewGame()

	var succeeded, failed int
	infoMsg("\nRunning connect4 test cases")
	// Illegal move test cases --------------------------------------------------

	infoMsg("Test that column 0 is an illegal move")
	game, err := runTestCase([]int{0}, game)
	if err != nil {
		succeeded++
		successMsg("Attempting to make move to illegal column 0 resulted in the expected error")
	} else {
		failed++
		failureMsg("Attempting to make move to illegal column 0 did not result in the expected error")
	}

	infoMsg("\nTest that column < 0 is an illegal move")
	game = connect4.NewGame()
	game, err = runTestCase([]int{-1}, game)
	if err != nil && err == connect4.ErrIllegalColumn {
		succeeded++
		successMsg("Attempting to make move to illegal column -1 resulted in the expected error")
	} else {
		failed++
		failureMsg("Attempting to make move to illegal column -1 did not result in the expected error")
	}

	infoMsg("\nTest that column > 7 is an illegal move")
	game = connect4.NewGame()
	game, err = runTestCase([]int{8}, game)
	if err != nil && err == connect4.ErrIllegalColumn {
		succeeded++
		successMsg("Attempting to make move to illegal column 8 resulted in the expected error")
	} else {
		failed++
		failureMsg("Attempting to make move to illegal column 8 did not result in the expected error")
	}

	infoMsg("\nTest that attempted move into a full column is an illegal move")
	game = connect4.NewGame()
	game, err = runTestCase([]int{1, 1, 1, 1, 1, 1, 1, 1}, game)
	if err != nil && err == connect4.ErrFullColumn {
		succeeded++
		successMsg("Attempting to make a move to a full column resulted in the expected error")
	} else {
		failed++
		failureMsg("Attempting to make a move to a full collumn did not result in the expected error")
	}

	// Game results test cases --------------------------------------------------

	infoMsg("\nTest that the first player will win with 4 vertical checkers")
	game = connect4.NewGame()
	game, err = runTestCase([]int{1, 2, 1, 2, 1, 2, 1}, game)
	if err == nil && game.IsOver() && game.Result() == connect4.RedPlayer {
		succeeded++
		successMsg("Game reported that Red won the game")
	} else {
		failed++
		failureMsg("Game did not report that Red won the game")
	}

	infoMsg("\nTest to ensure that attemping to make a move after the game is over is illegal")
	game, err = runTestCase([]int{5}, game)
	if err != nil && err == connect4.ErrGameOver {
		succeeded++
		successMsg("Attempting to make a move when the game is over resulted in the expected error")
	} else {
		failed++
		failureMsg("Attempting to make a move when the game is over did result in the expected error")
	}

	infoMsg("\nTest that game isn't over prematurely")
	game = connect4.NewGame()
	game, err = runTestCase([]int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}, game)
	if err == nil && !game.IsOver() {
		succeeded++
		successMsg("Game isn't over prematurely")
	} else {
		failed++
		failureMsg("Game is over in non ended state (or an error was generated when it shouldn't be")
	}

	infoMsg("\nTest that the second player wins with 4 horizontal checker")
	game = connect4.NewGame()
	game, err = runTestCase([]int{7, 1, 1, 3, 7, 4, 3, 2}, game)
	if err == nil && game.IsOver() && game.Result() == connect4.BlackPlayer {
		succeeded++
		successMsg("Game reported that Black won the game")
	} else {
		failed++
		failureMsg("Game did not report that Black won the game")
	}

	infoMsg("\nTest that the game ends in a draw when all positions are filled with no winner")
	game = connect4.NewGame()
	game, err = runTestCase([]int{1, 2, 3, 4, 5, 6, 7, 7, 6, 5, 4, 3, 2, 1, 1, 2, 3, 4, 6, 5, 7, 4, 3, 2, 1, 6, 5, 7, 7, 6, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 6, 7}, game)
	if err == nil && game.IsOver() && game.Result() == connect4.DrawResult {
		succeeded++
		successMsg("Game reported a drawn position")
	} else {
		failed++
		failureMsg("Game did not report a drawn position")
	}

	infoMsg("\nTest that the first player will win with 4 checkers Southwest to Northeast")
	game = connect4.NewGame()
	game, err = runTestCase([]int{1, 2, 6, 3, 2, 5, 4, 4, 3, 5, 3, 5, 5, 6, 6, 7, 7, 4, 4}, game)
	if err == nil && game.IsOver() && game.Result() == connect4.RedPlayer {
		succeeded++
		successMsg("Game reported that Red won the game")
	} else {
		failed++
		failureMsg("Game did not report that Red won the game")
	}

	infoMsg("\nTest that the second player will win with 4 checkers Southeast to Northwest")
	game = connect4.NewGame()
	game, err = runTestCase([]int{2, 6, 3, 5, 4, 1, 2, 5, 2, 2, 3, 4, 3, 3, 1, 4}, game)
	if err == nil && game.IsOver() && game.Result() == connect4.BlackPlayer {
		succeeded++
		successMsg("Game reported that Black won the game")
	} else {
		failed++
		failureMsg("Game did not report that Black won the game")
	}

	printTotalTestResults(succeeded, failed)
}

func printTotalTestResults(succeeded, failed int) {
	total := succeeded + failed
	message := fmt.Sprintf("Passed %d out of %d tests", succeeded, total)

	infoMsg("\n\nTest results:")
	if failed > 0 {
		failureMsg(message)
	} else {
		successMsg(message)
	}
}

func runTestCase(moves []int, game *connect4.Game) (*connect4.Game, error) {
	for i := range moves {
		err := game.MakeMove(moves[i])
		if err != nil {
			return game, err
		}
	}

	return game, nil
}

func successMsg(message string) {
	fmt.Printf("%s%s%s\n", indent, colorGreen, message)
}

func failureMsg(message string) {
	fmt.Printf("%s%s%s\n", indent, colorRed, message)
}

func infoMsg(message string) {
	fmt.Printf("%s%s\n", colorCyan, message)
}
