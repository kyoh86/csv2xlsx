# csv2xlsx

Convert and combine csv files to a xlsx file

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
