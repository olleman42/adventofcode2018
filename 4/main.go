package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var guards = make(map[int]*guard)

func main() {
	h, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	d, err := ioutil.ReadAll(h)
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(d), "\n")
	evts, err := parseEvents(rows)
	if err != nil {
		log.Fatal(err)
	}

	// sort events
	sort.Sort(evts)

	for _, e := range evts {
		fmt.Println(e.datetime.Format("2006-01-02 15:04"), e.contents)
	}
	err = parseShiftsGuards(evts)
	if err != nil {
		log.Fatal(err)
	}

	laziest := &guard{}
	for _, g := range guards {
		if g.totalSleepTime > laziest.totalSleepTime {
			laziest = g
		}
	}
	ugh := *laziest
	log.Println(*laziest)
	log.Println(ugh.id * ugh.totalSleepTime)
	// walk through all minutes in an hour and see if it's inside any of this guard's intervals
	//minuteCounter := make([]int, 0)
	// sleepiestMinute := 0
	// sleepiestIndex := 0
	// for m := 0; m < 60; m++ {
	// 	fakeMinute := strconv.Itoa(m)

	// 	testTime, err := time.Parse("4", fakeMinute)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	c := 0

	// 	for _, i := range ugh.sleepIntervals {
	// 		if testTime.Minute() >= i.start.Minute() && testTime.Minute() < i.end.Minute() {
	// 			c++
	// 		}
	// 	}
	// 	//minuteCounter = append(minuteCounter, c)
	// 	log.Println(m, c)
	// 	if c > sleepiestMinute {
	// 		sleepiestMinute = c
	// 		sleepiestIndex = m
	// 	}
	// }
	// fmt.Println(sleepiestMinute)
	// fmt.Println(sleepiestIndex * ugh.id)
	guardsOnMinute := make([]map[int]int, 0)
	minuteLazyGuardID := 0 // id
	laziestMinute := 0
	numberOfSleepsOnThisMinute := 0
	for m := 0; m < 60; m++ {
		fakeMinute := strconv.Itoa(m)

		testTime, err := time.Parse("4", fakeMinute)
		if err != nil {
			log.Fatal(err)
		}

		minuteMap := make(map[int]int)

		for _, g := range guards {
			timesOnThisMinute := 0
			for _, i := range g.sleepIntervals {
				if testTime.Minute() >= i.start.Minute() && testTime.Minute() < i.end.Minute() {
					timesOnThisMinute++
				}
			}
			minuteMap[g.id] = timesOnThisMinute
		}
		// get the laziest of this minute and compare to cached dude
		for id, t := range minuteMap {
			if t > numberOfSleepsOnThisMinute {
				minuteLazyGuardID = id
				laziestMinute = m
				numberOfSleepsOnThisMinute = t
			}
		}

		guardsOnMinute = append(guardsOnMinute, minuteMap)

	}
	for m, merp := range guardsOnMinute {
		log.Println("Minute", m)
		for id, hits := range merp {
			log.Println(id, hits)
		}
	}
	log.Println(minuteLazyGuardID, laziestMinute)
	log.Println(minuteLazyGuardID * laziestMinute)

}

type event struct {
	datetime time.Time
	contents string
}

type sleepInterval struct {
	start, end time.Time
}

type guard struct {
	id             int
	totalSleepTime int
	sleepIntervals []sleepInterval
}

type events []*event

func (evs events) Len() int      { return len(evs) }
func (evs events) Swap(i, j int) { evs[i], evs[j] = evs[j], evs[i] }
func (evs events) Less(i, j int) bool {
	return evs[i].datetime.Before(evs[j].datetime)
}

func parseEvents(rows []string) (events, error) {
	events := make([]*event, 0)
	for _, row := range rows {
		evt := event{}
		dt, err := time.Parse("2006-01-02 15:04", row[1:17])
		if err != nil {
			return nil, err
		}

		contents := row[19 : len(row)-1]

		evt.datetime = dt
		evt.contents = contents
		events = append(events, &evt)
	}
	return events, nil
}

func parseShiftsGuards(evts events) error {

	var latestGuard *guard
	sleepTime := time.Now()
	for _, evt := range evts {
		if evt.contents[0:5] == "Guard" {
			var err error
			latestGuard, err = getGuard(evt.contents, evt.datetime)
			if err != nil {
				return nil
			}
			continue
		}
		if evt.contents == "falls asleep" {
			sleepTime = evt.datetime
			ivl := sleepInterval{start: evt.datetime}
			latestGuard.sleepIntervals = append(latestGuard.sleepIntervals, ivl)
			continue
		}
		if evt.contents == "wakes up" {
			latestGuard.sleepIntervals[len(latestGuard.sleepIntervals)-1].end = evt.datetime
			latestGuard.totalSleepTime = latestGuard.totalSleepTime + int(evt.datetime.Sub(sleepTime).Minutes())
		}
	}
	return nil
}

func getGuard(contents string, date time.Time) (*guard, error) {
	x := strings.Split(contents, " ")
	id, err := strconv.ParseInt(x[1][1:len(x[1])], 0, 0)
	if err != nil {
		return nil, err
	}
	if g, ok := guards[int(id)]; ok {
		return g, nil
	}
	g := &guard{}
	g.sleepIntervals = make([]sleepInterval, 0)
	g.id = int(id)
	guards[g.id] = g
	return g, nil

}
