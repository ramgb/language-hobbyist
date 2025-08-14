package knights

import (
	"fmt"
	"strconv"
)

func Start() {
	var boardSize int
	fmt.Println("Enter board size: ")
	fmt.Scanln(&boardSize)
	fmt.Println("Solving Knights Tour for " + strconv.Itoa(boardSize) + " X " + strconv.Itoa(boardSize))
}
