package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	t "runtime/msan"
	"sort"
	"strings"
)

var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

	// steps := make([]step, 0)
	steps := parseSteps(d)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range letters {
		steps = append(steps, step{ID: string(l)})
	}

	stepmap := buildStepTree(steps)

	order := strings.Builder{}
	batchEmpty := false
	for !batchEmpty {
		nextBatch := []string{}
		for _, s := range stepmap {
			if s.done {
				continue
			}
			// check if deps have resolved
			ready := true
			for _, d := range s.dependents {
				if d == "" {
					continue
				}
				if !stepmap[d].done {
					ready = false
					break
				}
			}
			if ready {
				nextBatch = append(nextBatch, s.ID)
			}
		}

		sort.Strings(nextBatch)
		// log.Println(nextBatch)
		order.WriteString(nextBatch[0])
		stepmap[nextBatch[0]].done = true
		log.Println(order.String())
		if len(order.String()) == len(letters) {
			break
		}

	}
	// os.Exit(0)

	for _, t := range stepmap {
		t.done = false
		t.started = false
	}

	tick := 0
	// tasks := order.String()
	elves := []*worker{}
	for i := 0; i < 5; i++ {
		elves = append(elves, &worker{})
	}
	// finished := ""
	// done := false
	for {
		// check if any elves have free slots
		elvesBusy := false
		for _, e := range elves {
			if !e.isBusy() {

				if e.currentJob != "" {
					stepmap[e.currentJob].done = true
				}
				// is it possible to start another?
				freetask := ""
				foundfree := false
				for _, t := range stepmap {

					if t.done || t.started {
						continue
					}

					task := *t
					imfree := true
					// log.Println(task.dependents)
					for _, d := range task.dependents {
						if d == "" {
							continue
						}
						if !stepmap[d].done {
							// log.Println("Setting foundfree false on dep d", task.ID)
							imfree = false
							break
							// continue
						}
					}
					if !imfree {
						continue
					}
					freetask = task.ID
					foundfree = true
					break

				}

				if !foundfree {
					continue
				}
				stepmap[freetask].started = true
				e.currentJob = freetask
				e.timeLeft = stepmap[freetask].buildTime
				elvesBusy = true

			} else {
				elvesBusy = true
			}
		}

		if !elvesBusy {
			break
		}

		fmt.Print(tick)
		for _, e := range elves {
			fmt.Print(*e)
		}
		fmt.Println()
		tick++
		for _, e := range elves {
			if e.isBusy() {
				e.timeLeft--
			}
		}

		// log.Println("what")

		// log.Println(tick)

		// if tick > 65 {

		// 	os.Exit(0)
		// }

		// time time for everyone

		// check if everyone is done
	}
	log.Println(tick)
}

func getNextFreeTask(stepmap map[string]*step) (step, error) {
	candidates := []string
	for _, s := range stepmap {
		stp := *s
		if stp.done || stp.started {
			continue
		}

		allDepsDone := true
		for _, d := range stp.dependents {
			if d == "" {
				continue
			}
			if !stepmap[d].done {
				allDepsDone = false
				break
			}
		}

		if allDepsDone {
			candidates = append(candidates, stp.ID)
		}
	}

	if len(candidates) == 0 {
		return step{}, error.Error("no candidates")
	}

	sort.Strings(candidates)

	return *stepmap[candidates[0]], nil

}

type worker struct {
	currentJob string
	timeLeft   int
}

func (w worker) isBusy() bool { return w.timeLeft > 0 }

type step struct {
	ID, pre       string
	started, done bool
	buildTime     int
	dependents    []string
}

func hasDeps(msteps map[string]*step, letter rune) bool {
	if _, ok := msteps[string(letter)]; !ok {
		return false
	}
	return true
}

func buildStepTree(steps []step) map[string]*step {
	used := make(map[string]*step)
	for _, s := range steps {
		if _, ok := used[s.ID]; !ok {
			for _, pre := range steps {
				if pre.ID == s.ID {
					s.dependents = append(s.dependents, pre.pre)
					// steps[i] = s
				}
			}
			s.buildTime = 60 + strings.Index(letters, s.ID) + 1
			sx := s
			used[s.ID] = &sx
		}
	}
	for _, v := range used {
		log.Println(v)
	}

	return used
}

func parseSteps(data []byte) []step {
	d := strings.Split(string(data), "\r\n")
	steps := make([]step, 0)
	for _, r := range d {
		parts := strings.Split(r, " ")
		id := parts[7]
		prestep := parts[1]
		// ts := findStep(steps, id)
		// if ts != nil {
		// 	// ts.dependents = append(ts.dependents, prestep)
		// 	// continue
		// }
		steps = append(steps, step{ID: id, dependents: make([]string, 0), pre: prestep})
	}
	// log.Println(steps)
	return steps
}

func findStep(steps []step, query string) *step {
	for _, s := range steps {
		if s.ID == query {
			return &s
		}
	}
	return nil
}
