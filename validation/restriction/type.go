package restriction

type ImportResExcelRow struct {
	ImportHSCode        string   `xml:"import_hscode"`
	ImportingCountry    []string `xml:"importing_country"`
	ExportingCountry    []string `xml:"exporting_country"`
	CountryOfOrigin     []string `xml:"country_of_origin"`
	ExemptedCountry     []string `xml:"exempted_country"`
	Conventions         string   `xml:"conventions"`
	RestrictionCategory string   `xml:"restriction_category"`
	GovPartnerAgency    string   `xml:"other_or_partner_governing_agency"`
	Status              string   `xml:"status"`
	UniqueKey           string   `xml:"unique_key"`
	SourceUniqueID      string   `xml:"source_unique_id"`
	EffectiveAt         string   `xml:"effective_at"`
	ExpiryAt            string   `xml:"expiry_at"`
}

type ExportRestrictionExcelRow struct {
	ExportHSCode     string   `xml:"export_hscode"`
	ExportingCountry []string `xml:"exporting_country"`
	ImportingCountry []string `xml:"importing_country"`
	ExemptedCountry  []string `xml:"exempted_country"`
	CountryOfOrigin  []string `xml:"country_of_origin"`
	Status           string   `xml:"status"`
	UniqueKey        string   `xml:"unique_key"`
	SourceUniqueID   string   `xml:"source_unique_id"`
	EffectiveAt      string   `xml:"effective_at"`
	ExpiryAt         string   `xml:"expiry_at"`
}
