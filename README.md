# GoLife

An implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in Go.

GoLife uses a Binary Search Tree data structure to store all active cells in the game.
The playing field is 4294967295 x 4294967295 (uint32max x uint32max) in size because the grid of cells is not stored as a static array.

GoLife runs Conway's Game of Life directly on the terminal window.
The basic logic for the Game of Life has been implemented, however there are many features that I wish to add that have not been added yet.

## Installation

To install GoLife, simply use the `go get` command.

```
go get github.com/DougWagner/golife
```
