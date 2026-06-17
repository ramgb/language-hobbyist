package main

import (
	"fmt"
	"neural-net/src/neuralnet"
)

func main() {
	network := neuralnet.NewNetwork([]int{2, 1}, 2)
	trainingInput := [][]float64{{0.0, 0.0}, {0.0, 1.0}, {1.0, 0.0}, {1.0, 1.0}}
	trainingOutput := [][]float64{{0.0}, {1.0}, {1.0}, {0.0}}
	network.Train(trainingInput, trainingOutput, 0.1)

	for index := range trainingInput {
		output := network.Activate(trainingInput[index])
		fmt.Print("output for ", trainingInput[index], " = ", output)
	}
}
