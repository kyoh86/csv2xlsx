# csv2xlsx

Convert and combine csv files to a xlsx file

[![PkgGoDev](https://pkg.go.dev/badge/kyoh86/csv2xlsx)](https://pkg.go.dev/kyoh86/csv2xlsx)
[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/csv2xlsx)](https://goreportcard.com/report/github.com/kyoh86/csv2xlsx)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/csv2xlsx.svg)](https://codecov.io/gh/kyoh86/csv2xlsx)
[![Release](https://github.com/kyoh86/csv2xlsx/workflows/Release/badge.svg)](https://github.com/kyoh86/csv2xlsx/releases)

## Installation

```sh
go get github.com/kyoh86/csv2xlsx
```

## Usage

```sh
csv2xlsx --output=OUTPUT [--delimiter=","] <sources>...
```

### Flags

Flag             | Description
-----------------|--------------------------------------------------------------------
--help           | Show context-sensitive help (also try --help-long and --help-man)
--output=OUTPUT  | Output XLSX file name
--delimiter=","  | Delimiter of input CSV files

### Args
Arg        | Description
-----------|-----------------
<sources>  | Input CSV files

### Hint

If you want to load TSV (Tab Splited Values):

```sh
csv2xlsx --output=OUTPUT.xlsx --delimiter=$',' SOURCE.csv
```

## Thanks

[csv2xlsx](https://github.com/tealeg/csv2xlsx)

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
