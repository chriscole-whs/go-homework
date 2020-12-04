# connect-4-go-homework
This is a take home assignment to complete an implementation of Connect 4 in Go.

## Overview
This package defines a set of methods, constants, and variables in the connect4 package.
main.go acts as a driver for a set of tests to ensure that the logic in the template methods
have been implemented correctly.

## Running the Tests
From the base directory of the project:
```
go run -v main.go
```

In it's current unmodified state, the output should look like this:
```
$ go run -v main.go

Running connect4 test cases
Test that column 0 is an illegal move
  Attempting to make move to illegal column 0 did not result in the expected error

Test that column < 0 is an illegal move
  Attempting to make move to illegal column -1 did not result in the expected error

Test that column > 7 is an illegal move
  Attempting to make move to illegal column 8 did not result in the expected error

Test that attempted move into a full column is an illegal move
  Attempting to make a move to a full collumn did not result in the expected error

Test that the first player will win with 4 vertical checkers
  Game did not report that Red won the game

Test to ensure that attemping to make a move after the game is over is illegal
  Attemping to make a move when the game is over did result in the expected error

Test that game isn't over prematurely
  Game isn't over prematurely

Test that the second player wins with 4 horizontal checkers
  Game did not report that Black won the game

Test that the game ends in a draw when all positions are filled with no winner
  Game did not report a drawn position

Test that the first player will win with 4 checkers Southwest to Northeast
  Game did not report that Red won the game

Test that the second player will win with 4 checkers Southeast to Northwest
  Game did not report that Black won the game


Test results:
  Passed 1 out of 11 tests
```

### Submission Instructions
Clone this repository, and make a branch of your own.
Come up with a design for the connect4.Game struct and implement the methods in the connect4 package. 
Once you have an implementation that passes the tests in main.go, create a gist (https://gist.github.com/) for
the connect4 package and email it back to your contact at WhiteHat.

### Tips
You shouldn't need any libraries outside of the core language to complete this task. 
A working solution isn't your only goal here; special attention needs to be paid to good software
development practices like being conscious of performance, readability, encapsulation, and selecting appropriate, informative names for 
functions and variables.

Feel free to include any data members in the Game struct that you feel are needed to complete the exercise.

Good luck!

## Reference Materials
Golang: https://golang.org/doc/

Connect 4 behavior: https://www.mathsisfun.com/games/connect4.html
