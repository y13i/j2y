package j2yLib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetInputBytes(source, str string) []byte {
	var inputBytes []byte

	switch source {
	case "ARGV":
		inputBytes = []byte(str)
	case "FILE":
		var err error
		inputBytes, err = ioutil.ReadFile(str)

		if err != nil {
			fmt.Println("file read error.")
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
	case "STDIN":
		scanner := bufio.NewScanner(os.Stdin)

		var inputLines []string

		for scanner.Scan() {
			inputLines = append(inputLines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		inputBytes = []byte(strings.Join(inputLines, ""))
	default:
		fmt.Println("unknown source.")
		os.Exit(1)
	}

	return inputBytes
}
