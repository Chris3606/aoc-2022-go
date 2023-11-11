package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Dir struct {
	Name      string
	ParentDir *Dir
	SubDirs   []*Dir
	Files     []File
}

func NewDir(name string, parent *Dir) Dir {
	return Dir{name, parent, nil, nil}
}

func (dir *Dir) AddSubDir(name string) *Dir {
	subDir := new(Dir)
	subDir.Name = name
	subDir.ParentDir = dir
	dir.SubDirs = append(dir.SubDirs, subDir)
	return subDir
}

func (dir *Dir) GetSubDir(name string) *Dir {
	for _, dir := range dir.SubDirs {
		if dir.Name == name {
			return dir
		}
	}

	return nil
}

func (dir *Dir) GetOrAddSubDir(name string) *Dir {
	subDir := dir.GetSubDir(name)
	if subDir != nil {
		return subDir
	}

	return dir.AddSubDir(name)
}

func (dir *Dir) getSize() int {
	var size int
	for _, dir := range dir.SubDirs {
		size += dir.getSize()
	}

	for _, file := range dir.Files {
		size += file.Size
	}

	return size
}

func GetDirectories(root *Dir) []*Dir {
	var dirs []*Dir

	for _, subDir := range root.SubDirs {
		dirs = append(dirs, GetDirectories(subDir)...)
	}

	dirs = append(dirs, root)

	return dirs
}

func parseInput7(input string) (Dir, error) {
	f, err := os.Open(input)
	if err != nil {
		return Dir{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	root := NewDir("/", nil)
	curDir := &root
	scanner.Scan() // Skip first line since we've set the root dir
	for scanner.Scan() {
		text := scanner.Text()
		if text[0] == '$' { // command
			cmdParts := strings.Split(text[2:], " ")
			switch cmdParts[0] {
			case "ls": // Current directory is already set
				continue
			case "cd":
				if cmdParts[1] == ".." {
					curDir = curDir.ParentDir
				} else {
					curDir = curDir.GetSubDir(cmdParts[1])
					if curDir == nil {
						panic("CD command to directory which does not exist")
					}
				}
				continue
			default:
				panic("Unsupported command")
			}
		} else { // Dir or file listing
			cmdParts := strings.Split(text, " ")
			if cmdParts[0] == "dir" { // Create dir
				curDir.AddSubDir(cmdParts[1])
			} else { // Create new file
				size, err := strconv.Atoi(cmdParts[0])
				if err != nil {
					return Dir{}, err
				}

				curDir.Files = append(curDir.Files, File{cmdParts[1], size})
			}
		}
	}

	return root, nil
}

func Day07A(input string) int {
	fsRoot, err := parseInput7(input)
	CheckError(err)

	dirs := GetDirectories(&fsRoot)

	var sum int
	for _, dir := range dirs {
		size := dir.getSize()
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func Day07B(input string) int {
	const totalSpace = 70000000
	const requiredFreeSpace = 30000000

	fsRoot, err := parseInput7(input)
	CheckError(err)

	currentFreeSpace := totalSpace - fsRoot.getSize()
	additionalSpaceNeeded := requiredFreeSpace - currentFreeSpace

	minSize := math.MaxInt

	dirs := GetDirectories(&fsRoot)
	for _, dir := range dirs {
		size := dir.getSize()
		if size >= additionalSpaceNeeded && size < minSize {
			minSize = size
		}
	}

	return minSize

}
