package neuralnet

type Layer struct {
	index       int
	perceptrons []*Perceptron
}

func NewLayer(layerIndex int, inputDimensions int, numNeurons int) *Layer {
	layer := &Layer{
		index:       layerIndex,
		perceptrons: make([]*Perceptron, numNeurons),
	}

	for i := range numNeurons {
		layer.perceptrons[i] = NewPerceptron(inputDimensions)
	}
	return layer
}

func (l *Layer) Size() int {
	return len(l.perceptrons)
}

func (l *Layer) Activate(inputs []float64) []float64 {
	outputs := make([]float64, len(l.perceptrons))

	for i, p := range l.perceptrons {
		outputs[i] = p.Activate(inputs)
	}
	return outputs
}
