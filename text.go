package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var monthMap = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

// TODO: write test
func splitText(text string) []string {
	loc := header.FindStringIndex(text)
	if loc == nil {
		return nil
	}
	text = text[loc[0]:]
	summaries := header.Split(text, -1)[1:]
	reformatted := []string{}
	for _, summary := range summaries {
		reformatted = append(reformatted, strings.ReplaceAll(summary, "\n\n", "\n"))
	}
	return reformatted
}

func headers(text string) []string {
	return header.FindAllString(text, -1)
}

// TODO: write test
func summaries(text string, headers []string) []string {
	summaries := splitText(text)
	if len(headers) != len(summaries) {
		panic("must be same number of headers and summaries")
	}
	out := []string{}
	for i := range headers {
		out = append(out, headers[i]+summaries[i])
	}
	return out
}

// TODO: write test
func writeSummaries(out io.Writer, inputFile string, summaries []string, fileNameData [][]string) {
	for i, record := range fileNameData {
		billType := record[0]
		billNumber := record[1]
		date := record[2]
		outFileName := sOutFileName(inputFile, billType, billNumber, date)
		f, err := os.Create(outFileName)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(out, f.Name())
		f.WriteString(summaries[i])
		f.Close()
	}
}

func sOutFileName(inputFile, billType, billNumber, date string) string {
	base := filepath.Base(inputFile)
	ext := filepath.Ext(inputFile)
	newBase := strings.TrimSuffix(base, ext)
	return fmt.Sprintf("%s_%s%s_%s.txt", newBase, billType, billNumber, date)
}

func getFilenameData(records [][]string) [][]string {
	out := [][]string{}
	for _, record := range records {
		billType := clean(record[0])
		billNumber := record[1]
		date := convertDate(record[3])
		out = append(out, []string{billType, billNumber, date})
	}
	return out
}

func convertDate(date string) string {
	parts := strings.Split(date, " ")
	month := parts[0]
	day := clean(parts[1])
	year := parts[2]
	return fmt.Sprintf("%s%s%02s", year, monthMap[month], day)
}

func clean(billType string) string {
	punct := regexp.MustCompile(`[ .,;:'\"?]`)
	cleaned := punct.ReplaceAllString(billType, "")
	return strings.ToLower(cleaned)
}
