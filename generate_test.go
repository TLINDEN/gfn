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
	"testing"
)

var tests = []struct {
	name string
	want bool
	code string
}{
	{
		name: "code-ok",
		want: true,
		code: "!s(na|ha|ma|va)v",
	},
	{
		name: "code-fail",
		want: false,
		code: "!s(na|ha|ma|vav",
	},
}

var conf = &Config{
	Number: 3,
}

func TestGenerate(t *testing.T) {
	for _, tt := range tests {
		testname := fmt.Sprintf("generate-%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			conf.Code = tt.code

			names, err := Generate(conf)

			if err != nil {
				if tt.want {
					t.Errorf("Generate() returned an unexpected error: %s", err)
				}
				return
			}

			if len(names) != 3 {
				t.Errorf("Generate() returned wrong number of results\nExp: %+v\nGot: %+v\n",
					conf.Number, len(names))
				return
			}

			for idx, name := range names {
				if len(name) == 0 {
					t.Errorf("Generate() returned empty results\nIndex: %d\nGot: <%s>\n",
						idx, name)
					return
				}
			}
		})
	}

}
