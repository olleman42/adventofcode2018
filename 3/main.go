package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	h, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	d, err := ioutil.ReadAll(h)
	if err != nil {
		log.Fatal(err)
	}

	pieces, err := parsePieces(d)
	if err != nil {
		log.Fatal(err)
	}

	xoverlaps := 0
	canvas := [1200][1200]int{}
	for i := 0; i < len(canvas); i++ {
		row := canvas[i]
		for j := 0; j < len(row); j++ {
			// totContains := 0
			overlap := make([]*piece, 0)
			for _, p := range pieces {
				// if p.contains()
				if p.contains(int64(i), int64(j)) {
					// totContains++
					overlap = append(overlap, p)
				}
			}
			if len(overlap) > 1 {
				for _, px := range overlap {
					px.claimed = true
				}
				xoverlaps++
				// log.Println("found overlap", i, j)
			}

		}
	}

	log.Println(xoverlaps)
	for _, p := range pieces {
		if !p.claimed {
			log.Println("untouched:", p.id)
		}
	}

}

func parsePieces(in []byte) ([]*piece, error) {
	rows := strings.Split(string(in), "\n")

	pieces := make([]*piece, 0)
	for _, w := range rows {
		piece := piece{}
		parts := strings.Split(w, " ")

		id, err := strconv.ParseInt(parts[0][1:], 0, 64)
		if err != nil {
			return nil, err
		}

		xy := strings.Split(parts[2][0:len(parts[2])-1], ",")
		x, err := strconv.ParseInt(xy[0], 0, 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseInt(xy[1], 0, 64)
		if err != nil {
			return nil, err
		}

		wh := strings.Split(parts[3], "x")
		width, err := strconv.ParseInt(wh[0], 0, 64)
		if err != nil {
			return nil, err
		}
		height, err := strconv.ParseInt(wh[1], 0, 64)
		if err != nil {
			return nil, err
		}

		piece.id = id
		piece.x = x
		piece.y = y
		piece.width = width
		piece.height = height

		pieces = append(pieces, &piece)
	}
	return pieces, nil
}

type piece struct {
	id, width, height, x, y int64
	claimed                 bool
}

func (p *piece) contains(x, y int64) bool {
	if p.x < x && x <= p.x+p.width && p.y < y && y <= p.y+p.height {
		// p.claimed = true
		return true
	}
	return false
}
