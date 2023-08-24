package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hueristiq/hqgolog"
	"github.com/hueristiq/xtee/internal/configuration"
	"github.com/spf13/pflag"
)

var (
	quietMode bool
	dryRun    bool
	trim      bool
)

func init() {
	pflag.BoolVarP(&quietMode, "quiet", "q", false, "")
	pflag.BoolVar(&dryRun, "dry-run", false, "")
	pflag.BoolVar(&trim, "sources", false, "")

	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, configuration.BANNER)

		h := "USAGE:\n"
		h += fmt.Sprintf("  %s [OPTIONS]\n", configuration.NAME)

		h += "\nOPTIONS:\n"
		h += " -q, --quite bool                 quiet mode (no output at all)\n"
		h += "     --dry-run bool               don't append anything to the file, just print the new lines to stdout\n"
		h += "     --trim bool                  trim leading and trailing whitespace before comparison\n"

		fmt.Fprintln(os.Stderr, h)
	}

	pflag.Parse()
}

func main() {
	fn := pflag.Arg(0)

	lines := make(map[string]bool)

	var f io.WriteCloser

	if fn != "" {
		// read the whole file into a map if it exists
		r, err := os.Open(fn)
		if err == nil {
			sc := bufio.NewScanner(r)

			for sc.Scan() {
				if trim {
					lines[strings.TrimSpace(sc.Text())] = true
				} else {
					lines[sc.Text()] = true
				}
			}

			r.Close()
		}

		if !dryRun {
			// re-open the file for appending new stuff
			f, err = os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				hqgolog.Fatal().Msg(err.Error())
			}

			defer f.Close()
		}
	}

	// read the lines, append and output them if they're new
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()

		if trim {
			line = strings.TrimSpace(line)
		}

		if lines[line] {
			continue
		}

		// add the line to the map so we don't get any duplicates from stdin
		lines[line] = true

		if !quietMode {
			fmt.Println(line)
		}

		if !dryRun {
			if fn != "" {
				fmt.Fprintf(f, "%s\n", line)
			}
		}
	}
}
