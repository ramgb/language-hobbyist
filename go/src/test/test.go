package test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	fmt.Println("This is a test package")
}

func Tester() {
	var name string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter input string : ")
	name, _ = reader.ReadString('\n')
	name = name[:len(name)-1]
	fmt.Println(name, " has ", len(strings.ReplaceAll(name, " ", "")), " characters")
}
