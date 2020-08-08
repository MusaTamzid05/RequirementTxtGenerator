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
		infoSlice, err := ls.getLibsFrom(path)

		if err != nil {
			fmt.Println(err)
		}

		for _, info := range infoSlice {
			fmt.Println(info.String())
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

func (ls *LibrarySearcher) generateLibInfo(fullLibName, path, line string) LibInfo {

	libData := strings.Split(fullLibName, ".")
	return CreateLibInfo(path, libData[0], line)
}

func (ls *LibrarySearcher) getLibsFrom(path string) ([]LibInfo, error) {

	infoSlice := []LibInfo{}

	file, err := os.Open(path)

	if err != nil {
		return infoSlice, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "import") {
			data := strings.Split(line, " ")
			info := ls.generateLibInfo(data[1], path, line)
			infoSlice = append(infoSlice, info)

		} else if strings.HasPrefix(line, "from") {
			if strings.Contains(line, "import") {
				data := strings.Split(line, " ")
				info := ls.generateLibInfo(data[1], path, line)
				infoSlice = append(infoSlice, info)

			}
		}
	}

	return infoSlice, nil
}
