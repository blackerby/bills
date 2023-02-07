package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/gen2brain/go-fitz"
)

var header = regexp.MustCompile(`(?m)(((?:S|H)\.? ?(?:R\.?)? (?:J\.? Res\. ?)?)(\w{1,5})\.? ((?:M(?:rs?|essrs)\.) .+?)(?:[;,:])? (\w{1,9} \d{1,2}[.,] \d{4})[.—]? ?\n?(?:\((['0-9a-zA-Z ]+)\))?(?:\.|.+\.|:|.+:)?)`)

func main() {
	// var wg sync.WaitGroup

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
	var wg sync.WaitGroup

	doc, err := fitz.New(inputFile)
	if err != nil {
		return err
	}
	defer doc.Close()

	// Page numbers are indexed from 0, so subtract 1 from page number entered by user.
	text, err := doc.Text(pageNumber - 1)
	if err != nil {
		return err
	}

	headers := headers(text)
	if headers == nil {
		outFile := pageOutFileName(inputFile, pageNumber)
		fmt.Printf("Header pattern not found, writing full page to %s\n", outFile)
		f, err := os.Create(outFile)
		if err != nil {
			panic(err)
		}
		f.WriteString(text)
		f.Close()
		return nil
	}
	summaries := summaries(text, headers)
	records := getMetadata(headers)
	fileNameData := getFilenameData(records[1:])

	mdOutFile := mdOutFileName(inputFile, pageNumber)
	f, err := os.Create(mdOutFile)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeMetadata(records, f)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeSummaries(os.Stdout, inputFile, summaries, fileNameData)
	}()

	wg.Wait()

	return nil
}

func pageOutFileName(inputFile string, pageNumber int) string {
	base := filepath.Base(inputFile)
	ext := filepath.Ext(inputFile)
	newBase := strings.TrimSuffix(base, ext)
	return fmt.Sprintf("%s_%03d.txt", newBase, pageNumber)
}
