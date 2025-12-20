# Advent of Code 2025

Just my solutions for [Advent of Code 2025](https://adventofcode.com/2025).

Uses [Cobra](https://github.com/spf13/cobra) as CLI framework and [Go-Funk](https://github.com/thoas/go-funk) for some array processing tools.

Requirements installation:

```bash
go mod tidy
```

How to run specific puzzle (input files included in /input-files):

```bash
go run . day <day:1-25> <part:1/2> <input-file>
```

Because of TDD approach, tests are also included:

```bash
go test
```

Optionally, one can use existing Makefile.
