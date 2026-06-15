package main

import (
	"fmt"
	"neural-net/src/neuralnet"
)

func main() {
	newPerceptron := neuralnet.NewPerceptron(2)
	output := newPerceptron.Activate([]float64{1.0, 1.0})
	fmt.Println(output)
}
