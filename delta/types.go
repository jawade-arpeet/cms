package delta

import (
	"cms/validation/restriction"
	"encoding/xml"
)

type ImportResXMLRow struct {
	XMLName             xml.Name `xml:"import_restriction"`
	ImportHSCode        string   `xml:"import_hscode"`
	ImportingCountry    StrList  `xml:"importing_country"`
	ExportingCountry    StrList  `xml:"exporting_country"`
	CountryOfOrigin     StrList  `xml:"country_of_origin"`
	ExemptedCountry     StrList  `xml:"exempted_country"`
	Conventions         string   `xml:"conventions"`
	RestrictionCategory string   `xml:"restriction_category"`
	GovPartnerAgency    string   `xml:"other_or_partner_governing_agency"`
	Status              string   `xml:"status"`
	UniqueKey           string   `xml:"unique_key"`
	SourceUniqueID      string   `xml:"source_unique_id"`
	EffectiveAt         string   `xml:"effective_at"`
	ExpiryAt            string   `xml:"expiry_at"`
	IsActive            bool     `xml:"is_active"`
	IsRestricted        bool     `xml:"is_restricted"`
}

func NewImportResDeltaRow(row restriction.ImportResExcelRow) ImportResXMLRow {
	return ImportResXMLRow{
		ImportHSCode:        row.ImportHSCode,
		ImportingCountry:    StrList(row.ImportingCountry),
		ExportingCountry:    StrList(row.ExportingCountry),
		CountryOfOrigin:     StrList(row.CountryOfOrigin),
		ExemptedCountry:     StrList(row.ExemptedCountry),
		Conventions:         row.Conventions,
		RestrictionCategory: row.RestrictionCategory,
		GovPartnerAgency:    row.GovPartnerAgency,
		Status:              row.Status,
		UniqueKey:           row.UniqueKey,
		SourceUniqueID:      row.SourceUniqueID,
		EffectiveAt:         row.EffectiveAt,
		ExpiryAt:            row.ExpiryAt,
		IsActive:            true,
		IsRestricted:        true,
	}
}
