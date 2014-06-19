package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Perceptron struct {
	output float64
	Weight []float64
}

type NeuralNet struct {
	MiddleNode []Perceptron
	OutputNode []Perceptron
}

func nodeInit(weightNum int) Perceptron {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var node Perceptron
	node.Weight = make([]float64, weightNum)
	for i := 0; i < weightNum; i++ {
		node.Weight[i] = (rnd.Float64() * 2.0) - 1.0
	}
	return node
}

func makeNodes(weightNum int, nodeNum int) []Perceptron {
	var nodes = make([]Perceptron, nodeNum)
	for i, _ := range nodes {
		nodes[i] = nodeInit(weightNum)
	}
	return nodes
}

func (net *NeuralNet) init(inputNum int, middleNum int, outputNum int) {
	net.MiddleNode = makeNodes(inputNum, middleNum)
	net.OutputNode = makeNodes(middleNum, outputNum)
}

func (net *NeuralNet) calc(inputs []float64) {
	var midInputs []float64
	midInputs = calcOutput(inputs, net.MiddleNode)
	for i, v := range midInputs {
		fmt.Printf("midInputs[%d] = %f\n", i, v)
	}

	var outputs []float64
	outputs = calcOutput(midInputs, net.OutputNode)
	for i, v := range outputs {
		fmt.Printf("outputs[%d] = %f\n", i, v)
	}

}

func calcOutput(inputs []float64, nodes []Perceptron) []float64 {
	var outputs = make([]float64, len(nodes))
	for i, node := range nodes {
		var sum float64 = 0
		for j, input := range inputs {
			sum += node.Weight[j] * input
		}
		if sigmoid(sum) > 0.5 {
			outputs[i] = 1.0
		} else {
			outputs[i] = 0.0
		}
	}
	return outputs
}

func sigmoid(value float64) float64 {
	gain := -5.0
	return (1.0 / (1 + math.Pow(math.E, gain*value)))
}

func main() {
	var nn NeuralNet
	nn.init(2, 3, 1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("w[%d][%d] = %f\n", i, j, nn.MiddleNode[i].Weight[j])
		}
	}
	input := []float64{1.0, 1.0}
	nn.calc(input)
}
