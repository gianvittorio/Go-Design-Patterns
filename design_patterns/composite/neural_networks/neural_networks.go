package main

type NeuronInterface interface {
	Iterate() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Iterate() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (nl *NeuronLayer) Iterate() []*Neuron {
	result := make([]*Neuron, 0)

	for _, neuron := range nl.Neurons {
		result = append(result, &neuron)
	}

	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iterate() {
		for _, r := range right.Iterate() {
			l.ConnectTo(r)
		}
	}
}

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)
}
