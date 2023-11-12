package day11

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	curItems       []int
	operation      Operation
	testMod        int
	trueId         int
	falseId        int
	itemsInspected int
}

type Operation func(int) int

func NewAddValOperation(val int) Operation {
	return func(i int) int { return i + val }
}

func NewAddSelfOperation() Operation {
	return func(i int) int { return i + i }
}

func NewMultiplyValOperation(val int) Operation {
	return func(i int) int { return i * val }
}

func NewMultiplySelfOperation() Operation {
	return func(i int) int { return i * i }
}

func parseInput(path string) ([]Monkey, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	monkeysData, err := utils.ReadGroups(f, func(s string) (string, error) { return s, nil })
	if err != nil {
		return nil, err
	}

	var monkeys []Monkey
	for _, monkeyData := range monkeysData {
		monkeyData = monkeyData[1:] // First line is just the ID, which we don't need
		if len(monkeyData) != 5 {
			panic("Too much monkey business, not valid data.")
		}

		// Starting items
		itemData := monkeyData[0][len("  Starting items: "):]
		startingItems, err := utils.ReadItems(strings.NewReader(itemData), ", ", strconv.Atoi)
		if err != nil {
			return nil, err
		}

		// Operation
		monkeyData = monkeyData[1:]
		var opData, operandData string
		_, err = fmt.Sscanf(monkeyData[0], "  Operation: new = old %s %s", &opData, &operandData)
		if err != nil {
			return nil, err
		}

		var op Operation
		switch opData {
		case "+":
			if operandData == "old" {
				op = NewAddSelfOperation()
			} else {
				val, err := strconv.Atoi(operandData)
				if err != nil {
					return nil, err
				}
				op = NewAddValOperation(val)
			}
		case "*":
			if operandData == "old" {
				op = NewMultiplySelfOperation()
			} else {
				val, err := strconv.Atoi(operandData)
				if err != nil {
					return nil, err
				}
				op = NewMultiplyValOperation(val)
			}
		default:
			panic("Unsupported operator.")
		}

		// Test divisible
		monkeyData = monkeyData[1:]
		var testMod int
		fmt.Sscanf(monkeyData[0], "  Test: divisible by %d", &testMod)

		// True ID
		monkeyData = monkeyData[1:]
		var trueId int
		fmt.Sscanf(monkeyData[0], "    If true: throw to monkey %d", &trueId)

		// True ID
		monkeyData = monkeyData[1:]
		var falseId int
		fmt.Sscanf(monkeyData[0], "    If false: throw to monkey %d", &falseId)

		monkey := Monkey{curItems: startingItems, operation: op, testMod: testMod, trueId: trueId, falseId: falseId}
		monkeys = append(monkeys, monkey)
	}

	return monkeys, nil
}

func simulateMonkeyBusiness(monkeys []Monkey, worryControl func(int) int) {
	for i := range monkeys {
		monkey := &monkeys[i]

		for _, item := range monkey.curItems {
			// Increase worry level
			item = monkey.operation(item)

			// Apply worry control
			item = worryControl(item)

			// Figure out the correct monkey to throw the item to
			var id int
			if (item % monkey.testMod) == 0 {
				id = monkey.trueId
			} else {
				id = monkey.falseId
			}

			// Throw the item
			monkeys[id].curItems = append(monkeys[id].curItems, item)
		}

		// Track how many items we inspected, and clear the list
		monkey.itemsInspected += len(monkey.curItems)
		monkey.curItems = []int{}
	}
}

func monkeyBusinessLevel(monkeys []Monkey) int {
	var m1, m2 int

	for _, m := range monkeys {
		if m.itemsInspected > m1 {
			m2 = m1
			m1 = m.itemsInspected
		} else if m.itemsInspected > m2 {
			m2 = m.itemsInspected
		}
	}

	return m1 * m2
}

func PartA(path string) int {
	monkeys, err := parseInput(path)
	utils.CheckError(err)

	for i := 0; i < 20; i++ {
		simulateMonkeyBusiness(monkeys, func(i int) int { return i / 3 })
	}

	return monkeyBusinessLevel(monkeys)
}

func PartB(path string) int {
	monkeys, err := parseInput(path)
	utils.CheckError(err)

	// Find the LCM of all the test numbers
	var testValues []int
	for _, m := range monkeys {
		testValues = append(testValues, m.testMod)
	}
	lcm := utils.LCM(testValues[0], testValues[1], testValues[2:]...)

	for i := 0; i < 10000; i++ {
		simulateMonkeyBusiness(monkeys, func(i int) int { return i % lcm })
	}

	return monkeyBusinessLevel(monkeys)
}
