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
	b, err := ioutil.ReadAll(h)

	adjustments := strings.Split(string(b), "\n")
	log.Println(adjustments)

	seenFreqs := make([]int64, 0)
	var tot int64
	for {

		for _, v := range adjustments {
			// log.Println(v)
			l := len(v)
			if len(v) > 0 {
				vv, err := strconv.ParseInt(v[1:l], 0, 32)
				if err != nil {
					log.Fatal(err)
				}
				switch string(v[0]) {
				case "+":
					tot = tot + vv
				case "-":
					tot = tot - vv
				}
				if containsFreq(seenFreqs, tot) {
					log.Fatal(tot)
				}
				seenFreqs = append(seenFreqs, tot)

			}
		}
	}

	log.Println(tot)
}

func containsFreq(freqs []int64, freq int64) bool {
	for _, v := range freqs {
		if v == freq {
			return true
		}
	}
	return false
}
