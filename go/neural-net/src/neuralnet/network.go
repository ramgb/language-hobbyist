package neuralnet

import "math"

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

func (n *Network) Train(inputs [][]float64, outputs [][]float64, learningRate float64) {
	if len(inputs) == 0 || len(outputs) == 0 {
		panic("No training input/output provided")
	}
	if len(inputs) != len(outputs) {
		panic("Cardinality of training inputs and outputs do not match")
	}

	if n.layers[len(n.layers)-1].Size() != len(outputs[0]) {
		panic("Output layer cardinality is not equal to expected output cardinality")
	}

	// Implement stochastic gradient
	for index := range inputs {
		output := n.Activate(inputs[index])

		squaredError := 0.0
		for outputIndex := range output {
			squaredError += 0.5 * math.Pow((output[outputIndex]-outputs[index][outputIndex]), 2)
		}
	}
}
