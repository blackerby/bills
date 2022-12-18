package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/gen2brain/go-fitz"
)

var header = regexp.MustCompile(`(?m)(((?:S|H)\.? ?(?:R\.?)? (?:J\.? Res\. ?)?)(\w{1,5})\.? ((?:M(?:rs?|essrs)\.) .+?)(?:[;,:])? (\w{1,9} \d{1,2}[.,] \d{4})[.—]? ?\n?(?:\((['0-9a-zA-Z ]+)\))?(?:\.|.+\.|:|.+:)?)`)

func main() {
	inputFile := flag.String("file", "", "PDF file to extract text from")
	pageNumber := flag.Int("page", 1, "Page number of PDF file to extract text from")
	flag.Parse()

	if *inputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*inputFile, *pageNumber, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(inputFile string, pageNumber int, out io.Writer) error {

	doc, err := fitz.New(inputFile)
	if err != nil {
		return err
	}
	defer doc.Close()

	text, err := doc.Text(pageNumber)
	if err != nil {
		return err
	}

	headers := headers(text)
	summaries := summaries(text)
	records := getMetadata(headers)
	fileNameData := getFilenameData(records[1:])

	mdOutFile := mdOutFileName(inputFile, pageNumber)
	f, err := os.Create(mdOutFile)
	if err != nil {
		return err
	}
	writeMetadata(records, f)

	writeSummaries(os.Stdout, inputFile, summaries, fileNameData)
	return nil
}
