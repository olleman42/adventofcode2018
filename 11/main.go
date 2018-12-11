package main

import (
	"fmt"
	"log"
	"strconv"
)

var (
	debug = false
)

func main() {
	serial := 7857
	// serial := 18
	// serial := 18

	size := 300

	field := make(map[int]map[int]*point)

	for y := 1; y <= size; y++ {
		// field = append(field, []point{})
		field[y] = make(map[int]*point)
		for x := 1; x <= size; x++ {
			p := &point{
				x: x,
				y: y,
			}

			rackID := x + 10
			powerLevel := rackID * y
			powerLevel = powerLevel + serial
			powerLevel = powerLevel * rackID
			plString := strconv.Itoa(powerLevel)
			// fmt.Println(plString)
			if len(plString) < 3 {
				fmt.Println(plString)
			}
			hundredth, err := strconv.ParseInt(plString[len(plString)-3:len(plString)-2], 0, 64)
			// fmt.Println(hundredth)
			if err != nil {
				log.Fatal(err)
			}
			powerLevel = int(hundredth) - 5
			p.level = powerLevel
			p.lastSize = p.level

			field[y][x] = p
			// field[y] = append(field[y], p)

		}
		// fmt.Println(field[y])
		// os.Exit(0)
	}

	highest := 0
	highX, highY, highSize := 0, 0, 0
	for dsize := 1; dsize <= size; dsize++ {
		fmt.Println(dsize)
		for y := 1; y <= size; y++ {
			for x := 1; x <= size; x++ {
				pp := field[y][x].GetNeighbourTotal(field, dsize)
				if pp > highest {
					fmt.Println("New top:", pp)
					highest = pp
					// fmt.Println(pp)
					highX = x
					highY = y
					highSize = dsize
				}
			}
		}
	}
	fmt.Println("done", highest, "(", highX, ",", highY, ",", highSize, ")")
	// debug = true
	// fmt.Println(field[16][243].level)
	// fmt.Println(field[16][244].level)
	// fmt.Println(field[16][245].level)

	// fmt.Println(field[17][243].level)
	// fmt.Println(field[17][244].level)
	// fmt.Println(field[17][245].level)

	// fmt.Println(field[18][243].level)
	// fmt.Println(field[18][244].level)
	// fmt.Println(field[18][245].level)
}

type point struct {
	x, y, level, lastSize int
}

func (p *point) GetNeighbourTotal(field map[int]map[int]*point, size int) int {
	if debug {
		fmt.Println(p.x, p.y)
	}

	total := p.lastSize
	if _, ok := field[p.x][p.y+size-1]; !ok {
		return 0
	}
	if _, ok := field[p.x+size-1][p.y]; !ok {
		return 0
	}
	if size > 1 {

		for dy := 0; dy < size-1; dy++ {
			total += getLevel(p.x+size-1, p.y+dy, field)
			// total += thisLevel
		}
		for dx := 0; dx < size-1; dx++ {
			total += getLevel(p.x+dx, p.y+size-1, field)
			// total += thisLevel
		}
		total += getLevel(p.x+size-1, p.y+size-1, field)
	}

	p.lastSize = total
	// p.squareLevel = total
	return total
	// // OWN
	// total := p.level
	// // RIGHT
	// total += getLevel(p.x+1, p.y, field)
	// // RIGHT RIGHT
	// total += getLevel(p.x+2, p.y, field)
	// // DOWN
	// total += getLevel(p.x, p.y+1, field)
	// // DOWN DOWN
	// total += getLevel(p.x, p.y+2, field)
	// // DOWN LEFT
	// total += getLevel(p.x+1, p.y+1, field)
	// // DOWN LEFT LEFT
	// total += getLevel(p.x+2, p.y+1, field)
	// // DOWN DOWN LEFT
	// total += getLevel(p.x+1, p.y+2, field)
	// // DOWN DOWN LEFT LEFT
	// total += getLevel(p.x+2, p.y+2, field)
	// p.squareLevel = total

	// return total
}

func getLevel(x int, y int, field map[int]map[int]*point) int {
	if n, ok := field[y][x]; ok {
		if debug {
			fmt.Println(x, y, n.level)
		}
		return n.level
	}
	return 0
}
