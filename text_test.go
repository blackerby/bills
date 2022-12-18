package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHeaders(t *testing.T) {
	t.Parallel()

	want := []string{
		"S. 24. Messrs. Costigan and Wagner; January 4, 1935 (Judiciary).",
		"S. 32. Mr. Copeland; January 4, 1935 (Public Lands and Surveys).",
		"S. 33. Mr. Copeland; January 4, 1935 (Commerce).",
		"S. 70. Mr. Wheeler; January 4, 1935 (Agriculture and Forestry).",
		"S. 81. Mr. George; January 4, 1935 (Agriculture and Forestry).",
		"S. 84. Mr. George; January 4, 1935 (Judiciary).",
	}

	text, err := os.ReadFile("testdata/74_2_003.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := headers(string(text))

	if !cmp.Equal(want, got) {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestConvertDate(t *testing.T) {
	t.Parallel()

	want := "19350104"
	got := convertDate("January 4, 1935")

	if want != got {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestGetFilenameData(t *testing.T) {
	t.Parallel()

	want := [][]string{
		{"s", "24", "19350104"},
		{"s", "32", "19350104"},
		{"s", "33", "19350104"},
		{"s", "70", "19350104"},
		{"s", "81", "19350104"},
		{"s", "84", "19350104"},
	}

	records := [][]string{
		{"bill_type", "bill_number", "sponsor", "date", "committee"},
		{"S. ", "24", "Messrs. Costigan and Wagner", "January 4, 1935", "Judiciary"},
		{"S. ", "32", "Mr. Copeland", "January 4, 1935", "Public Lands and Surveys"},
		{"S. ", "33", "Mr. Copeland", "January 4, 1935", "Commerce"},
		{"S. ", "70", "Mr. Wheeler", "January 4, 1935", "Agriculture and Forestry"},
		{"S. ", "81", "Mr. George", "January 4, 1935", "Agriculture and Forestry"},
		{"S. ", "84", "Mr. George", "January 4, 1935", "Judiciary"},
	}[1:] // this is what's done in the call in main.go

	got := getFilenameData(records)

	if !cmp.Equal(want, got) {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestSOutFileName(t *testing.T) {
	t.Parallel()

	want := "74_2_s24_19350104.txt"
	got := sOutFileName(inputFile, "s", "24", "19350104")

	if want != got {
		t.Errorf("expected %q, got %q", want, got)
	}
}
