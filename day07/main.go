package main

import (
	"AoC2022"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type directory struct {
	parent   *directory
	children map[string]*directory
	files    []file
}

func (d *directory) totalSize() int {
	var size int
	for _, file := range d.files {
		size += file.size
	}

	for _, child := range d.children {
		size += child.totalSize()
	}

	return size
}

type command interface {
	command() string
}

type cd struct {
	target string
}

func (_ cd) command() string {
	return "cd"
}

type file struct {
	size int
	name string
}

type ls struct {
	files          []file
	directoryNames []string
}

func (_ ls) command() string {
	return "ls"
}

func evalStdin(lines []string) []command {
	var commands []command
	for i := 0; i < len(lines); i++ {
		var nextCommand command
		if strings.HasPrefix(lines[i], "$ cd") {
			nextCommand = cd{lines[i][5:]}
		} else {
			var files []file
			var directoryNames []string

			for i+1 < len(lines) && !strings.HasPrefix(lines[i+1], "$") {
				parts := strings.Split(lines[i+1], " ")
				if parts[0] == "dir" {
					directoryNames = append(directoryNames, parts[1])
				} else {
					fileSize, _ := strconv.Atoi(parts[0])
					files = append(files, file{fileSize, parts[1]})
				}
				i++
			}
			nextCommand = ls{files, directoryNames}
		}
		commands = append(commands, nextCommand)
	}

	return commands
}

func getDirectories(lines []string) []*directory {
	var directories []*directory
	var cwd *directory
	commands := evalStdin(lines)

	for _, command := range commands {
		switch command.command() {
		case "cd":
			target := command.(cd).target
			switch target {
			case "..":
				cwd = cwd.parent
			case "/":
				cwd = &directory{parent: cwd}
				directories = append(directories, cwd)
			default:
				cwd = cwd.children[target]
				directories = append(directories, cwd)
			}
		case "ls":
			cwd.files = command.(ls).files
			children := make(map[string]*directory)
			for _, directoryName := range command.(ls).directoryNames {
				children[directoryName] = &directory{parent: cwd}
			}
			cwd.children = children
		}
	}
	return directories
}

func evalA(lines []string) int {
	directories := getDirectories(lines)

	var smallDirectorySizes []int
	for _, directory := range directories {
		if directory.totalSize() <= 100000 {
			smallDirectorySizes = append(smallDirectorySizes, directory.totalSize())
		}
	}

	var smallDirectoriesTotalSize int
	for _, size := range smallDirectorySizes {
		smallDirectoriesTotalSize += size
	}

	return smallDirectoriesTotalSize
}

func evalB(lines []string) int {
	directories := getDirectories(lines)
	sort.Slice(directories, func(i, j int) bool {
		return directories[i].totalSize() < directories[j].totalSize()
	})
	maxSize := directories[len(directories)-1].totalSize()
	println(maxSize)

	for _, directory := range directories {
		if directory.totalSize() > 30000000-(70000000-maxSize) {
			return directory.totalSize()
		}
	}

	return -1
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)

	resA := evalA(lines)
	resB := evalB(lines)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 7
	//debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	//filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	//eval(filenameDebug, true)
	eval(filename, false)
}
