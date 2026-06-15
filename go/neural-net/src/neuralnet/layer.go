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

	for index, perceptron := range l.perceptrons {
		outputs[index] = perceptron.Activate(inputs)
	}
	return outputs
}

func (l *Layer) BackPropagate(partialGradientAccumulators []float64, nextLayerOutputs []float64) ([]float64, []float64) {
	// iterate through all perceptrons in this layer and update their weights.
	// return new partial gradients and outputs for next layers.

	newPartialGradientAccumulators := make([]float64, len(l.perceptrons))
	newNextLayerOutputs := make([]float64, len(l.perceptrons))

	for index, perceptron := range l.perceptrons {
		newPartialGradientAccumulators[index], newNextLayerOutputs[index] = perceptron.UpdateWeights(partialGradientAccumulators, nextLayerOutputs)
	}
	return newPartialGradientAccumulators, newNextLayerOutputs
}
