package main

import (
	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
	"aoc/day05"
	"aoc/day06"
	"aoc/day07"
	"aoc/day08"
	"aoc/day09"
	"aoc/day10"
	"aoc/day11"
	"aoc/utils"
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
	dayStr := fmt.Sprintf("%02d", day)
	file := "inputs/day" + dayStr
	if sample {
		file += "_sample"
	}
	file += ".txt"

	switch day {
	case 1:
		printResult(day, sample, day01.PartA(file), day01.PartB(file))
	case 2:
		printResult(day, sample, day02.PartA(file), day02.PartB(file))
	case 3:
		printResult(day, sample, day03.PartA(file), day03.PartB(file))
	case 4:
		printResult(day, sample, day04.PartA(file), day04.PartB(file))
	case 5:
		printResult(day, sample, day05.PartA(file), day05.PartB(file))
	case 6:
		printResult(day, sample, day06.PartA(file), day06.PartB(file))
	case 7:
		printResult(day, sample, day07.PartA(file), day07.PartB(file))
	case 8:
		printResult(day, sample, day08.PartA(file), day08.PartB(file))
	case 9:
		printResult(day, sample, day09.PartA(file), day09.PartB(file))
	case 10:
		printResult(day, sample, day10.PartA(file), day10.PartB(file))
	case 11:
		printResult(day, sample, day11.PartA(file), day11.PartB(file))
	// case 12:
	// 	printResult(day, sample, Day12A(file), Day12B(file))
	// case 13:
	// 	printResult(day, sample, Day13A(file), Day13B(file))
	// case 14:
	// 	printResult(day, sample, Day14A(file), Day14B(file))
	// case 15:
	// 	printResult(day, sample, Day15A(file), Day15B(file))
	// case 16:
	// 	printResult(day, sample, Day16A(file), Day16B(file))
	// case 17:
	// 	printResult(day, sample, Day17A(file), Day17B(file))
	// case 18:
	// 	printResult(day, sample, Day18A(file), Day18B(file))
	// case 19:
	// 	printResult(day, sample, Day19A(file), Day19B(file))
	// case 20:
	// 	printResult(day, sample, Day20A(file), Day20B(file))
	// case 21:
	// 	printResult(day, sample, Day21A(file), Day21B(file))
	// case 22:
	// 	printResult(day, sample, Day22A(file), Day22B(file))
	// case 23:
	// 	printResult(day, sample, Day23A(file), Day23B(file))
	// case 24:
	// 	printResult(day, sample, Day24A(file), Day24B(file))
	// case 25:
	// 	printResult(day, sample, Day25A(file), Day25B(file))
	default:
		panic("Unsupported day parameter.")
	}
}

func main() {
	//dayPtr := flag.Int("n", 1, "The day to run.")
	//samplePtr := flag.Bool("s", false, "When set, runs with the sample input instead of the real input.")

	args := os.Args[1:]
	if len(args) == 0 {
		panic("Invalid usage: exe DAY_NUM [--sample | -s].")
	}
	day, err := strconv.Atoi(args[0])
	utils.CheckError(err)

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
