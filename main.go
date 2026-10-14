package main

import (
	"cms/delta"
	"cms/koala"
	"os"
)

func main() {
	var rows = []delta.ImportResXMLRow{{
		ImportHSCode:        "01012100",
		ImportingCountry:    []string{"IN", "EU"},
		ExportingCountry:    []string{"US"},
		CountryOfOrigin:     []string{"US", "CN", "CO"},
		ExemptedCountry:     []string{"NP"},
		Conventions:         "Sample convention",
		RestrictionCategory: "Restricted",
		GovPartnerAgency:    "Sample agency",
		Status:              "U",
		UniqueKey:           "IN_import_restriction_001",
		SourceUniqueID:      "SRC_001",
		EffectiveAt:         "2026-01-01",
		ExpiryAt:            "2050-01-01",
		IsActive:            true,
		IsRestricted:        true,
	}, {
		ImportHSCode:        "98012100",
		ImportingCountry:    []string{"IN", "CN", "SK"},
		ExportingCountry:    []string{"US"},
		CountryOfOrigin:     []string{"US", "SK"},
		ExemptedCountry:     []string{"NP", "SK"},
		Conventions:         "Sample convention",
		RestrictionCategory: "Restricted",
		GovPartnerAgency:    "Sample agency",
		Status:              "U",
		UniqueKey:           "IN_import_restriction_001",
		SourceUniqueID:      "SRC_001",
		EffectiveAt:         "2026-01-01",
		ExpiryAt:            "2050-01-01",
		IsActive:            true,
		IsRestricted:        true,
	}}

	body, err := koala.ExportGzipXML(rows, "import_restrictions")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("output.xml.gz", body, 0644); err != nil {
		panic(err)
	}
}
