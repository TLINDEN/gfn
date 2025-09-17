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
	"os"
	"runtime/debug"

	"log/slog"

	"github.com/tlinden/yadu"
)

func main() {
	os.Exit(Main(os.Stdout))
}

func TMain() int {
	return Main(os.Stdout)
}

func Main(output io.Writer) int {
	// parse config file and command line parameters, if any
	conf, err := InitConfig(output)
	if err != nil {
		return Die(err)
	}

	if conf.Showversion {
		_, err := fmt.Fprintf(output, "This is gfn version %s\n", VERSION)
		if err != nil {
			log.Fatalf("failed to print to output: %s", err)
		}

		return 0
	}

	//  enable  debugging,  if  needed.   We  only  use  log/slog  for
	// debugging, so there's no need to configure it outside debugging
	if conf.Debug {
		logLevel := &slog.LevelVar{}
		// we're using a more verbose logger in debug mode
		buildInfo, _ := debug.ReadBuildInfo()
		opts := &yadu.Options{
			Level:     logLevel,
			AddSource: true,
		}

		logLevel.Set(slog.LevelDebug)

		handler := yadu.NewHandler(output, opts)
		debuglogger := slog.New(handler).With(
			slog.Group("program_info",
				slog.Int("pid", os.Getpid()),
				slog.String("go_version", buildInfo.GoVersion),
			),
		)
		slog.SetDefault(debuglogger)
	}

	// just show what we have
	if conf.Listshortcuts {
		ListTemplates(conf, output)
		return 0
	}

	// code argument is mandatory
	if len(conf.Code) == 0 {
		_, err := fmt.Fprintln(output, Usage)
		if err != nil {
			log.Fatalf("failed to print to output: %s", err)
		}

		return 1
	}

	// check if we can use a template, otherwise consider the argument
	// to be FN code
	if Exists(conf.Templates, conf.Code) {
		slog.Debug("Argument resolves to template code",
			"name", conf.Code, "code", conf.Templates[conf.Code])

		conf.Code = conf.Templates[conf.Code]
	}

	// all prepared, run baby run
	names, err := Generate(conf)
	if err != nil {
		return Die(err)
	}

	if err = PrintColumns(conf, names, output); err != nil {
		return Die(err)
	}

	return 0
}
