package neuralnet

import (
	"fmt"
	"math"
)

type Network struct {
	layers []*Layer
}

func NewNetwork(layerDimensions []int, inputDimensions int) *Network {
	network := &Network{
		layers: make([]*Layer, len(layerDimensions)),
	}
	layerInputDimensions := inputDimensions
	for index := range len(layerDimensions) {
		network.layers[index] = NewLayer(index, layerInputDimensions, layerDimensions[index])
		layerInputDimensions = layerDimensions[index]
	}
	return network
}

func (n *Network) Activate(inputs []float64) []float64 {
	layerInputs := inputs
	var layerOutputs []float64
	for _, layer := range n.layers {
		layerOutputs = layer.Activate(layerInputs)
		layerInputs = layerOutputs
	}
	return layerOutputs
}

func (n *Network) Print() {
	fmt.Println("Network State")
	for index := range n.layers {
		n.layers[index].Print()
	}
	fmt.Println("--------------")
}

func (n *Network) Train(inputs [][]float64, expectedOutput [][]float64, learningRate float64) {
	if len(inputs) == 0 || len(expectedOutput) == 0 {
		panic("No training input/output provided")
	}
	if len(inputs) != len(expectedOutput) {
		panic("Cardinality of training inputs and output do not match")
	}

	if n.layers[len(n.layers)-1].Size() != len(expectedOutput[0]) {
		panic("Output layer cardinality is not equal to expected output cardinality")
	}

	// Implement stochastic gradient

	squaredError := 1.0
	iterations := 0
	for squaredError > 0.001 {
		fmt.Println("Iteration #", iterations)
		iterations += 1
		squaredError = 0.0
		for index := range inputs {
			actualOutput := n.Activate(inputs[index])
			partialGradientAccumulators := make([]float64, len(expectedOutput[0]))
			for outputIndex := range actualOutput {
				squaredError += 0.5 * math.Pow((actualOutput[outputIndex]-expectedOutput[index][outputIndex]), 2)
				partialGradientAccumulators[outputIndex] = actualOutput[outputIndex] - expectedOutput[index][outputIndex]
			}

			for i := len(n.layers) - 1; i >= 0; i-- {
				partialGradientAccumulators = n.layers[i].BackPropagate(partialGradientAccumulators, learningRate)
			}

		}
		fmt.Println("squaredError = ", squaredError)
	}
	fmt.Println("#iterations completed = ", iterations)
}
