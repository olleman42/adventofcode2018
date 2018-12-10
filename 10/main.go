package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
	h.Close()

	rows := strings.Split(string(d), "\r\n")

	points := make([]*point, 0)
	for _, r := range rows {
		re := regexp.MustCompile("[0-9]+|-[0-9]+")

		c := re.FindAllString(r, -1)
		// fmt.Println(c)
		// fmt.Println(len(c))
		x, err := strconv.ParseInt(c[0], 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseInt(c[1], 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		velx, err := strconv.ParseInt(c[2], 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		vely, err := strconv.ParseInt(c[3], 0, 64)
		if err != nil {
			log.Fatal(err)
		}

		p := point{
			x:    x,
			y:    y,
			velx: velx,
			vely: vely,
		}
		points = append(points, &p)
	}

	for _, p := range points {
		fmt.Println(p)
	}

	// for i := 0; i < 11000; i++ {
	// 	for _, p := range points {
	// 		p.Step()
	// 	}
	// }

	// for _, p := range points {
	// 	fmt.Println(p)
	// }
	width := int64(10000000)
	i := 0
	for {
		for _, p := range points {
			p.Step()
		}
		// find the longest distance between left and right
		leftbound := int64(0)
		rightbound := int64(0)
		for _, p := range points {
			if p.x < leftbound {
				leftbound = p.x
			}
			if p.x > rightbound {
				rightbound = p.x
			}
		}
		newWidth := rightbound - leftbound
		// if newWidth > width {
		// 	fmt.Println(i)
		// 	break
		// }
		width = newWidth
		fmt.Println(width)
		if i > 10556 {
			fmt.Println(i)
			break
		}
		i++
	}

	for _, p := range points {
		fmt.Println(p)
	}

	im := image.NewNRGBA(image.Rectangle{Max: image.Point{X: 400, Y: 400}})
	for _, p := range points {
		im.Set(int(p.x), int(p.y), color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	}

	ih, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(ih, im); err != nil {
		log.Fatal(err)
	}
	ih.Close()
	// im.
}

type point struct {
	x, y, velx, vely int64
}

func (p *point) Step() {
	p.x += p.velx
	p.y += p.vely
}
