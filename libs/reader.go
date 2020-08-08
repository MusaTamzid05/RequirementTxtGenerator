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

		if len(libMap) == 0 {
			continue
		}

		for name, lines := range libMap {
			fmt.Println("File path: ", name)
			for index, line := range lines {
				fmt.Printf("%d = %s\n", index, line)
			}
		}

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

	libLines := []string{}

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "import") {
			libLines = append(libLines, line)

		} else if strings.HasPrefix(line, "from") {
			if strings.Contains(line, "import") {
				libLines = append(libLines, line)
			}
		}
	}

	if len(libLines) > 0 {
		libMap[path] = libLines
	}

	return libMap, nil
}
