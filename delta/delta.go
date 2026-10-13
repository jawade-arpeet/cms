package delta

import (
	"cms/validation/restriction"
)

type Delta[T any] struct {
	Inserted []T
	Updated  []T
	Deleted  []T
}

type ImportResDeltaRow struct {
	restriction.ImportResExcelRow
	IsActive     bool `xml:"is_active"`
	IsRestricted bool `xml:"is_restricted"`
}

func newImportResDeltaRow(row restriction.ImportResExcelRow) ImportResDeltaRow {
	return ImportResDeltaRow{
		ImportResExcelRow: row,
		IsActive:          true,
		IsRestricted:      true,
	}
}

func GenerateResDelta(
	countryISO2 string,
	contentType string,
	records []restriction.ImportResExcelRow,
) *Delta[ImportResDeltaRow] {
	inserted := make([]ImportResDeltaRow, 0, len(records))
	updated := make([]ImportResDeltaRow, 0, len(records))
	deleted := make([]ImportResDeltaRow, 0, len(records))

	for _, rec := range records {
		deltaRow := newImportResDeltaRow(rec)

		switch deltaRow.Status {
		case "I":
			inserted = append(inserted, deltaRow)
		case "U":
			updated = append(updated, deltaRow)
		case "D":
			deleted = append(deleted, deltaRow)
		}
	}

	return &Delta[ImportResDeltaRow]{Inserted: inserted, Updated: updated, Deleted: deleted}
}

func GenerateInsertedDelta() {}
func GenerateUpdatedDelta()  {}
func GenerateDeletedDelta()  {}
