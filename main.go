package main

import (
	"fmt"
	"io"
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
		ListTemplates(output)
		return 0
	}

	if len(conf.Code) == 0 {
		fmt.Fprintln(output, Usage)
	}

	if Exists(Templates, conf.Code) {
		conf.Code = Templates[conf.Code]
	}

	names, err := Generate(conf.Count, conf.Code)
	if err != nil {
		return Die(err)
	}

	if err = PrintColumns(names, output); err != nil {
		return Die(err)
	}

	return 0
}
