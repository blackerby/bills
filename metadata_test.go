package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMetadata(t *testing.T) {
	t.Parallel()

	want := [][]string{
		{"bill_type", "bill_number", "sponsor", "date", "committee"},
		{"S. ", "24", "Messrs. Costigan and Wagner", "January 4, 1935", "Judiciary"},
		{"S. ", "32", "Mr. Copeland", "January 4, 1935", "Public Lands and Surveys"},
		{"S. ", "33", "Mr. Copeland", "January 4, 1935", "Commerce"},
		{"S. ", "70", "Mr. Wheeler", "January 4, 1935", "Agriculture and Forestry"},
		{"S. ", "81", "Mr. George", "January 4, 1935", "Agriculture and Forestry"},
		{"S. ", "84", "Mr. George", "January 4, 1935", "Judiciary"},
	}

	text, err := os.ReadFile("testdata/74_2_003.txt")
	if err != nil {
		t.Fatal(err)
	}

	headers := headers(string(text))
	got := getMetadata(headers)

	if !cmp.Equal(want, got) {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestMdOutFileName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		page int
		want string
	}{
		{"One digit page", 3, "74_2_003.csv"},
		{"Two digit page", 12, "74_2_012.csv"},
		{"Three digit page", 121, "74_2_121.csv"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want := tc.want
			got := mdOutFileName(inputFile, tc.page)
			if want != got {
				t.Errorf("expected %q, got %q", want, got)
			}
		})
	}
}
