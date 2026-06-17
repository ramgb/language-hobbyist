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
	activatedInput     []float64
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
		activatedInput:     make([]float64, len(weights)),
		activatedOutput:    0.0,
		activationFunction: activationFunctionType,
	}
}

func (p *Perceptron) Size() int {
	return len(p.weights)
}

func (p *Perceptron) Activate(inputs []float64) float64 {
	p.activatedInput = inputs
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

func (p *Perceptron) UpdateWeights(partialGradientAccumulator float64, learningRate float64) []float64 {
	// Update each weight based on partial gradients of the next layer.

	if p.state != Active {
		panic("Activation function should already be run before updating weights")
	}
	currentNodeGradientAccumulator := make([]float64, len(p.weights))
	currentActivatedOutput := p.activatedOutput
	for index := range p.weights {
		currentNodeGradientAccumulator[index] = p.weights[index] * partialGradientAccumulator
		p.weights[index] -= learningRate * partialGradientAccumulator * currentActivatedOutput * (1.0 - currentActivatedOutput) * p.activatedInput[index]
	}

	p.state = InActive
	p.activatedOutput = 0.0
	return currentNodeGradientAccumulator
}
