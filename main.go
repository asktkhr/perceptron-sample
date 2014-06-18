package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Perceptron struct {
	output int
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
		node.Weight[i] = rnd.Float64()
	}
	return node
}

func makeNodes(weightNum int, nodeNum int) []Perceptron {
	var nodes = make([]Perceptron, nodeNum)
	for i := 0; i < nodeNum; i++ {
		nodes[i] = nodeInit(weightNum)
	}
	return nodes
}

func (net *NeuralNet) init(inputNum int, middleNum int, outputNum int) {
	net.MiddleNode = makeNodes(inputNum, middleNum)
	net.OutputNode = makeNodes(middleNum, outputNum)
}

/*
func (net *NeuralNet) calc(input []int) int {

}
*/

func main() {
	var nn NeuralNet
	nn.init(2, 3, 1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("w[%d][%d] = %f\n", i, j, nn.MiddleNode[i].Weight[j])
		}
	}

}
