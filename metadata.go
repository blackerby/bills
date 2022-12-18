package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"
)

func getMetadata(headers []string) [][]string {
	records := [][]string{
		{"bill_type", "bill_number", "sponsor", "date", "committee"},
	}
	for _, h := range headers {
		groups := header.FindAllStringSubmatch(h, -1)
		for _, group := range groups {
			records = append(records, group[2:])
		}
	}
	return records
}

// TODO: write test
func writeMetadata(records [][]string, out io.Writer) {
	w := csv.NewWriter(out)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func mdOutFileName(inputFile string, pageNumber int) string {
	base := filepath.Base(inputFile)
	ext := filepath.Ext(inputFile)
	newBase := strings.TrimSuffix(base, ext)
	return fmt.Sprintf("%s_%03d.csv", newBase, pageNumber)
}
