package neuralnet

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

// TODO#5) : backpropagation implementation
