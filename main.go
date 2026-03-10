package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/tsv-align/internal/aligner"
)

func main() {
	var (
		maxWidth  = flag.Int("max-width", 0, "Maximum column width (0 = unlimited)")
		padding   = flag.Int("padding", 2, "Spaces between columns")
		noColor   = flag.Bool("no-color", false, "Disable color output")
		truncate  = flag.Bool("truncate", true, "Truncate long values with ellipsis")
	)
	flag.Parse()

	var reader io.Reader
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	opts := aligner.Options{
		MaxWidth: *maxWidth,
		Padding:  *padding,
		Color:    !*noColor,
		Truncate: *truncate,
	}

	if err := processInput(reader, opts); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func processInput(reader io.Reader, opts aligner.Options) error {
	scanner := bufio.NewScanner(reader)
	var rows [][]string

	for scanner.Scan() {
		row := aligner.ParseTSVLine(scanner.Text())
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	a := aligner.New(opts)
	aligned := a.Align(rows)

	for _, line := range aligned {
		fmt.Println(line)
	}

	return nil
}
