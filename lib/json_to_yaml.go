package j2yLib

import (
	"fmt"
	"github.com/ghodss/yaml"
	"os"
)

func JsonToYaml(inputBytes []byte) []byte {
	buffer, err := yaml.JSONToYAML(inputBytes)

	if err != nil {
		fmt.Println("JSON -> YAML convert error.")
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	return buffer
}
