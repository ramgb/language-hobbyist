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
		fmt.Printf(fmt.Sprintf("%s -> %s", source, destination), "\n")
		return
	}
	solve(numDisks-1, source, intermediate, destination)
	fmt.Printf(fmt.Sprintf("%s -> %s", source, destination), "\n")
	solve(numDisks-1, intermediate, destination, source)
}
