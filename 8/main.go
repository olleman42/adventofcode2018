package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	d, err := getData()
	if err != nil {
		log.Fatal(err)
	}

	license := []int{}
	sd := strings.Split(string(d), " ")
	for _, i := range sd {
		num, err := strconv.ParseInt(i, 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		license = append(license, int(num))
	}

	fmt.Println(license)

	fmt.Println(parseNode(license))
	licenseNodes, _ := parseNode(license)
	fmt.Println(tallyMetadata(licenseNodes))
	fmt.Println(getNodeValues(licenseNodes))

}

type node struct {
	childCount, metadataCount int
	children                  []node
	metadata                  []int
	nodeValue                 int
}

func parseNode(input []int) (node, []int) {
	n := node{
		childCount:    input[0],
		metadataCount: input[1],
	}

	input = input[2:]

	for i := 0; i < n.childCount; i++ {
		child, remains := parseNode(input)
		n.children = append(n.children, child)
		input = remains
	}

	for i := 0; i < n.metadataCount; i++ {
		n.metadata = append(n.metadata, input[i])
	}

	input = input[n.metadataCount:]

	return n, input
}

func getNodeValues(n node) int {
	totalNodeValue := 0
	if n.childCount == 0 {
		for _, i := range n.metadata {
			totalNodeValue = totalNodeValue + i
		}
		return totalNodeValue
	}

	for _, i := range n.metadata {
		if i == 0 || i > n.childCount {
			continue
		}
		totalNodeValue = totalNodeValue + getNodeValues(n.children[i-1])
	}

	return totalNodeValue

}

func tallyMetadata(n node) int {
	metadataCount := 0
	for _, n := range n.metadata {
		metadataCount = metadataCount + n
	}
	for _, c := range n.children {
		metadataCount += tallyMetadata(c)
	}

	return metadataCount
}

func getData() ([]byte, error) {

	h, err := os.Open("./input.txt")
	if err != nil {
		return []byte{}, err
	}

	d, err := ioutil.ReadAll(h)
	if err != nil {
		return []byte{}, err
	}

	return d, nil
}
