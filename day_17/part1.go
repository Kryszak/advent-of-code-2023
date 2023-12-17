package day17

import (
	"fmt"
	"github.com/Kryszak/aoc2023/common"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

type node struct {
	x, y  int
	value int
}

func printMap(nodes [][]node) {
	for _, rowValue := range nodes {
		for _, colValue := range rowValue {
			fmt.Print(colValue)
		}
		fmt.Println()
	}
}

func loadInput(path string) (nodes [][]node) {
	fileScanner := common.FileScanner(path)
	x := 0
	for fileScanner.Scan() {
		var line []node
		for y, value := range fileScanner.Text() {
			line = append(line, node{x, y, common.Atoi(string(value))})
		}
		nodes = append(nodes, line)
		x++
	}

	return nodes
}

func bfs(nodes [][]node, endNode node, minStraight, maxStraight int) int {
	type queueEntry struct {
		x, y     int
		dir      common.Direction
		heatLoss int
		straight int
	}
	type cacheEntry struct {
		x, y     int
		dir      common.Direction
		straight int
	}

	q := pq.NewWith(func(a, b any) int {
		p1 := a.(queueEntry).heatLoss
		p2 := b.(queueEntry).heatLoss
		return p1 - p2
	})

	q.Enqueue(queueEntry{
		x:        0,
		y:        1,
		straight: 1,
		dir:      common.East,
	})
	q.Enqueue(queueEntry{
		x:        1,
		y:        0,
		straight: 1,
		dir:      common.South,
	})
	cache := make(map[cacheEntry]int)

	for !q.Empty() {
		t, _ := q.Dequeue()
		entry := t.(queueEntry)

		if entry.x < 0 || entry.x >= len(nodes) || entry.y < 0 || entry.y >= len(nodes[entry.x]) {
			continue
		}

		heat := nodes[entry.x][entry.y].value + entry.heatLoss
		if entry.x == endNode.x && entry.y == endNode.y {
			return heat
		}

		ce := cacheEntry{x: entry.x, y: entry.y, dir: entry.dir, straight: entry.straight}
		if v, exists := cache[ce]; exists {
			if v <= heat {
				continue
			}
		}
		cache[ce] = heat

		if entry.straight >= minStraight {
			xLeft, yLeft := entry.x, entry.y
			xRight, yRight := entry.x, entry.y
			var dirLeft, dirRight common.Direction

			switch entry.dir {
			case common.North, common.South:
				{
					yLeft--
					dirLeft = common.West

					yRight++
					dirRight = common.East
				}
			case common.East, common.West:
				{
					xLeft--
					dirLeft = common.North

					xRight++
					dirRight = common.South
				}
			}

			q.Enqueue(queueEntry{
				x:        xLeft,
				y:        yLeft,
				dir:      dirLeft,
				heatLoss: heat,
				straight: 1,
			})

			q.Enqueue(queueEntry{
				x:        xRight,
				y:        yRight,
				dir:      dirRight,
				heatLoss: heat,
				straight: 1,
			})
		}

		if entry.straight < maxStraight {
			x, y := entry.x, entry.y
			switch entry.dir {
			case common.North:
				x--
			case common.South:
				x++
			case common.East:
				y++
			case common.West:
				y--
			}
			q.Enqueue(queueEntry{
				x:        x,
				y:        y,
				dir:      entry.dir,
				heatLoss: heat,
				straight: entry.straight + 1,
			})
		}
	}
	return 0
}

func Part1(path string) (answer int) {
	input := loadInput(path)
	answer = bfs(input, input[len(input)-1][len(input[0])-1], 1, 3)
	return answer
}
