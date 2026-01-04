package main

import (
	"fmt"
	"neural-net/src/perceptron"
)

func main() {
	newPerceptron := perceptron.NewPerceptron(2)
	output := newPerceptron.Activate([]float64{1.0, 1.0})
	fmt.Println(output)
}
