package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const emptyFile = "(empty)"

//Node - interface for workong with Directory and File
type Node interface {
	String() string
}

//Directory - struct for saving dir data
type Directory struct {
	Name     string
	Children []Node
}

//File - struct for saving files data
type File struct {
	Name string
	Size int64
}

func (file File) String() string {
	if file.Size == 0 {
		return file.Name + " " + emptyFile
	}

	return file.Name + "(" + strconv.FormatInt(file.Size, 10) + "b)"
}

func (dir Directory) String() string {
	return dir.Name
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())

	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	nodes, err := readFolder(path, []Node{}, printFiles)

	if err != nil {
		return fmt.Errorf("Error - %v", err)
	}

	//display a structured folder
	display(out, nodes, []string{})

	return nil
}

func readFolder(path string, nodes []Node, printFiles bool) ([]Node, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	files, err := file.Readdir(0)

	if err != nil {
		return nil, err
	}

	//Sort all files
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, f := range files {
		if !(f.IsDir() || printFiles) {
			continue
		}

		var node Node

		if f.IsDir() {
			children, _ := readFolder(filepath.Join(path, f.Name()), []Node{}, printFiles)
			node = Directory{Name: f.Name(), Children: children}
		} else {
			node = File{Name: f.Name(), Size: f.Size()}
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func display(out io.Writer, nodes []Node, prfx []string) {
	if len(nodes) == 0 {
		return
	}

	fmt.Fprintf(out, "%s", strings.Join(prfx, ""))

	node := nodes[0]

	//check if current file last is
	if len(nodes) == 1 {
		fmt.Fprintf(out, "%s%s\n", "└───", node)

		//display children
		if directory, ok := node.(Directory); ok {
			display(out, directory.Children, append(prfx, "│\t"))
		}
		return
	}

	fmt.Fprintf(out, "%s%s\n", "├───", node)

	//display children
	if directory, ok := node.(Directory); ok {
		display(out, directory.Children, append(prfx, "│\t"))
	}

	display(out, nodes[1:], prfx)

}
