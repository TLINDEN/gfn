package main

import (
	"fmt"
	"math/rand"
	"time"

	fn "github.com/s0rg/fantasyname"
)

func Generate(count int, code string) ([]string, error) {
	rand.Seed(time.Now().UnixNano())
	reg := map[string]int{}

	gen, err := fn.Compile(code, fn.Collapse(true))
	if err != nil {
		return nil, fmt.Errorf("Could not compile FN code:", err)
	}

	for i := 0; i < count; i++ {
		name := gen.String()
		if !exists(reg, name) {
			reg[name] = 1
		}
	}

	names := make([]string, len(reg))

	i := 0
	for k, _ := range reg {
		names[i] = k
		i++
	}

	return names, nil
}
