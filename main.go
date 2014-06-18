package main

import (
	"fmt"
	"math/rand"
)

type Perceptron struct {
	output int
	Weight []float64
}

func (node *Perceptron) init(weightNum int) {
	node.Weight = make([]float64, weightNum)
	for i := 0; i < weightNum; i++ {
		node.Weight[i] = rand.Float64()
	}
}

func main() {
	node := new(Perceptron)
	node.init(6)
	for i := 0; i < len(node.Weight); i++ {
		fmt.Println(node.Weight[i])
	}
	// create network

	// learning

	// wait input

}
