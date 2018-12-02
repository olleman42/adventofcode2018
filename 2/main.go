package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	h, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	c, err := ioutil.ReadAll(h)
	if err != nil {
		log.Fatal(err)
	}

	ids := strings.Split(string(c), "\n")

	// dCount := 0
	// tCount := 0
	// for _, id := range ids {
	// 	countDubsTrips(&dCount, &tCount, id)
	// 	log.Println(dCount, tCount)
	// }

	// log.Println(dCount * tCount)

	for _, id := range ids {
		for _, idx := range ids {
			if id != idx {
				diffCount := 0
				diffIdx := 0
				for i := 0; i < len(id); i++ {
					if id[i] != idx[i] {
						diffCount++
						diffIdx = i
					}
				}
				if diffCount < 2 {

					log.Println(id[0:diffIdx] + id[diffIdx+1:len(id)])
					os.Exit(0)
				}
			}
		}
	}
}

func countDubsTrips(dCount *int, tCount *int, word string) {
	charMap := map[rune]int{}

	for _, c := range word {
		if _, ok := charMap[c]; ok {
			charMap[c]++
			continue
		}
		charMap[c] = 1
	}

	dLock := false
	tLock := false

	for _, v := range charMap {
		if v == 2 && !dLock {
			*dCount++
			dLock = true
			continue
		}
		if v == 3 && !tLock {
			*tCount++
			tLock = true
			continue
		}
	}
}
