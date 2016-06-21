package j2yLib

import (
	"fmt"
	"io/ioutil"
	"os"
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
	default:
		fmt.Println("unknown source.")
		os.Exit(1)
	}

	return inputBytes
}
