package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "github.com/codegangsta/cli"
  "github.com/ghodss/yaml"
  "github.com/bitly/go-simplejson"
)

func main() {
  app := cli.NewApp()

  app.Name    = "j2y"
  app.Usage   = "convert JSON to YAML"
  app.Version = "0.0.4"

  app.Flags = []cli.Flag {
    cli.StringFlag {
      Name:  "output, o",
      Value: "STDOUT",
      Usage: "path of output destination file path",
    },

    cli.BoolFlag {
      Name:  "eval, e",
      Usage: "treat argument as raw data instead of file path",
    },

    cli.BoolFlag {
      Name:  "reverse, r",
      Usage: "convert YAML to JSON",
    },

    cli.BoolFlag {
      Name:  "minify, m",
      Usage: "output JSON in single line format",
    },
  }

  app.Action = func(command *cli.Context) {
    var source      string
    var outputBytes []byte

    if command.Bool("eval") {
      source = "ARGV"
    } else {
      if command.Args().First() == "" {
        fmt.Println("`j2y --help` to view usage.")
        os.Exit(1)
      }

      source = "FILE"
    }

    inputBytes := getInputBytes(source, command.Args().First())

    if command.Bool("reverse") {
      outputBytes = yamlToJson(inputBytes, command.Bool("minify"))
    } else {
      outputBytes = jsonToYaml(inputBytes)
    }

    output(outputBytes, command.String("output"))
  }

  app.Run(os.Args)
}

func getInputBytes(source, str string) []byte {
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

func yamlToJson(inputBytes []byte, minify bool) []byte {
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

func jsonToYaml(inputBytes []byte) []byte {
  buffer, err := yaml.JSONToYAML(inputBytes)

  if err != nil {
    fmt.Println("JSON -> YAML convert error.")
    fmt.Printf("err: %v\n", err)
    os.Exit(1)
  }

  return buffer
}

func output(outputBytes []byte, dest string) {
  if dest == "STDOUT" {
    fmt.Println(string(outputBytes))
  } else {
    ioutil.WriteFile(dest, outputBytes, 0644)
  }

  return
}
