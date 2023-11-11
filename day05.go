package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type StackMove struct {
	qty  int
	from int
	to   int
}

type Stack []byte

func parseInput5(input string) ([]Stack, []StackMove, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// separate stack data and move data
	seenNewLine := false
	var stackData []string
	var moveData []string
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			seenNewLine = true
		} else {
			if seenNewLine {
				moveData = append(moveData, scanner.Text())
			} else {
				stackData = append(stackData, scanner.Text())
			}
		}
	}

	// Process stack data
	var stacks []Stack
	stackIdData := stackData[len(stackData)-1]
	// iterate over every character that is the number identifying a stack.  This logic
	// only works up to 9 stacks, but fortunately that's all the input gives us.
	for i := 1; i < len(stackIdData); i += 4 {
		// Build the stack by traversing up the other lines and checking the corresponding chars
		var stack Stack
		for j := len(stackData) - 2; j >= 0; j-- {
			data := stackData[j]
			if data[i] == ' ' {
				break
			}
			stack = append(stack, data[i])
		}
		stacks = append(stacks, stack)
	}

	var moves []StackMove
	for _, moveTxt := range moveData {
		var move StackMove
		_, err = fmt.Sscanf(moveTxt, "move %d from %d to %d", &move.qty, &move.from, &move.to)
		if err != nil {
			return nil, nil, err
		}

		moves = append(moves, move)
	}

	return stacks, moves, nil
}

func executeStackMove9000(stacks []Stack, move StackMove) {
	for i := 0; i < move.qty; i++ {
		from := move.from - 1
		to := move.to - 1

		// Remove element from its stack
		elem := stacks[from][len(stacks[from])-1]
		stacks[from] = stacks[from][:len(stacks[from])-1]

		// Put it on the destination stack
		stacks[to] = append(stacks[to], elem)
	}
}

func executeStackMove9001(stacks []Stack, move StackMove) {
	from := move.from - 1
	to := move.to - 1

	// Grab elements that we are moving and remove them from old stack
	elems := stacks[from][len(stacks[from])-move.qty:]
	stacks[from] = stacks[from][:len(stacks[from])-len(elems)]

	// Copy them onto the new stack (preserve order)
	stacks[to] = append(stacks[to], elems...)
}

func getMessage(stacks []Stack) string {
	var buff bytes.Buffer

	for _, stack := range stacks {
		buff.WriteByte(stack[len(stack)-1])
	}

	return buff.String()
}

func Day05A(input string) string {
	stacks, moves, err := parseInput5(input)
	CheckError(err)

	for _, move := range moves {
		executeStackMove9000(stacks, move)
	}

	return getMessage(stacks)
}

func Day05B(input string) string {
	stacks, moves, err := parseInput5(input)
	CheckError(err)

	for _, move := range moves {
		executeStackMove9001(stacks, move)
	}

	return getMessage(stacks)
}
