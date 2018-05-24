package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

const (
	usage = `sepfile separates a target file by a keyword line.

Each separated file names are added a dot and a sequence number after the target file name.

Usage:

	sepfile <srcfile> <kwdline> [<outpath>]

Each arguments are:

	<srcfile>
		a target file.
	<kwdline>
		a keyword line for separating.
		if an error occurred then surround it by '"'.
	[<outpath>]
		an output file path.
		default is ".".

`
)

func main() {
	if ln := len(os.Args); ln != 3 && ln != 4 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(2)
	}
	srcfile := os.Args[1]
	if fi, err := os.Stat(srcfile); err != nil || fi.IsDir() {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	kwdline := os.Args[2]
	outpath := "."
	if len(os.Args) == 4 {
		outpath = os.Args[3]
		if fi, err := os.Stat(outpath); err != nil || !fi.IsDir() {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}

	if err := separate(srcfile, kwdline, outpath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func separate(srcfile, kwdline, outpath string) error {
	f, err := os.Open(srcfile)
	if err != nil {
		return err
	}
	defer f.Close()

	basename := filepath.Base(srcfile)
	rtncd := returnCode()

	seq := 1
	lines := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if text := scanner.Text(); text != kwdline {
			lines = append(lines, text)
		} else {
			output(outpath, outName(basename, seq), lines, rtncd)
			seq++
			lines = make([]string, 0)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if len(lines) > 0 {
		output(outpath, outName(basename, seq), lines, rtncd)
	}

	return nil
}

func outName(basename string, seq int) string {
	return basename + "." + strconv.Itoa(seq)
}

func returnCode() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else {
		return "\n"
	}
}

func output(outpath, sepfile string, lines []string, rtncd string) error {
	f, err := os.Create(filepath.Join(outpath, sepfile))
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range lines {
		if _, err := f.WriteString(line + rtncd); err != nil {
			return err
		}
	}

	return nil
}
