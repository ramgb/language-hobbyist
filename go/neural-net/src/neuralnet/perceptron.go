package neuralnet

import (
	"math"
	"math/rand/v2"
)

type ActivationFnType int

type WeightInitType int

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

type Perceptron struct {
	weights            []float64
	bias               float64
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
		return 1.0 / (1.0 + math.Exp(-sum))
	case ReluActivation:
		panic("Not implemented")
	default:
		return 0.0
	}
}

func (p *Perceptron) GetWeight(weightIndex int) float64 {
	return p.weights[weightIndex]
}

func (p *Perceptron) UpdateWeight(weightIndex int, newWeight float64) {
	p.weights[weightIndex] = newWeight
}
