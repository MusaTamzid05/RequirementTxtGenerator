package libs

import (
	"bufio"
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
		libMap, err := ls.getLibsFrom(path)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(libMap)
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

func (ls *LibrarySearcher) getLibsFrom(path string) (map[string][]string, error) {

	libMap := make(map[string][]string)
	file, err := os.Open(path)

	if err != nil {
		return libMap, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	libraryNames := []string{}

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "import") {
			data := strings.Split(line, " ")
			libraryNames = append(libraryNames, data[1])

		}
	}

	if len(libraryNames) > 0 {
		names := strings.Split(path, "/")
		libMap[names[len(names)-1]] = libraryNames
	}

	return libMap, nil
}
