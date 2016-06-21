package main

import (
	"./src/j2y_lib"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name    = "j2y"
	app.Usage   = "convert JSON to YAML"
	app.Version = "0.0.5"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "STDOUT",
			Usage: "path of output destination file path",
		},

		cli.BoolFlag{
			Name:  "eval, e",
			Usage: "treat argument as raw data instead of file path",
		},

		cli.BoolFlag{
			Name:  "reverse, r",
			Usage: "convert YAML to JSON",
		},

		cli.BoolFlag{
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

		inputBytes := j2yLib.GetInputBytes(source, command.Args().First())

		if command.Bool("reverse") {
			outputBytes = j2yLib.YamlToJson(inputBytes, command.Bool("minify"))
		} else {
			outputBytes = j2yLib.JsonToYaml(inputBytes)
		}

		j2yLib.Output(outputBytes, command.String("output"))
	}

	app.Run(os.Args)
}
