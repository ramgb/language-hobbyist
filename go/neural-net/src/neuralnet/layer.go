package neuralnet

type Layer struct {
	Index       int
	Perceptrons []Perceptron
}

func NewLayer(index int, inputDimensions int, numNeurons int) Layer {
	layer := Layer{
		Perceptrons: make([]Perceptron, numNeurons),
	}

	for i := range numNeurons {
		layer.Perceptrons[i] = *NewPerceptron(inputDimensions)
	}
	return layer
}

func (l *Layer) Activate(inputs []float64) []float64 {
	outputs := make([]float64, len(l.Perceptrons))

	for i, p := range l.Perceptrons {
		outputs[i] = p.Activate(inputs)
	}
	return outputs
}
