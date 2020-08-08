package main

import (
	"requirements_text_generator/libs"
)

func main() {

	librarySearch := libs.LibrarySearcher{}
	librarySearch.Search("/home/musa/git_clones/cp-vton-plus")
}
