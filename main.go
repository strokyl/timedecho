package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		minisecond, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Duration(minisecond) * time.Millisecond)
		fmt.Println(minisecond)
	}
}
