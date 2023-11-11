package main

import (
	"fmt"
	"os"
	"strconv"
)

func printResult[T1 any, T2 any](day int, sample bool, partA T1, partB T2) {
	formatA := "Day %dA:"
	formatB := "Day %dB:"
	if sample {
		formatA = "Day %dA (sample):"
		formatB = "Day %dB (sample):"
	}
	fmt.Printf(formatA+"\n%v\n", day, partA)
	fmt.Printf("\n"+formatB+"\n%v\n", day, partB)
}

func runCode(day int, sample bool) {
	file := "inputs/day01"
	if sample {
		file += "_sample"
	}
	file += ".txt"

	switch day {
	case 1:
		printResult(day, sample, Day01A(file), Day01B(file))
	case 2:
		fmt.Printf("Day %dA:\n%v\n", day, Day02A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day02B(file))
	case 3:
		fmt.Printf("Day %dA:\n%v\n", day, Day03A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day03B(file))
	case 4:
		fmt.Printf("Day %dA:\n%v\n", day, Day04A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day04B(file))
	case 5:
		fmt.Printf("Day %dA:\n%v\n", day, Day05A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day05B(file))
	case 6:
		fmt.Printf("Day %dA:\n%v\n", day, Day06A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day06B(file))
	case 7:
		fmt.Printf("Day %dA:\n%v\n", day, Day07A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day07B(file))
	case 8:
		fmt.Printf("Day %dA:\n%v\n", day, Day08A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day08B(file))
	case 9:
		fmt.Printf("Day %dA:\n%v\n", day, Day09A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day09B(file))
	case 10:
		fmt.Printf("Day %dA:\n%v\n", day, Day10A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day10B(file))
	case 11:
		fmt.Printf("Day %dA:\n%v\n", day, Day11A(file))
		fmt.Printf("\nDay %dB:\n%v\n", day, Day11B(file))
	// case 12:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day12A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day12B(file))
	// case 13:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day13A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day13B(file))
	// case 14:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day14A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day14B(file))
	// case 15:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day15A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day15B(file))
	// case 16:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day16A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day16B(file))
	// case 17:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day17A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day17B(file))
	// case 18:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day18A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day18B(file))
	// case 19:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day19A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day19B(file))
	// case 20:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day20A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day20B(file))
	// case 21:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day21A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day21B(file))
	// case 22:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day22A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day22B(file))
	// case 23:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day23A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day23B(file))
	// case 24:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day24A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day24B(file))
	// case 25:
	// 	fmt.Printf("Day %dA:\n%v\n", day, Day25A(file))
	// 	fmt.Printf("\nDay %dB:\n%v\n", day, Day25B(file))
	default:
		panic("Unsupported day parameter.")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("Invalid usage: exe DAY_NUM [--sample | -s].")
	}
	day, err := strconv.Atoi(args[0])
	CheckError(err)

	sample := false
	if len(args) > 1 && (args[1] == "--sample" || args[1] == "-s") {
		sample = true
	}

	runCode(day, sample)

	// switch day {
	// case 1:
	// 	fmt.Printf("Day %dA: %d\n", day, Day10A("inputs/day10.txt"))
	// 	fmt.Printf("Day %dB:\n%s\n", day, Day10B("inputs/day10.txt"))
	// }
	// fmt.Printf("Day 10A: %d\n", Day10A("inputs/day10.txt"))
	// fmt.Printf("Day 10B:\n%s\n", Day10B("inputs/day10.txt"))
}
