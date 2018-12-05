package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("starting")
	h, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	d, err := ioutil.ReadAll(h)
	if err != nil {
		log.Fatal(err)
	}

	data := string(d)

	sample := "abcdefghijklmnopqrstuvwxyz"
	shortestL := ""
	shortestPolySize := 100000
	for _, l := range sample {
		var newPolymer strings.Builder
		for _, c := range data {
			if c == l || c == rune(l+32) || c == rune(l-32) {
				//newPolymer.WriteString(string(c))
				continue
			}
			newPolymer.WriteString(string(c))

		}
		//log.Println(newPolymer.String())
		reducedSize := reducePolymer(newPolymer.String())
		log.Println(string(l), reducedSize)
		if reducedSize < shortestPolySize {
			shortestL = string(l)
			shortestPolySize = reducedSize
		}
	}
	log.Println(shortestL, shortestPolySize)
	//log.Println(reducePolymer(data))

}

func reducePolymer(polymer string) int {
	original := len(polymer)
	reduced := 0
	for reduced != original {
		var newOne strings.Builder
		skipOne := false

		for i, c := range polymer {
			if skipOne {
				skipOne = false
				continue
			}
			if i == len(polymer)-1 {
				newOne.WriteRune(c)
				continue
			}

			if c < 96 {
				//log.Println(c, rune(data[i+1]-32))
				if c == rune(polymer[i+1]-32) {
					skipOne = true
					continue
				}
				//log.Println(i, rune(c+32))
			} else {
				//log.Println(c, rune(data[i+1]+32))
				if c == rune(polymer[i+1]+32) {
					skipOne = true
					continue
				}
				//log.Println(i, rune(c-32))
			}
			newOne.WriteRune(c)

			//fmt.Println(i, string(c+1))
		}
		//fmt.Println(newOne, len(newOne))
		//log.Println(original, len(newOne))
		//log.Println(newOne)
		reduced = original + 0
		original = len(newOne.String())
		polymer = newOne.String()
	}
	return original

}
