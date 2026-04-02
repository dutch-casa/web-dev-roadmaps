package main

import "fmt"

// Part of speech as design
//
// The naming conventions:
//   Types      → nouns          (Student, Order)
//   Functions  → verbs          (calculate, send, parse)
//   Booleans   → predicates     (isActive, hasPermission)
//   Errors     → failure reason (ErrNotFound, ErrTimeout)
//   Collections→ plural nouns   (users, scores)
//
// Rename everything in this program to follow these conventions.
// The code should read almost like English sentences when you're done.

type data struct {
	a    string
	b    string
	c    int
	d    bool
	list []float64
}

func check(d data) bool {
	return d.c >= 18 && d.d
}

func getStuff(d data) float64 {
	t := 0.0
	for _, x := range d.list {
		t += x
	}
	if len(d.list) == 0 {
		return 0
	}
	return t / float64(len(d.list))
}

func doThing(d data) string {
	if !check(d) {
		return "DENIED"
	}
	avg := getStuff(d)
	status := "regular"
	if avg >= 90 {
		status = "honors"
	}
	return fmt.Sprintf("%s %s (%s) - avg: %.1f - %s",
		d.a, d.b, "approved", avg, status)
}

func main() {
	things := []data{
		{a: "Jordan", b: "Lee", c: 20, d: true, list: []float64{92, 88, 95}},
		{a: "Sam", b: "Park", c: 17, d: true, list: []float64{85, 79, 91}},
		{a: "Alex", b: "Chen", c: 22, d: false, list: []float64{96, 94, 98}},
		{a: "Casey", b: "Jones", c: 19, d: true, list: []float64{72, 68, 75}},
	}

	for _, t := range things {
		fmt.Println(doThing(t))
	}
}
