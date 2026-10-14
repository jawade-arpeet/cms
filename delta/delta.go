package delta

import (
	"cms/validation/restriction"
)

type Delta[T any] struct {
	Inserted []T
	Updated  []T
	Deleted  []T
}

func GenerateResDelta(
	countryISO2 string,
	contentType string,
	records []restriction.ImportResExcelRow,
) *Delta[ImportResXMLRow] {
	inserted := make([]ImportResXMLRow, 0, len(records))
	updated := make([]ImportResXMLRow, 0, len(records))
	deleted := make([]ImportResXMLRow, 0, len(records))

	for _, rec := range records {
		deltaRow := NewImportResDeltaRow(rec)

		switch deltaRow.Status {
		case "I":
			inserted = append(inserted, deltaRow)
		case "U":
			updated = append(updated, deltaRow)
		case "D":
			deleted = append(deleted, deltaRow)
		}
	}

	return &Delta[ImportResXMLRow]{Inserted: inserted, Updated: updated, Deleted: deleted}
}
