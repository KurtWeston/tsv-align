package aligner

import (
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

type Options struct {
	MaxWidth int
	Padding  int
	Color    bool
	Truncate bool
}

type Aligner struct {
	opts   Options
	colors []*color.Color
}

func New(opts Options) *Aligner {
	if opts.Padding < 0 {
		opts.Padding = 0
	}
	return &Aligner{
		opts: opts,
		colors: []*color.Color{
			color.New(color.FgCyan),
			color.New(color.FgGreen),
			color.New(color.FgYellow),
			color.New(color.FgMagenta),
		},
	}
}

func ParseTSVLine(line string) []string {
	return strings.Split(line, "\t")
}

func (a *Aligner) Align(rows [][]string) []string {
	if len(rows) == 0 {
		return nil
	}

	widths := a.calculateWidths(rows)
	result := make([]string, len(rows))

	for i, row := range rows {
		result[i] = a.formatRow(row, widths)
	}

	return result
}

func (a *Aligner) calculateWidths(rows [][]string) []int {
	maxCols := 0
	for _, row := range rows {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	widths := make([]int, maxCols)
	for _, row := range rows {
		for j, cell := range row {
			w := runewidth.StringWidth(cell)
			if a.opts.MaxWidth > 0 && w > a.opts.MaxWidth {
				w = a.opts.MaxWidth
			}
			if w > widths[j] {
				widths[j] = w
			}
		}
	}

	return widths
}

func (a *Aligner) formatRow(row []string, widths []int) string {
	var parts []string
	for i, cell := range row {
		formatted := a.formatCell(cell, widths[i], i)
		parts = append(parts, formatted)
	}
	return strings.Join(parts, strings.Repeat(" ", a.opts.Padding))
}

func (a *Aligner) formatCell(cell string, width int, colIndex int) string {
	cellWidth := runewidth.StringWidth(cell)

	if a.opts.MaxWidth > 0 && cellWidth > a.opts.MaxWidth {
		if a.opts.Truncate {
			cell = a.truncate(cell, a.opts.MaxWidth)
			cellWidth = a.opts.MaxWidth
		}
	}

	padding := width - cellWidth
	if padding < 0 {
		padding = 0
	}

	formatted := cell + strings.Repeat(" ", padding)

	if a.opts.Color && colIndex < len(a.colors) {
		return a.colors[colIndex%len(a.colors)].Sprint(formatted)
	}

	return formatted
}

func (a *Aligner) truncate(s string, maxWidth int) string {
	if maxWidth < 3 {
		return strings.Repeat(".", maxWidth)
	}

	width := 0
	for i, r := range s {
		rw := runewidth.RuneWidth(r)
		if width+rw > maxWidth-3 {
			return s[:i] + "..."
		}
		width += rw
	}

	return s
}
