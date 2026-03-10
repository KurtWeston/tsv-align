# tsv-align

Format tab-separated values into visually aligned columns for human readability while preserving TSV structure

## Features

- Read TSV data from stdin or file and output aligned columns to stdout
- Calculate proper column widths based on content including Unicode characters
- Preserve exact TSV structure with tabs between columns
- Add configurable padding between columns for readability
- Support for truncating long values with ellipsis when max width is specified
- Optional color-coding of alternating columns for visual clarity
- Handle edge cases: empty fields, very long values, mixed character widths
- Configurable flags: --max-width, --padding, --no-color, --truncate
- Fast streaming processing for large TSV files
- Proper handling of East Asian wide characters and emoji

## How to Use

Use this project when you need to:

- Quickly solve problems related to tsv-align
- Integrate go functionality into your workflow
- Learn how go handles common patterns

## Installation

```bash
# Clone the repository
git clone https://github.com/KurtWeston/tsv-align.git
cd tsv-align

# Install dependencies
go build
```

## Usage

```bash
./main
```

## Built With

- go

## Dependencies

- `github.com/mattn/go-runewidth`
- `github.com/fatih/color`

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
