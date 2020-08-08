package main

import (
	"flag"
	"fmt"
	"requirements_text_generator/libs"
)

func main() {

	pathPtr := flag.String("path", "", "Path to python project")
	flag.Parse()

	if *pathPtr == "" {
		fmt.Println("Usage: python_project_path")
		return
	}

	librarySearch := libs.LibrarySearcher{}
	librarySearch.Search(*pathPtr)
}
