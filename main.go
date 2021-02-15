package main

import (
	"flag"
	"io/ioutil"
	"regexp"
	"strings"
)

var path = flag.String("path", "./", "Path to go source file.")

func main() {
	flag.Parse()

	if !strings.HasSuffix(*path, ".go") {
		return
	}

	read, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	res := format(string(read))
	if string(read) == res {
		return
	}

	if err = ioutil.WriteFile(*path, []byte(res), 0); err != nil {
		panic(err)
	}
}

var importRegex = regexp.MustCompile(`(?m)import \(([\w\s"./_-]+)\)`)
var newRegex = regexp.MustCompile(`(?m)[\n]+`)
var tabRegex = regexp.MustCompile(`(?m)[\t]+`)

func format(s string) string {
	r := importRegex.FindAllString(s, 1)
	if len(r) == 0 {
		return s
	}

	r[0] = newRegex.ReplaceAllString(r[0], "\n")
	r[0] = tabRegex.ReplaceAllString(r[0], "\t")

	return importRegex.ReplaceAllString(s, r[0])
}
