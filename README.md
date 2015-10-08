# j2y

A tiny CLI to covert JSON/YAML.

Many thanks to [ghodss/yaml](https://github.com/ghodss/yaml).

## Download

Multi platform binaries are available.

- [Releases](https://github.com/y13i/j2y/releases)

## Installation

Put the downloaded/built binary to one of your `$PATH` directories.

## Usage

```
NAME:
   j2y - convert JSON to YAML

USAGE:
   j2y [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --output, -o "STDOUT"	path of output destination file path
   --eval, -e			treat argument as raw data instead of file path
   --reverse, -r		convert YAML to JSON
   --minify, -m			output JSON in single line format
   --help, -h			show help
   --version, -v		print the version
```

### Convert JSON to YAML

Example JSON from [JSON Tutorial](http://www.w3schools.com/json/).

```
$ cat example.json
{"employees":[
    {"firstName":"John", "lastName":"Doe"},
    {"firstName":"Anna", "lastName":"Smith"},
    {"firstName":"Peter", "lastName":"Jones"}
]}
$ j2y example.json
employees:
- firstName: John
  lastName: Doe
- firstName: Anna
  lastName: Smith
- firstName: Peter
  lastName: Jones
```

#### Output to file

```
$ j2y -o example.yml example.json
$ cat example.yml
employees:
- firstName: John
  lastName: Doe
- firstName: Anna
  lastName: Smith
- firstName: Peter
  lastName: Jones
```

### Convert YAML to JSON

```
$ j2y -r example.yml
{
  "employees": [
    {
      "firstName": "John",
      "lastName": "Doe"
    },
    {
      "firstName": "Anna",
      "lastName": "Smith"
    },
    {
      "firstName": "Peter",
      "lastName": "Jones"
    }
  ]
}
```

#### Single line output

```
$ j2y -r -m example.yml
{"employees":[{"firstName":"John","lastName":"Doe"},{"firstName":"Anna","lastName":"Smith"},{"firstName":"Peter","lastName":"Jones"}]}
```

### Eval input from command argument

```
$ j2y -e '{"employees":[{"firstName":"John","lastName":"Doe"},{"firstName":"Anna","lastName":"Smith"},{"firstName":"Peter","lastName":"Jones"}]}'
employees:
- firstName: John
  lastName: Doe
- firstName: Anna
  lastName: Smith
- firstName: Peter
  lastName: Jones
```
