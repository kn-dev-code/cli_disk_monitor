package file_handler

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type File struct {
	FileName string
	FileSize int64
}

func HandleFiles() {
	scanner := bufio.NewScanner(os.Stdin)
	var paths []string

	// 1. Collect up to 3 file paths
	for len(paths) < 3 {
		fmt.Print("Enter your requested file path: ")
		if !scanner.Scan() {
			break
		}
		paths = append(paths, scanner.Text())
	}

	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()

			info, err := os.Stat(p)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", p, err)
				return
			}

			f := File{
				FileName: info.Name(),
				FileSize: info.Size(),
			}
			fmt.Printf("File: %s | Size: %d bytes\n", f.FileName, f.FileSize)
		}(path)
	}

	wg.Wait()
}
