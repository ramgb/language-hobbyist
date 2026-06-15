package neuralnet

import (
	"math"
	"math/rand/v2"
)

type ActivationFnType int

type WeightInitType int

type ActivationState int

// For StandardGaussian
const sigma = 0.1
const mu = 0.0

const (
	SigmoidActivation ActivationFnType = iota
	ReluActivation
)

const (
	StandardGaussian WeightInitType = iota
)

const (
	InActive ActivationState = iota
	Active
)

type Perceptron struct {
	weights            []float64
	bias               float64
	state              ActivationState
	activatedOutput    float64
	activationFunction ActivationFnType
}

func NewPerceptron(inputDimensions int) *Perceptron {
	return NewPerceptronWithActivationFn(inputDimensions, SigmoidActivation, StandardGaussian)
}

func NewPerceptronWithActivationFn(inputDimensions int, activationFunctionType ActivationFnType, weightInitType WeightInitType) *Perceptron {
	weights := make([]float64, inputDimensions)

	for index := range len(weights) {
		weights[index] = rand.NormFloat64()*sigma + mu
	}
	return &Perceptron{
		weights:            weights,
		bias:               0.0,
		state:              InActive,
		activatedOutput:    0.0,
		activationFunction: activationFunctionType,
	}
}

func (p *Perceptron) Activate(inputs []float64) float64 {
	sum := p.bias

	if len(inputs) != len(p.weights) {
		panic("Number of inputs doesn't match the number of weights")
	}
	for i, input := range inputs {
		sum += p.weights[i] * input
	}
	switch p.activationFunction {
	case SigmoidActivation:
		p.activatedOutput = 1.0 / (1.0 + math.Exp(-sum))
		p.state = Active
		return p.activatedOutput
	case ReluActivation:
		panic("Not implemented")
	default:
		panic("No activation function set")
	}
}

func (p *Perceptron) UpdateWeights(partialGradientAccumulator []float64, nextLayerOutputs []float64) (float64, float64) {
	// Update each weight based on partial gradients of the next layer.

	if p.state != Active {
		panic("Activation function should already be run before updating weights")
	}
	currentNodeGradientAccumulator := 0.0
	currentActivatedOutput := p.activatedOutput
	for index := range p.weights {
		currentNodeGradientAccumulator += p.weights[index] * partialGradientAccumulator[index]
		p.weights[index] = partialGradientAccumulator[index] * nextLayerOutputs[index] * (1.0 - nextLayerOutputs[index]) * currentActivatedOutput
	}

	p.state = InActive
	p.activatedOutput = 0.0
	return currentNodeGradientAccumulator, currentActivatedOutput
}
