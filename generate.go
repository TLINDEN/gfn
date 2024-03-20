/*
Copyright © 2024 Thomas von Dein

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	fn "github.com/s0rg/fantasyname"
)

// Actual fantasy name generation
func Generate(conf *Config) ([]string, error) {
	rand.Seed(time.Now().UnixNano())

	//  we register  each generated  word to  avoid duplicates,  which
	// naturally happens every while
	reg := map[string]int{}

	// library call
	gen, err := fn.Compile(conf.Code, fn.Collapse(true))
	if err != nil {
		return nil, fmt.Errorf("could not compile FN code: %w", err)
	}

	// fetch requested number of names
	for i := 0; len(reg) < conf.Number; i++ {
		name := gen.String()

		if !Exists(reg, name) {
			reg[name] = 1

			if conf.WordWidth < len(name) {
				conf.WordWidth = len(name)
			}
		}

		// static codes  (like 'akx',  which is  no FN  code, just a
		// literal) generates just 1 item
		if i > conf.Number*2 {
			break
		}
	}

	fmt.Println("fetched names")

	slog.Debug("Generated fantasy names from code",
		"code", conf.Code, "count-names", len(reg))

	// adjust columns, if needed
	if conf.WordWidth*conf.Columns > MaxWidth {
		conf.Columns = MaxWidth / conf.WordWidth
	}

	// we just return a slice of names
	names := make([]string, len(reg))

	i := 0
	for k, _ := range reg {
		names[i] = k
		i++
	}

	return names, nil
}
