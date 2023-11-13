package day13

import (
	"aoc/utils"
	"bufio"
	"os"
	"slices"
	"strconv"
)

type ElementType byte

const (
	Value ElementType = iota
	SubPacket
)

type Element struct {
	Tag       ElementType
	Val       int
	SubPacket Packet
}

type Packet []Element

func parsePacket(data string) (Packet, error) {
	// Cut off beginning [ and ending ]
	data = data[1 : len(data)-1]
	//reader := bufio.NewReader(strings.NewReader(data))

	var packet Packet
	//for reader.Buffered() > 0 {
	for i := 0; i < len(data); i++ {
		// List element; so find the end and parse the list to a sub-packet
		if data[i] == '[' {
			start := i
			offset := i + 1
			stack := []byte{data[i]}

			for len(stack) > 0 {
				cur := data[offset]
				if cur == '[' {
					stack = append(stack, cur)
				} else if cur == ']' {
					stack = stack[:len(stack)-1]
				}

				offset++
			}

			subPacket, err := parsePacket(data[start:offset])
			if err != nil {
				return nil, err
			}
			packet = append(packet, Element{Tag: SubPacket, SubPacket: subPacket})

			i = offset - 1
		} else if data[i] == ',' { // Skip commas
			continue
		} else { // Value; so find the next comma (or end of slice), then parse int
			s := data[i:]
			idx := len(s)
			for cur := range s {
				if s[cur] == ',' {
					idx = cur
					break
				}
			}

			// Slice to just the integer and convert it
			s = s[:idx]
			val, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}

			// Add value
			packet = append(packet, Element{Tag: Value, Val: val})

			// Move up in original slice
			i += len(s) - 1
		}
	}

	return packet, nil

}

type PacketPair struct {
	P1 Packet
	P2 Packet
}

func parseInput(path string) ([]PacketPair, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var packetPairs []PacketPair
	for scanner.Scan() {
		firstPairData := scanner.Text()
		if firstPairData == "" {
			continue
		}

		if !scanner.Scan() {
			panic("Invalid data format")
		}
		secondPairData := scanner.Text()

		p1, err := parsePacket(firstPairData)
		if err != nil {
			return nil, err
		}
		p2, err := parsePacket(secondPairData)
		if err != nil {
			return nil, err
		}

		packetPairs = append(packetPairs, PacketPair{p1, p2})
	}

	return packetPairs, nil
}

func comparePackets(p1, p2 Packet) int {
	for i := range p1 {
		if len(p2) <= i {
			break
		}

		val := compareElements(p1[i], p2[i])
		if val != 0 {
			return val
		}
	}

	if len(p1) < len(p2) {
		return -1
	} else if len(p1) == len(p2) {
		return 0
	} else {
		return 1
	}
}

func compareElements(e1, e2 Element) int {
	switch e1.Tag {
	case Value:
		switch e2.Tag {
		case Value:
			if e1.Val < e2.Val {
				return -1
			} else if e1.Val == e2.Val {
				return 0
			} else {
				return 1
			}
		case SubPacket:
			return compareElements(Element{Tag: SubPacket, SubPacket: Packet{e1}}, e2)
		}
	case SubPacket:
		switch e2.Tag {
		case Value:
			return compareElements(e1, Element{Tag: SubPacket, SubPacket: Packet{e2}})
		case SubPacket:
			return comparePackets(e1.SubPacket, e2.SubPacket)
		}
	}

	panic("Invalid packets.")
}

func PartA(path string) int {
	pairs, err := parseInput(path)
	utils.CheckError(err)

	var sum int
	for i, p := range pairs {
		val := comparePackets(p.P1, p.P2)
		if val <= 0 {
			sum += (i + 1)
		}
		// fmt.Println(p.P1)
		// fmt.Println(p.P2)
		// fmt.Println()
	}

	return sum
}

func PartB(path string) int {
	pairs, err := parseInput(path)
	utils.CheckError(err)

	// Flatten list
	var packets []Packet
	for _, p := range pairs {
		packets = append(packets, p.P1, p.P2)
	}

	// Add divider packets
	d1 := Packet{Element{Tag: SubPacket, SubPacket: Packet{Element{Tag: Value, Val: 2}}}}
	d2 := Packet{Element{Tag: SubPacket, SubPacket: Packet{Element{Tag: Value, Val: 6}}}}
	packets = append(packets, d1, d2)

	// Sort
	slices.SortFunc(packets, comparePackets)

	// Find decoder key via the location of divider packets
	decoderKey := 1
	for i, p := range packets {
		if comparePackets(p, d1) == 0 || comparePackets(p, d2) == 0 {
			decoderKey *= (i + 1)
		}
	}

	return decoderKey
}
