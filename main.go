package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	os.Exit(Main(os.Stdout))
}

func Main(output io.Writer) int {
	conf, err := InitConfig(output)
	if err != nil {
		return Die(err)
	}

	if conf.Showversion {
		fmt.Fprintf(output, "This is gfn version %s\n", VERSION)

		return 0
	}

	if conf.Listshortcuts {
		for _, name := range Templates {
			fmt.Println(name)
		}
		return 0
	}

	// FIXME: this is a slice of Template{}, turn it into a map
	if Contains(Templates, conf.Code) {
		conf.Code = Templates[conf.Code]
	}

	names, err := Generate(conf.Count, conf.Code)

	return 0
}

func exists[K comparable, V any](m map[K]V, v K) bool {
	if _, ok := m[v]; ok {
		return true
	}
	return false
}

func Die(err error) int {
	log.Fatal("Error", err.Error())

	return 1
}

// find an item in a list, generic variant
func Contains[E comparable](s []E, v E) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}

	return false
}
