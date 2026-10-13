package addCvd

type AddCvdExcelRow struct {
	ImportHSCode                    string   `xml:"import_hscode"`
	ImportingCountry                []string `xml:"importing_country"`
	ExportingCountry                []string `xml:"exporting_country"`
	CountryOfOrigin                 []string `xml:"country_of_origin"`
	ExemptedCountry                 []string `xml:"exempted_country"`
	Category                        string   `xml:"category"`
	Add                             bool     `xml:"add"`
	Cvd                             bool     `xml:"cvd"`
	ControllingAuthority            string   `xml:"controlling_authority"`
	ExclusionNotes                  string   `xml:"exclusion_notes"`
	Notes                           string   `xml:"notes"`
	CaseNo                          string   `xml:"case_no"`
	ReferenceDocumentOrNotification string   `xml:"reference_document_or_notification"`
	Status                          string   `xml:"status"`
	UniqueKey                       string   `xml:"unique_key"`
	SourceUniqueID                  string   `xml:"source_unique_id"`
	EffectiveAt                     string   `xml:"effective_at"`
	ExpiryAt                        string   `xml:"expiry_at"`
}
