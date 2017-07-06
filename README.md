# GoLife

An implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in Go.

GoLife uses a Binary Search Tree data structure to store all active cells in the game.
The playing field is 256x256 because the coordinates are stored as 8-bit integers.
Changing the data type of the coordinates in the Cell stucture will increase the size of the playing field.

GoLife runs Conway's Game of Life directly on the terminal window.
Use the arrow keys to scroll up, down, left, or right during execution of GoLife.
GoLife is still a work in progress, and there are many features that I wish to add that have not been added yet.

## Installation

To install GoLife, simply use the `go get` command.

```
go get github.com/DougWagner/golife
```

GoLife works on Linux systems (developed and tested on Ubuntu).
It builds and runs on MacOS systems, however there are some issues.
It does not compile on Windows systems.
