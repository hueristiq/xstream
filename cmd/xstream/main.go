package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hueristiq/hqgolog"
	"github.com/hueristiq/xstream/internal/configuration"
	"github.com/spf13/pflag"
)

var (
	soak           bool
	trim           bool
	uniqueOutput   bool
	appendToOutput bool
	quiet          bool
	preview        bool
)

func init() {
	pflag.BoolVar(&soak, "soak", false, "")
	pflag.BoolVar(&trim, "trim", false, "")
	pflag.BoolVarP(&uniqueOutput, "unique", "u", false, "")
	pflag.BoolVarP(&appendToOutput, "append", "a", false, "")
	pflag.BoolVarP(&quiet, "quiet", "q", false, "")
	pflag.BoolVarP(&preview, "preview", "p", false, "")

	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, configuration.BANNER)

		h := "\nUSAGE:\n"
		h += fmt.Sprintf(" %s [OPTION]... [FILE]\n", configuration.NAME)

		h += "\nINPUT OPTIONS:\n"
		h += "     --soak bool        soak up all input before writing to file\n"

		h += "\nMANIPULATION:\n"
		h += "     --trim bool        trim leading and trailing whitespace\n"

		h += "\nOUTPUT OPTIONS:\n"
		h += " -u, --unique bool      output unique lines\n"
		h += " -a, --append bool      append lines to output\n"
		h += " -q, --quiet bool       suppress output to stdout\n"
		h += " -p, --preview bool     preview new lines, without writing to file\n"

		fmt.Fprintln(os.Stderr, h)
	}

	pflag.Parse()
}

func main() {
	destination := pflag.Arg(0)

	var err error

	var writer io.WriteCloser

	uniqueDestinationLinesMap := map[string]bool{}

	if destination != "" && uniqueOutput && appendToOutput {
		uniqueDestinationLinesMap, err = readFileIntoMap(destination, trim)
		if err != nil && !os.IsNotExist(err) {
			hqgolog.Fatal().Msg(err.Error())
		}
	}

	if destination != "" && !preview {
		writer, err = getWriteCloser(destination, appendToOutput)
		if err != nil {
			hqgolog.Fatal().Msg(err.Error())
		}

		defer writer.Close()
	}

	if soak {
		if err = processInputInSoakMode(uniqueDestinationLinesMap, destination, writer); err != nil {
			hqgolog.Fatal().Msg(err.Error())
		}
	} else {
		if err = processInputInDefaultMode(uniqueDestinationLinesMap, destination, writer); err != nil {
			hqgolog.Fatal().Msg(err.Error())
		}
	}
}

func processInputInSoakMode(uniqueDestinationLinesMap map[string]bool, destination string, df io.WriteCloser) (err error) {
	var inputLinesSlice []string

	inputLinesSlice, err = readStdinIntoSlice(trim)
	if err != nil {
		return
	}

	for _, line := range inputLinesSlice {
		if uniqueOutput {
			if uniqueDestinationLinesMap[line] {
				continue
			}

			uniqueDestinationLinesMap[line] = true
		}

		if !quiet {
			fmt.Println(line)
		}

		if !preview && destination != "" {
			fmt.Fprintf(df, "%s\n", line)
		}
	}

	return
}

func processInputInDefaultMode(uniqueDestinationLinesMap map[string]bool, destination string, df io.WriteCloser) (err error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if trim {
			line = strings.TrimSpace(line)
		}

		if uniqueOutput {
			if uniqueDestinationLinesMap[line] {
				continue
			}

			uniqueDestinationLinesMap[line] = true
		}

		if !quiet {
			fmt.Println(line)
		}

		if !preview && destination != "" {
			fmt.Fprintf(df, "%s\n", line)
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func readFileIntoMap(file string, trim bool) (lines map[string]bool, err error) {
	lines = map[string]bool{}

	var f *os.File

	f, err = os.Open(file)
	if err != nil {
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if trim {
			line = strings.TrimSpace(line)
		}

		if _, ok := lines[line]; ok {
			continue
		}

		lines[line] = true
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func getWriteCloser(file string, appendToFile bool) (writer io.WriteCloser, err error) {
	directory := filepath.Dir(file)

	if _, err = os.Stat(directory); os.IsNotExist(err) {
		if err = os.MkdirAll(directory, os.ModePerm); err != nil {
			return
		}
	}

	if appendToFile {
		writer, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	} else {
		writer, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	}

	return
}

func readStdinIntoSlice(trim bool) (lines []string, err error) {
	lines = []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if trim {
			line = strings.TrimSpace(line)
		}

		lines = append(lines, line)
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
