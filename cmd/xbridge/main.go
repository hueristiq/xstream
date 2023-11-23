package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hueristiq/hqgolog"
	"github.com/hueristiq/xbridge/internal/configuration"
	"github.com/spf13/pflag"
)

var (
	trim    bool
	preview bool
	silent  bool
)

func init() {
	pflag.BoolVarP(&preview, "preview", "p", false, "")
	pflag.BoolVarP(&silent, "silent", "s", false, "")
	pflag.BoolVar(&trim, "trim", false, "")

	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, configuration.BANNER)

		h := "\nUSAGE:\n"
		h += fmt.Sprintf(" %s [OPTIONS]\n", configuration.NAME)

		h += "\nMANIPULATION:\n"
		h += "     --trim bool        enable leading and trailing whitespace trimming\n"

		h += "\nOUTPUT:\n"
		h += " -p, --preview bool     enable preview mode i.e show new lines, without writing the changes\n"
		h += " -s, --silent bool      enable silent mode i.e suppress stdout output\n"

		fmt.Fprintln(os.Stderr, h)
	}

	pflag.Parse()
}

func main() {
	destinationFile := pflag.Arg(0)

	lines := make(map[string]bool)

	var f io.WriteCloser

	if destinationFile != "" {
		lines = readFileIntoMap(destinationFile, trim)

		if !preview {
			var err error

			f, err = getWriteCloser(destinationFile)
			if err != nil {
				hqgolog.Fatal().Msg(err.Error())
			}

			defer f.Close()
		}
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()

		if trim {
			line = strings.TrimSpace(line)
		}

		if lines[line] {
			continue
		}

		lines[line] = true

		if !silent {
			fmt.Println(line)
		}

		if !preview && destinationFile != "" {
			fmt.Fprintf(f, "%s\n", line)
		}
	}
}

func readFileIntoMap(file string, trim bool) (lines map[string]bool) {
	lines = make(map[string]bool)

	r, err := os.Open(file)
	if err != nil {
		return
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

	return
}

func getWriteCloser(fn string) (writer io.WriteCloser, err error) {
	return os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
}
