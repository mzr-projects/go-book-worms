package main

import (
	"testing"
)

type testCase struct {
	bookWormsFile string
	want          []BookWorm
	wantErr       bool
}

var (
	handMaidsTale = Book{
		Title:  "The Handmaid's Tale",
		Author: "Margaret Atwood",
	}
	oryxAndCrake = Book{
		Title:  "Oryx and Crake",
		Author: "Margaret Atwood",
	}
	janeAery = Book{
		Title:  "Jane Eyre",
		Author: "Charlotte Brontë",
	}
	bellJar = Book{
		Title:  "The Bell Jar",
		Author: "Sylvia Plath",
	}
)

func TestLoadBookworms(t *testing.T) {
	/*
	   Here we define a map of test cases.The key is the test case name, and the value is the testCase struct.
	*/
	tests := map[string]testCase{
		"books_in_json_file": {
			bookWormsFile: "testdata/books.json",
			want: []BookWorm{
				{Name: "Fadi", Books: []Book{handMaidsTale, bellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handMaidsTale, janeAery}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookWormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookWormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookWorms(testCase.bookWormsFile)
			if err != nil && !testCase.wantErr {
				t.Fatalf("expected error: %v, got none", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatalf("expected no error, got one %v", err.Error())
			}

			if !equalBookWorms(t, got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}

func equalBookWorms(t *testing.T, got []BookWorm, want []BookWorm) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for i := range got {
		if got[i].Name != want[i].Name {
			return false
		}

		if !equalBooks(t, got[i].Books, want[i].Books) {
			return false
		}
	}

	return true
}

func equalBooks(t *testing.T, books []Book, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}
