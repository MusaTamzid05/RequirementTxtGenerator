package libs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type LibrarySearcher struct {
}

func (ls *LibrarySearcher) Search(path string) {

	paths := ls.getPaths(path)

	if len(paths) == 0 {
		fmt.Println("No python file found")
		return
	}

	for _, path := range paths {
		fmt.Println(path)
	}
}

func (ls *LibrarySearcher) getPaths(path string) []string {

	paths := []string{}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
		}

		if strings.HasSuffix(path, ".py") {
			paths = append(paths, path)
		}
		return nil
	})

	return paths
}
