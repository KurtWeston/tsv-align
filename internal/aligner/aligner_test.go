package aligner

import (
	"testing"
)

func TestParseTSVLine(t *testing.T) {
	tests := []struct {
		name string
		input string
		want []string
	}{
		{"simple", "a\tb\tc", []string{"a", "b", "c"}},
		{"empty fields", "a\t\tc", []string{"a", "", "c"}},
		{"single field", "hello", []string{"hello"}},
		{"empty string", "", []string{""}},
		{"unicode", "日本語\t한국어\tEmoji😀", []string{"日本語", "한국어", "Emoji😀"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseTSVLine(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("ParseTSVLine() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ParseTSVLine()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestAligner_Align(t *testing.T) {
	tests := []struct {
		name string
		opts Options
		rows [][]string
		wantLen int
	}{
		{
			name: "basic alignment",
			opts: Options{Padding: 2, Color: false},
			rows: [][]string{{"a", "b"}, {"aaa", "bb"}},
			wantLen: 2,
		},
		{
			name: "empty rows",
			opts: Options{Padding: 2, Color: false},
			rows: [][]string{},
			wantLen: 0,
		},
		{
			name: "unicode width",
			opts: Options{Padding: 1, Color: false},
			rows: [][]string{{"日本", "a"}, {"x", "test"}},
			wantLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New(tt.opts)
			got := a.Align(tt.rows)
			if len(got) != tt.wantLen {
				t.Errorf("Align() returned %d lines, want %d", len(got), tt.wantLen)
			}
		})
	}
}

func TestAligner_Truncate(t *testing.T) {
	tests := []struct {
		name string
		input string
		maxWidth int
		wantSuffix string
	}{
		{"long string", "verylongstring", 8, "..."},
		{"short string", "short", 10, ""},
		{"exact fit", "exact", 5, ""},
		{"tiny width", "test", 2, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New(Options{})
			got := a.truncate(tt.input, tt.maxWidth)
			if len(got) > tt.maxWidth {
				t.Errorf("truncate() length = %d, exceeds maxWidth %d", len(got), tt.maxWidth)
			}
			if tt.wantSuffix != "" && got[len(got)-3:] != tt.wantSuffix {
				t.Errorf("truncate() = %v, want suffix %v", got, tt.wantSuffix)
			}
		})
	}
}

func TestAligner_MaxWidth(t *testing.T) {
	opts := Options{MaxWidth: 5, Padding: 1, Color: false, Truncate: true}
	a := New(opts)
	rows := [][]string{{"verylongtext", "short"}}
	result := a.Align(rows)

	if len(result) != 1 {
		t.Fatalf("Expected 1 row, got %d", len(result))
	}
}

func TestNew_NegativePadding(t *testing.T) {
	opts := Options{Padding: -5}
	a := New(opts)
	if a.opts.Padding != 0 {
		t.Errorf("New() with negative padding = %d, want 0", a.opts.Padding)
	}
}
