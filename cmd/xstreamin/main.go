package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hueristiq/hqgolog"
	"github.com/hueristiq/xstreamin/internal/configuration"
	"github.com/spf13/pflag"
)

var (
	quite  bool
	dryRun bool
	trim   bool
)

func init() {
	pflag.BoolVarP(&quite, "quiet", "q", false, "")
	pflag.BoolVarP(&dryRun, "dry-run", "d", false, "")
	pflag.BoolVar(&trim, "sources", false, "")

	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, configuration.BANNER)

		h := "\nUSAGE:\n"
		h += fmt.Sprintf(" %s [OPTIONS]\n", configuration.NAME)

		h += "\nOPTIONS:\n"
		h += " -d, --dry-run bool     preview lines that would be appended, without writing the changes\n"
		h += " -q, --quite bool       enable quiet mode i.e. suppress output lines\n"
		h += "     --trim bool        trim leading and trailing whitespace before comparison\n"

		fmt.Fprintln(os.Stderr, h)
	}

	pflag.Parse()
}

func main() {
	fn := pflag.Arg(0)

	lines := make(map[string]bool)

	var f io.WriteCloser

	if fn != "" {
		lines = readFileIntoMap(fn, trim)

		if !dryRun {
			var err error

			f, err = getWriteCloser(fn)
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

		if !quite {
			fmt.Println(line)
		}

		if !dryRun {
			if fn != "" {
				fmt.Fprintf(f, "%s\n", line)
			}
		}
	}
}

func readFileIntoMap(fn string, trim bool) map[string]bool {
	lines := make(map[string]bool)

	r, err := os.Open(fn)
	if err != nil {
		return lines
	}

	defer r.Close()

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if trim {
			line = strings.TrimSpace(line)
		}

		lines[line] = true
	}

	return lines
}

func getWriteCloser(fn string) (io.WriteCloser, error) {
	return os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
}
