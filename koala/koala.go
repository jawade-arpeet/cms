package koala

import (
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
