package day20

import (
	"slices"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

type module struct {
	name    string
	t       string
	inputs  []input
	outputs []string
	state   bool
}

type input struct {
	name  string
	state bool
}

type pulseCounter struct {
	lows  int
	highs int
}

func (c *pulseCounter) increment(signal bool) {
	if signal {
		c.highs++
	} else {
		c.lows++
	}
}

type queuedSignal struct {
	node   module
	sender string
	signal bool
}

const (
	low  = false
	high = true
)

func loadData(path string) map[string]module {
	scanner := common.FileScanner(path)

	machines := make(map[string]module)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		var name, t string
		if line[0][0] == '%' || line[0][0] == '&' {
			t = string(line[0][0])
			name = string(line[0][1:])
		} else {
			name = line[0]
		}
		var outputs []string
		for _, output := range strings.Split(line[1], ", ") {
			outputs = append(outputs, output)
		}
		machines[name] = module{name, t, make([]input, 0), outputs, false}
	}

	for key := range machines {
		for mod, value := range machines {
			if slices.Contains(value.outputs, key) {
				entry := machines[key]
				entry.inputs = append(entry.inputs, input{mod, false})
				machines[key] = entry
			}
		}
	}

	return machines
}

func processPulse(machines map[string]module, counter *pulseCounter) []queuedSignal {
	broadcaster := machines["broadcaster"]
	queue := []queuedSignal{{broadcaster, "button", low}}
	signals := make([]queuedSignal, 0)

	for len(queue) > 0 {
		current := queue[0]
		signals = append(signals, current)
		queue = queue[1:]
		counter.increment(current.signal)

		switch current.node.t {
		case "":
			for _, target := range current.node.outputs {
				queue = append(queue, queuedSignal{machines[target], current.node.name, low})
			}
		case "%":
			if current.signal == low {
				this := machines[current.node.name]
				this.state = !this.state
				machines[current.node.name] = this
				for _, target := range current.node.outputs {
					queue = append(queue, queuedSignal{machines[target], current.node.name, this.state})
				}
			}
		case "&":
			inputIndex := slices.IndexFunc(current.node.inputs, func(i input) bool {
				return i.name == current.sender
			})
			current.node.inputs[inputIndex].state = current.signal
			nextPulse := low
			for _, input := range current.node.inputs {
				if input.state != high {
					nextPulse = high
				}
			}
			for _, target := range current.node.outputs {
				queue = append(queue, queuedSignal{machines[target], current.node.name, nextPulse})
			}
		}
	}
	return signals
}

func processPulses(machines map[string]module, counter *pulseCounter, times int) {
	for i := 0; i < times; i++ {
		processPulse(machines, counter)
	}
}

func Part1(path string) (answer int) {
	machines := loadData(path)
	pulseCounter := pulseCounter{0, 0}

	processPulses(machines, &pulseCounter, 1000)

	answer = pulseCounter.lows * pulseCounter.highs

	return answer
}
