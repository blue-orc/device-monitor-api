package reader

import (
	"bufio"
	"fmt"
	"os"
)

func Read(filepath string) (lines []string) {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Read : %s", err)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return
}
