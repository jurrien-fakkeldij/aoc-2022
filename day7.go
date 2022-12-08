package main

import (
	"fmt"
	"internal/filereader"
	"strconv"
	"strings"
)

type File struct {
	parent *Folder
	size   int
	name   string
}

type Folder struct {
	folders []*Folder
	files   []*File
	parent  *Folder
	size    int
	name    string
}

func Day7() {
	fmt.Println("============== Day7 ==============")
	Day7Part1()
	Day7Part2()
}

func Day7Part1() {
	fmt.Println("============== Part1 =============")
	instructions := filereader.ReadFile("./input/day-7/test.txt")

	mainFolder := setupFolderStructure(instructions)

	calculateSizesForFolders(mainFolder, 0)
	smallFolders := findSmallFolders(mainFolder, 100000)
	totalSum := 0
	for _, folder := range smallFolders {
		totalSum += folder.size
	}

	if totalSum != 95437 {
		fmt.Println("Test went wrong:", totalSum)
	} else {
		fmt.Println("Test passed:", totalSum)
	}

	instructions = filereader.ReadFile("./input/day-7/input.txt")
	mainFolder = setupFolderStructure(instructions)

	calculateSizesForFolders(mainFolder, 0)
	smallFolders = findSmallFolders(mainFolder, 100000)
	totalSum = 0
	for _, folder := range smallFolders {
		totalSum += folder.size
	}

	fmt.Println("Output:", totalSum)
}

func findSmallFolders(mainFolder *Folder, minimumSize int) []*Folder {
	folders := []*Folder{}
	for _, folder := range mainFolder.folders {
		folders = append(folders, findSmallFolders(folder, minimumSize)...)
	}
	if mainFolder.size <= minimumSize {
		folders = append(folders, mainFolder)
	}
	return folders
}

func calculateSizesForFolders(mainFolder *Folder, level int) {
	for _, folder := range mainFolder.folders {
		calculateSizesForFolders(folder, level+1)
		folder.parent.size += folder.size
	}
	for _, file := range mainFolder.files {
		file.parent.size += file.size
	}
}

func displayFolderStructure(mainFolder *Folder, level int) {
	preString := ""
	for i := 0; i < level; i++ {
		preString += "  "
	}
	fmt.Println(preString, "-", mainFolder.name, "( dir, size=", mainFolder.size, " )")
	for _, folder := range mainFolder.folders {
		displayFolderStructure(folder, level+1)
	}
	for _, file := range mainFolder.files {
		fmt.Println(preString, "  ", "-", file.name, "( file, size=", file.size, ")")
	}
}

func setupFolderStructure(input []string) *Folder {
	currentFolder := &Folder{}
	mainFolder := currentFolder
	var previousFolder *Folder
	for _, line := range input {
		if line[0] == '$' {
			//instruction
			if line[2:4] == "cd" {
				//fmt.Println("instruction", line[5:])
				if line[5:] == ".." {
					//fmt.Println("go up", currentFolder.parent.name)
					currentFolder = currentFolder.parent
					previousFolder = currentFolder.parent
					continue
				}
				if line[5:] != "/" {
					previousFolder = currentFolder
					for _, folder := range currentFolder.folders {
						if folder.name == line[5:] {
							currentFolder = folder
							break
						}
					}
					currentFolder.parent = previousFolder
				}
				currentFolder.name = line[5:]
			} else if line[2:4] == "ls" {
				continue
			}
		} else if strings.Contains(line, "dir") {
			//			fmt.Println("adding dir", line)
			currentFolder.folders = append(currentFolder.folders, &Folder{name: line[4:]})
		} else {
			fileData := strings.Split(line, " ")
			name := fileData[1]
			size, _ := strconv.Atoi(fileData[0])
			currentFolder.files = append(currentFolder.files, &File{name: name, size: size, parent: currentFolder})
		}
	}
	return mainFolder
}

func Day7Part2() {
	fmt.Println("============== Part2 =============")
	instructions := filereader.ReadFile("./input/day-7/test.txt")

	mainFolder := setupFolderStructure(instructions)

	calculateSizesForFolders(mainFolder, 0)
	amountToDelete := 30000000 - (70000000 - mainFolder.size)
	bigFolders := findBigFolders(mainFolder, amountToDelete)
	minimum := -1
	for _, folder := range bigFolders {
		if minimum == -1 {
			minimum = folder.size
			continue
		}

		if folder.size < minimum {
			minimum = folder.size
		}
	}

	if minimum != 24933642 {
		fmt.Println("Test went wrong:", minimum)
	} else {
		fmt.Println("Test passed:", minimum)
	}

	instructions = filereader.ReadFile("./input/day-7/input.txt")
	mainFolder = setupFolderStructure(instructions)

	calculateSizesForFolders(mainFolder, 0)
	amountToDelete = 30000000 - (70000000 - mainFolder.size)
	bigFolders = findBigFolders(mainFolder, amountToDelete)
	minimum = -1
	for _, folder := range bigFolders {
		if minimum == -1 {
			minimum = folder.size
			continue
		}

		if folder.size < minimum {
			minimum = folder.size
		}
	}
	fmt.Println("Output:", minimum)
}

func findBigFolders(mainFolder *Folder, minimumSize int) []*Folder {
	folders := []*Folder{}
	for _, folder := range mainFolder.folders {
		folders = append(folders, findBigFolders(folder, minimumSize)...)
	}
	if mainFolder.size >= minimumSize {
		folders = append(folders, mainFolder)
	}
	return folders
}
