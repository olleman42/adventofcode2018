package main

import (
	"fmt"
	"log"
)

func main() {

	anotherOne()
}

func anotherOne() {
	// players := 9
	// marblepool := 27
	log.Println("starting")
	players := 427
	marblepool := 70723 * 100
	// head := node{}
	tail := node{value: 0}
	tail.previous = &tail
	tail.next = &tail
	scores := make(map[int]int)
	var currentNode *node
	for i := 1; i < marblepool; i++ {
		currentPlayer := ((i - 1) % players) + 1
		if i == 1 {
			newNode := node{value: i, previous: &tail, next: &tail}
			tail.next = &newNode
			tail.previous = &newNode
			currentNode = &newNode
			continue
		}

		if i%23 == 0 && i != 0 {
			scoreNow := i + currentNode.PrevSeven().value
			currentNode = currentNode.PrevSeven().Unlink()
			if _, ok := scores[currentPlayer]; !ok {
				scores[currentPlayer] = scoreNow
			} else {
				scores[currentPlayer] = scores[currentPlayer] + scoreNow
			}
			continue
		}

		newNode := node{value: i}
		currentNode = currentNode.next.DropIn(&newNode)

	}

	highestScore := 0
	for _, s := range scores {
		if s > highestScore {
			highestScore = s
		}
	}
	log.Println(highestScore)
}

type node struct {
	previous *node
	next     *node
	value    int
}

func (n node) PrevSeven() *node {
	return n.previous.previous.previous.previous.previous.previous.previous
}

func (n *node) Unlink() *node {
	newCurrent := n.next
	n.previous.next = n.next
	n.next.previous = n.previous
	return newCurrent
}

func (n *node) DropIn(newNode *node) *node {
	n.next.previous = newNode
	newNode.next = n.next
	newNode.previous = n
	n.next = newNode
	return newNode
}

func dumpNode(n *node) {
	fmt.Print(n.value, " ")
	if n.next.value != 0 {
		dumpNode(n.next)
		return
	}
	fmt.Print("\r\n")
}
