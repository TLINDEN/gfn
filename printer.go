package main

import (
	"fmt"
	"io"
	"sort"
)

func ListTemplates(output io.Writer) {
	names := []string{}

	for name := range Templates {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Fprintln(output, name)
	}
}

func PrintColumns(names []string, output io.Writer) error {
	count := len(names)

	// no need for the hassle to calculate columns
	if count <= Columns {
		for _, name := range names {
			fmt.Fprintln(output, name)
		}

		return nil
	}

	// get a transposed list of columns
	padlist, max := Getcolumns(names, Columns)

	// make sure there's enough spacing between the columns
	format := fmt.Sprintf("%%-%ds", max+1)

	for _, row := range padlist {
		for _, word := range row {
			fmt.Fprintf(output, format, word)
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
