package main

import (
	"strings"
	"testing"

	"github.com/user/tsv-align/internal/aligner"
)

func TestProcessInput_Success(t *testing.T) {
	input := "a\tb\tc\n1\t2\t3\n"
	reader := strings.NewReader(input)
	opts := aligner.Options{Padding: 2, Color: false}

	err := processInput(reader, opts)
	if err != nil {
		t.Errorf("processInput() error = %v, want nil", err)
	}
}

func TestProcessInput_EmptyInput(t *testing.T) {
	reader := strings.NewReader("")
	opts := aligner.Options{Padding: 2, Color: false}

	err := processInput(reader, opts)
	if err != nil {
		t.Errorf("processInput() with empty input error = %v, want nil", err)
	}
}

func TestProcessInput_UnicodeContent(t *testing.T) {
	input := "Name\tCity\n田中\t東京\nKim\t서울\n"
	reader := strings.NewReader(input)
	opts := aligner.Options{Padding: 1, Color: false}

	err := processInput(reader, opts)
	if err != nil {
		t.Errorf("processInput() with unicode error = %v, want nil", err)
	}
}

func TestProcessInput_WithMaxWidth(t *testing.T) {
	input := "verylongcolumnname\tshort\n"
	reader := strings.NewReader(input)
	opts := aligner.Options{MaxWidth: 10, Padding: 2, Color: false, Truncate: true}

	err := processInput(reader, opts)
	if err != nil {
		t.Errorf("processInput() with max width error = %v, want nil", err)
	}
}
