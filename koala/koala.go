package koala

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/xuri/excelize/v2"
)

func cleanText(val string) string {
	val = strings.TrimSpace(val)
	// add code for cleaning more illegal characters
	return val
}

func ReadExcel[T any](body io.ReadCloser) ([]T, error) {
	file, err := excelize.OpenReader(body)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadExcel, err)
	}

	defer file.Close()

	sheetName := file.GetSheetName(0)

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadExcel, err)
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("%w: no rows found in the excel sheet: %s", ErrReadExcel, sheetName)
	}

	header := rows[0]

	// check for duplicate columns
	colCounts := make(map[string]int, len(header))
	duplicateCols := make([]string, 0)
	cols := make([]string, len(header))

	for idx, col := range header {
		val := strings.ToLower(strings.TrimSpace(col))

		cols[idx] = val
		colCounts[val] += 1

		if colCounts[val] == 2 {
			duplicateCols = append(duplicateCols, val)
		}
	}

	if len(duplicateCols) > 0 {
		return nil, fmt.Errorf("%w: %s", ErrDuplicateColumn, strings.Join(duplicateCols, ", "))
	}

	return nil, nil
}

// ExportGzipXML returns gzip-compressed XML under the nonempty root xmlName.
// Rows use their XML names or Go type names.
func ExportGzipXML[T any](rows []T, xmlName string) ([]byte, error) {
	if xmlName == "" {
		return nil, fmt.Errorf("xml name cannot be empty string")
	}

	var output bytes.Buffer
	compressed := gzip.NewWriter(&output)
	defer compressed.Close()
	if _, err := io.WriteString(compressed, xml.Header); err != nil {
		return nil, fmt.Errorf("failed to export XML header: %w", err)
	}

	encoder := xml.NewEncoder(compressed)
	encoder.Indent("", "    ")
	root := xml.StartElement{Name: xml.Name{Local: xmlName}}

	if err := encoder.EncodeToken(root); err != nil {
		return nil, fmt.Errorf("failed to export XML root: %w", err)
	}
	if err := encoder.Encode(rows); err != nil {
		return nil, fmt.Errorf("failed to export XML rows: %w", err)
	}
	if err := encoder.EncodeToken(root.End()); err != nil {
		return nil, fmt.Errorf("failed to export XML root: %w", err)
	}
	if err := encoder.Close(); err != nil {
		return nil, fmt.Errorf("failed to finish XML export: %w", err)
	}
	if err := compressed.Close(); err != nil {
		return nil, fmt.Errorf("failed to finish gzip export: %w", err)
	}

	return output.Bytes(), nil
}
