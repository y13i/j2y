package j2yLib

import (
  "os"
  "fmt"
  "github.com/ghodss/yaml"
  "github.com/bitly/go-simplejson"
)

func YamlToJson(inputBytes []byte, minify bool) []byte {
  buffer, err := yaml.YAMLToJSON(inputBytes)

  if err != nil {
    fmt.Println("YAML -> JSON convert error.")
    fmt.Printf("err: %v\n", err)
    os.Exit(1)
  }

  if minify {
    return buffer
  } else {
    jsonData, err := simplejson.NewJson(buffer)

    if err != nil {
      fmt.Println("JSON parse error.")
      fmt.Printf("err: %v\n", err)
      os.Exit(1)
    }

    prettyBytes, err := jsonData.EncodePretty()

    if err != nil {
      fmt.Println("JSON encode error.")
      fmt.Printf("err: %v\n", err)
      os.Exit(1)
    }

    return prettyBytes
  }
}
