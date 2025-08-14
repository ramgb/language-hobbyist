package hanoi

import (
	"fmt"
)

func Start() {
	var numDisks int
	fmt.Println("Enter number of disks : ")
	fmt.Scanln(&numDisks)
	solve(numDisks, "A", "C", "B")
}

func solve(numDisks int, source string, destination string, intermediate string) {
	if numDisks == 1 {
		fmt.Println(source + " -> " + destination)
		return
	}
	solve(numDisks-1, source, intermediate, destination)
	fmt.Println(source + " -> " + destination)
	solve(numDisks-1, intermediate, destination, source)
}
