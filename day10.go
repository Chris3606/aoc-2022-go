package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	CyclesRemaining int
	Op              Operation
}

func (instruction Instruction) Perform(cpu *CPU) {
	instruction.Op.Perform(cpu)
}

type Operation interface {
	Perform(*CPU)
}

type Noop struct{}

func NewNoopInstruction() Instruction {
	return Instruction{1, Noop{}}
}

func (op Noop) Perform(cpu *CPU) {}

type Addx struct {
	val int
}

func NewAddxInstruction(val int) Instruction {
	return Instruction{2, Addx{val}}
}

func (op Addx) Perform(cpu *CPU) {
	cpu.Register += op.val
}

type CPU struct {
	CurrentCycle int
	Register     int
	Instructions []Instruction
}

func NewCPU(instructions []Instruction) CPU {
	return CPU{0, 1, instructions}
}

// returns the value of the register DURING the instruction
func (cpu *CPU) Tick() int {
	reg := cpu.Register

	cpu.Instructions[0].CyclesRemaining--
	if cpu.Instructions[0].CyclesRemaining == 0 {
		cpu.Instructions[0].Perform(cpu)
		cpu.Instructions = cpu.Instructions[1:]
	}

	cpu.CurrentCycle++
	return reg
}

func parseInput10(input string) ([]Instruction, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var instructions []Instruction
	for scanner.Scan() {
		text := scanner.Text()

		split := strings.Split(text, " ")
		var instruction Instruction
		switch split[0] {
		case "noop":
			instruction = NewNoopInstruction()
		case "addx":
			val, err := strconv.Atoi(split[1])
			if err != nil {
				return nil, err
			}
			instruction = NewAddxInstruction(val)
		default:
			panic("Unsupported operation")
		}

		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func Day10A(input string) int {
	instructions, err := parseInput10(input)
	CheckError(err)

	cpu := NewCPU(instructions)

	sum := 0
	for cpu.CurrentCycle <= 220 {
		duringCycleVal := cpu.Tick()
		if cpu.CurrentCycle == 20 || (cpu.CurrentCycle > 20 && (cpu.CurrentCycle-20)%40 == 0) {
			sum += (cpu.CurrentCycle * duringCycleVal)
		}
	}

	return sum
}

func Day10B(input string) string {
	instructions, err := parseInput10(input)
	CheckError(err)

	var crt bytes.Buffer

	cpu := NewCPU(instructions)

	for len(cpu.Instructions) > 0 {
		spritePos := cpu.Tick()
		drawnPixel := cpu.CurrentCycle - 1

		pixNormalized := drawnPixel % 40 // Normalize to the range of sprite positions (horizontal component only)
		if pixNormalized >= spritePos-1 && pixNormalized <= spritePos+1 {
			crt.WriteByte('#')
		} else {
			crt.WriteByte('.')
		}

		if cpu.CurrentCycle%40 == 0 {
			crt.WriteByte('\n')
		}
	}

	return crt.String()
}
