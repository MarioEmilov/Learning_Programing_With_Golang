package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	range1, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	range2, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	if num > range1 && num < range2 {
		fmt.Printf("%d is in the range [%d; %d)", num, range1, range2)
	} else {
		fmt.Printf("%d is not in the range [%d; %d)", num, range1, range2)
	}
}
