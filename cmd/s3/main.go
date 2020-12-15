package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/drsigned/s3/pkg/s3format"

	"github.com/drsigned/gos"
	"github.com/logrusorgru/aurora/v3"
)

type options struct {
	mode string
}

var (
	co options
	fo s3format.Options
)

func banner() {
	fmt.Fprintln(os.Stderr, aurora.BrightBlue(`
     _____
 ___|___ /
/ __| |_ \
\__ \___) |
|___/____/ v1.0.0
`).Bold())
}

func init() {
	flag.Usage = func() {
		banner()

		h := "USAGE:\n"
		h += "  s3 [MODE] [OPTIONS]\n"

		h += "\nMODES:\n"
		h += "  format            convert various s3 buckets format (see formats below)\n"

		h += "\nFORMAT OPTIONS (used with mode 'format'):\n"
		h += "  path              to path-style (e.g. https://s3.amazonaws.com/github-cloud)\n"
		h += "  name              to bucket name (e.g. github-cloud)\n"
		h += "  url               to s3 url (e.g. s3://github-cloud)\n"
		h += "  vhost             to virtual-hosted-style (e.g. github-cloud.s3.amazonaws.com)\n"

		fmt.Fprint(os.Stderr, h)
	}

	flag.Parse()

	co.mode = strings.ToLower(flag.Arg(0))

	if co.mode == "format" {
		fo.Format = strings.ToLower(flag.Arg(1))
	}
}

func main() {
	if !gos.HasStdin() {
		os.Exit(1)
	}

	buckets := make(chan string)

	go func() {
		defer close(buckets)

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			buckets <- scanner.Text()
		}
	}()

	switch co.mode {
	case "format":
		formatOptions, _ := s3format.ParseOptions(&fo)
		s3format.Format(buckets, formatOptions)
	default:
		log.Fatalln(errors.New("No mode specified"))
	}

}
