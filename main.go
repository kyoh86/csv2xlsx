package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tealeg/xlsx"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		sources   []string
		output    string
		delimiter string
	)

	app := kingpin.New("csv2xlsx", "Convert and combine csv files to a xlsx file")
	app.Arg("sources", "Input CSV files").Required().ExistingFilesVar(&sources)
	app.Flag("output", "Output XLSX file name").Short('o').Required().StringVar(&output)
	app.Flag("delimiter", "Delimiter of input CSV files").Short('d').Default("\t").StringVar(&delimiter)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if err := combineCSVtoXLSX(sources, output, delimiter); err != nil {
		log.Fatalln(err)
	}
}

func combineCSVtoXLSX(sources []string, output string, delimiter string) error {
	xlsxFile := xlsx.NewFile()
	for _, s := range sources {
		if err := appendCSVtoXLSX(s, xlsxFile, delimiter); err != nil {
			return err
		}
	}
	return xlsxFile.Save(output)
}

func appendCSVtoXLSX(csvPath string, xlsxFile *xlsx.File, delimiter string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open csv %s: %s", csvPath, err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	} else {
		reader.Comma = rune('\t')
	}
	sheet, err := xlsxFile.AddSheet(csvPath)
	if err != nil {
		return fmt.Errorf("failed to add sheet %s: %s", csvPath, err)
	}
	fields, err := reader.Read()
	for err == nil {
		row := sheet.AddRow()
		for _, field := range fields {
			cell := row.AddCell()
			cell.Value = field
		}
		fields, err = reader.Read()
	}
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to load row %s", err)
	}
	return nil
}
