package main

import (
	"bufio"
	"os"
)

type Rucksack struct {
	allItems     map[byte]bool
	compartment1 map[byte]bool
	compartment2 map[byte]bool
}

func NewRucksack() Rucksack {
	return Rucksack{map[byte]bool{}, map[byte]bool{}, map[byte]bool{}}
}

func parseInput3(input string) ([]Rucksack, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var rucks []Rucksack
	for scanner.Scan() {
		c1 := scanner.Text()[0 : len(scanner.Text())/2]
		c2 := scanner.Text()[len(c1):]
		ruck := NewRucksack()
		for i := range c1 {
			ruck.compartment1[c1[i]] = true
			ruck.allItems[c1[i]] = true
		}
		for i := range c1 {
			ruck.compartment2[c2[i]] = true
			ruck.allItems[c2[i]] = true
		}

		rucks = append(rucks, ruck)
	}

	return rucks, nil
}

func (ruck Rucksack) findSharedElem() byte {
	for k := range ruck.compartment1 {
		if ruck.compartment2[k] {
			return k
		}
	}

	panic("No shared element")
}

func findBadge(ruck1, ruck2, ruck3 Rucksack) byte {
	for k := range ruck1.allItems {
		if ruck2.allItems[k] && ruck3.allItems[k] {
			return k
		}
	}

	panic("No badge found")
}

func getPriority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return (int)(item - 'a' + 1)
	} else {
		return (int)(item-'A'+1) + 26
	}
}

func Day03A(input string) int {
	rucks, err := parseInput3(input)
	CheckError(err)

	var sum int
	for _, ruck := range rucks {
		sum += getPriority(ruck.findSharedElem())
	}

	return sum
}

func Day03B(input string) int {
	rucks, err := parseInput3(input)
	CheckError(err)

	var sum int
	for i := 0; i < len(rucks); i += 3 {
		sum += getPriority(findBadge(rucks[i], rucks[i+1], rucks[i+2]))
	}

	return sum
}
