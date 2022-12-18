package main

const (
	inputFile  = "./testdata/74_2.pdf"
	pageNumber = 1
	goldenFile = "./testdata/74_2_001.txt"
)

// simulating running `bills -file 74_2.pdf -page 1`

// func TestRun(t *testing.T) {
// 	t.Parallel()

// 	var mockStdOut bytes.Buffer

// 	if err := run(inputFile, pageNumber, &mockStdOut); err != nil {
// 		t.Fatal(err)
// 	}

// 	resultFile := strings.TrimSpace(mockStdOut.String())

// 	result, err := os.ReadFile(resultFile)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected, err := os.ReadFile(goldenFile)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !bytes.Equal(expected, result) {
// 		t.Logf("golden:\n%s\n", expected)
// 		t.Logf("result:\n%s\n", result)
// 		t.Error("Result content does not match golden file")
// 	}

// 	os.Remove(resultFile)
// }

// simulating running `bills -file 74_2.pdf -page 3 -split`
