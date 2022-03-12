package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const moduleBeg = "module "

func main() {
	file, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, moduleBeg) {
			line = strings.TrimSpace(line[len(moduleBeg):])
			fmt.Print(line)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
