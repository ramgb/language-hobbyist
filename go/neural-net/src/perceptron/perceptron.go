package perceptron

import "math"

type ActivationFunctionType int

const (
	StepActivation ActivationFunctionType = iota
	SigmoidActivation
)

type Perceptron struct {
	weights            []float64
	bias               float64
	activationFunction ActivationFunctionType
}

func NewPerceptron(inputDimensions int) *Perceptron {
	return NewPerceptronWithActivationFn(inputDimensions, StepActivation)
}

func NewPerceptronWithActivationFn(inputDimensions int, activationFunctionType ActivationFunctionType) *Perceptron {
	return &Perceptron{
		weights:            make([]float64, inputDimensions),
		bias:               0.0,
		activationFunction: activationFunctionType,
	}
}

func (p *Perceptron) Activate(inputs []float64) float64 {
	sum := p.bias
	for i, input := range inputs {
		sum += p.weights[i] * input
	}
	switch p.activationFunction {
	case StepActivation:
		if sum >= 0 {
			return 1.0
		}
		return 0.0
	case SigmoidActivation:
		return 1.0 / (1.0 + math.Exp(-sum))
	default:
		return 0.0
	}
}
