package main

import (
	"fmt"
	"github.com/scags9876/go-homework/internal/connect4"
)

func main() {
	game := connect4.NewGame()

	testGame(game)
}

func testGame(game connect4.Game) {
	moves := []int{ 3, 3, 2, 2, 5, 2 }
	for _, move := range moves {
		err := game.MakeMove(move)
		if err != nil {
			panic(err)
		}
		if game.IsOver() {
			fmt.Println(game.Result())
			return
		}
	}
}