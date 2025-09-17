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
	"io"
	"log"
	"log/slog"
	"sort"
)

// Output a list of hardcoded FN code templates
func ListTemplates(conf *Config, output io.Writer) {
	slog.Debug("Listing configured templates", "templates", conf.Templates)

	names := []string{}

	for name := range conf.Templates {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		_, err := fmt.Fprintln(output, name)
		if err != nil {
			log.Fatalf("failed to print to output: %w", err)
		}
	}
}

// Columnar output
func PrintColumns(conf *Config, names []string, output io.Writer) error {
	count := len(names)

	// no need for the hassle to calculate columns
	if count <= conf.Columns {
		for _, name := range names {
			fmt.Fprintln(output, name)
		}

		return nil
	}

	// get a transposed list of columns
	padlist, max := Getcolumns(names, conf.Columns)

	// make sure there's enough spacing between the columns
	format := fmt.Sprintf("%%-%ds", max+1)

	for _, row := range padlist {
		for _, word := range row {
			_, err := fmt.Fprintf(output, format, word)
			if err != nil {
				log.Fatalf("failed to print to output: %w", err)
			}
		}
		fmt.Fprintln(output)
	}

	return nil
}

func Getcolumns(names []string, columns int) ([][]string, int) {
	words := len(names)
	max := 0

	// we'll have a list of $columns columns
	padlist := make([][]string, columns)

	// initialize'em
	for col := 0; col < columns; col++ {
		padlist[col] = []string{}
	}

	// fill from input
	for idx := 0; idx < words; idx += columns {
		for col := 0; col < columns; col++ {
			if idx+col >= words {
				padlist[col] = append(padlist[col], "")
			} else {
				padlist[col] = append(padlist[col], names[idx+col])
				length := len(names[idx+col])
				if length > max {
					max = length
				}
			}

		}
	}

	// turn columns to rows
	return Transpose(padlist), max
}
