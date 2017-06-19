# GoLife

An implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in Go.

GoLife uses a Binary Search Tree data structure to store all active cells in the game.
The playing field is 256x256 because the coordinates are stored as 8-bit integers.
Changing the data type of the coordinates in the Cell stucture will increase the size of the playing field.

GoLife runs Conway's Game of Life directly on the terminal window.
The basic logic for the Game of Life has been implemented, however there are many features that I wish to add that have not been added yet.

## Installation

To install GoLife, simply use the `go get` command.

```
go get github.com/DougWagner/golife
```
