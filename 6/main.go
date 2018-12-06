package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	id, x, y int
}

func main() {
	h, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(h)
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(data), "\r\n")
	log.Println("starting")
	points, err := parsePoints(rows)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range points {
		log.Println(*p)
	}

	field := make([][]point, 0)
	size := 400
	closer := 0
	// field size 400-400 (just eyeing it)
	// walk through each point on field and find closest point
	for i := 0; i < size; i++ {

		field = append(field, make([]point, 0))
		for j := 0; j < size; j++ {
			// find closes point by manhattan distance
			// closest := 999
			thispoint := &point{x: j, y: i}
			totDistance := 0
			for _, p := range points {
				totDistance = totDistance + manhattanDistance(thispoint, p)
			}
			if totDistance < 10000 {
				closer++
			}
			// for _, p := range points {
			// 	if manhattanDistance(thispoint, p) < closest {
			// 		thispoint.id = p.id
			// 		closest = manhattanDistance(thispoint, p)
			// 	}

			// }
			// // check if equally close to another one
			// for _, p := range points {
			// 	if manhattanDistance(thispoint, p) == closest && p.id != thispoint.id {
			// 		thispoint.id = -1
			// 	}
			// }
			field[i] = append(field[i], *thispoint)

		}
	}
	log.Println(closer)
	// log.Println(field)
	// log.Println("done")
	// for i := range field {
	// 	for j := range field {
	// 		fmt.Print(field[i][j].id, "\t")
	// 	}
	// 	fmt.Println()
	// }

	// trash := make(map[int]*point)
	// // remove all edge ids
	// for _, p := range field[0] {
	// 	if p.id == -1 {
	// 		continue
	// 	}
	// 	trash[p.id] = points[p.id]
	// }
	// for _, p := range field[len(field)-1] {
	// 	if p.id == -1 {
	// 		continue
	// 	}
	// 	trash[p.id] = points[p.id]
	// }
	// for i := 0; i < size; i++ {
	// 	p := field[i][0]
	// 	if p.id == -1 {
	// 		continue
	// 	}
	// 	trash[p.id] = points[p.id]
	// 	p = field[i][size-1]
	// 	if p.id == -1 {
	// 		continue
	// 	}
	// 	trash[p.id] = points[p.id]
	// }

	// fmt.Println(trash)

	// tally := make(map[int]int)
	// // count number of ids excluding the trash
	// for y := range field {
	// 	for x := range field[y] {
	// 		if field[y][x].id == -1 {
	// 			continue
	// 		}
	// 		if _, ok := trash[field[y][x].id]; ok {
	// 			continue
	// 		}
	// 		if _, ok := tally[field[y][x].id]; ok {
	// 			tally[field[y][x].id]++
	// 		} else {
	// 			tally[field[y][x].id] = 1
	// 		}
	// 	}
	// }
	// fmt.Println(tally)
}

func manhattanDistance(a, b *point) int {
	x := 0
	y := 0
	if a.x > b.x {
		x = a.x - b.x
	} else {
		x = b.x - a.x
	}

	if a.y > b.y {
		y = a.y - b.y
	} else {
		y = b.y - a.y
	}

	return x + y
}

func parsePoints(rows []string) ([]*point, error) {
	points := make([]*point, 0)
	for i, r := range rows {
		xy := strings.Split(r, ", ")
		x, err := strconv.ParseInt(xy[0], 0, 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseInt(xy[1], 0, 64)
		if err != nil {
			return nil, err
		}

		p := &point{}
		p.x = int(x)
		p.y = int(y)
		p.id = i
		points = append(points, p)
	}
	return points, nil
}
