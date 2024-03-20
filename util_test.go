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
	"reflect"
	"testing"
)

var testexists = []struct {
	name string
	want bool
	key  int
	hash map[int]int
}{
	{
		name: "have-key",
		want: true,
		key:  1,
		hash: map[int]int{1: 4, 2: 5},
	},
	{
		name: "miss-key",
		want: false,
		key:  3,
		hash: map[int]int{1: 4, 2: 5},
	},
}

func TestExists(t *testing.T) {
	for _, tt := range testexists {
		testname := fmt.Sprintf("exists-%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			got := Exists(tt.hash, tt.key)

			if got != tt.want {
				t.Errorf("Exists() returned wrong result\nExp: %+v\nGot: %+v\n",
					tt.want, got)
			}
		})
	}

}

var testcontains = []struct {
	name string
	want bool
	key  int
	list []int
}{
	{
		name: "have-item",
		want: true,
		key:  1,
		list: []int{1, 2},
	},
	{
		name: "miss-item",
		want: false,
		key:  3,
		list: []int{1, 2},
	},
}

func TestContains(t *testing.T) {
	for _, tt := range testcontains {
		testname := fmt.Sprintf("contains-%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			got := Contains(tt.list, tt.key)

			if got != tt.want {
				t.Errorf("Contains() returned wrong result\nExp: %+v\nGot: %+v\n",
					tt.want, got)
			}
		})
	}

}

var testtranspose = []struct {
	name   string
	slice  [][]string
	expect [][]string
}{
	{
		name: "2x2-matrix",
		slice: [][]string{
			{"a1", "a2"},
			{"b1", "b2"},
		},
		expect: [][]string{
			{"a1", "b1"},
			{"a2", "b2"},
		},
	},
	{
		name: "2x3-matrix",
		slice: [][]string{
			{"a1", "a2", "a3"},
			{"b1", "b2", "b3"},
		},
		expect: [][]string{
			{"a1", "b1"},
			{"a2", "b2"},
			{"a3", "b3"},
		},
	},
}

func TestTranspose(t *testing.T) {
	for _, tt := range testtranspose {
		testname := fmt.Sprintf("transpose-%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			got := Transpose(tt.slice)

			if !reflect.DeepEqual(tt.expect, got) {
				t.Errorf("Transpose() returned wrong result\nExp: %+v\nGot: %+v\n",
					tt.expect, got)
			}
		})
	}

}
