package restriction

import (
	"cms/validation/checks"
	"cms/validation/issue"
	"fmt"

	set "github.com/deckarep/golang-set/v3"
)

func Validate(
	countryISO2 string,
	contentType string,
	taxableHSCodes []string,
	agencies []string,
	conventions []string,
	records []ImportResExcelRow,
) []issue.Issue {
	seenUniqueKeys := set.NewSet[string]()
	seenRowKeys := set.NewSet[string]()
	taxableHSCodeSet := set.NewSet(taxableHSCodes...)
	agencySet := set.NewSet(agencies...)
	conventionSet := set.NewSet(conventions...)

	scriptRef := fmt.Sprintf("%s_%s", countryISO2, contentType)

	duplicateRowIssue := issue.New(
		issue.SevertityError,
		"row",
		"found duplicate rows",
	)

	invalidHSCodeIssue := issue.New(
		issue.SevertityError,
		"import_hscode",
		"found invalid hs codes",
	)

	invalidImportingCountryIssue := issue.New(
		issue.SevertityError,
		"importing_country",
		"found invalid importing country",
	)

	invalidOriginCountryIssue := issue.New(
		issue.SevertityError,
		"country_of_origin",
		"found invalid country of origin",
	)

	invalidExportingCountryIssue := issue.New(
		issue.SevertityError,
		"exporting_country",
		"found invalid exporting country",
	)

	invalidExemptedCountryIssue := issue.New(
		issue.SevertityError,
		"exempted_country",
		"found invalid exempted country",
	)

	invalidConventionIssue := issue.New(
		issue.SevertityError,
		"conventions",
		"found invalid conventions",
	)

	invalidCategoryIssue := issue.New(
		issue.SevertityError,
		"restriction_category",
		"found invalid restriction category",
	)

	invalidAgencyIssue := issue.New(
		issue.SevertityError,
		"other_or_partner_governing_agency",
		"found invalid governing partner agencies",
	)

	invalidStatusIssue := issue.New(
		issue.SevertityError,
		"status",
		"found invalid status codes",
	)

	invalidUniqueKeyIssue := issue.New(
		issue.SevertityError,
		"unique_key",
		"found invalid unique key",
	)

	duplicateUniqueKeyIssue := issue.New(
		issue.SevertityError,
		"unique_key",
		"found duplicate unique values",
	)

	invalidEffectiveAtIssue := issue.New(issue.SevertityError, "effective_at", "found invalid effective at values")

	for idx, rec := range records {
		rowIdx := idx + 2
		rowKey := fmt.Sprintf("%#v", rec)

		if checks.CheckDuplicateRows(seenRowKeys, rowKey) {
			duplicateRowIssue.Add(rowIdx, "")
		} else {
			seenRowKeys.Add(rowKey)
		}

		if !checks.CheckValidTaxableHSCode(taxableHSCodeSet, rec.ImportHSCode) {
			invalidHSCodeIssue.Add(rowIdx, rec.ImportHSCode)
		}

		for _, country := range rec.ImportingCountry {
			if !checks.CheckValidISO2(country) {
				invalidImportingCountryIssue.Add(rowIdx, country)
			}
		}

		for _, country := range rec.ExportingCountry {
			if !checks.CheckValidISO2(country) {
				invalidExportingCountryIssue.Add(rowIdx, country)
			}
		}

		for _, country := range rec.CountryOfOrigin {
			if !checks.CheckValidISO2(country) {
				invalidOriginCountryIssue.Add(rowIdx, country)
			}
		}

		for _, country := range rec.ExemptedCountry {
			if !checks.CheckValidISO2(country) {
				invalidExemptedCountryIssue.Add(rowIdx, country)
			}
		}

		if !checks.CheckValidResConventions(conventionSet, rec.Conventions) {
			invalidConventionIssue.Add(rowIdx, rec.Conventions)
		}

		if !checks.CheckValidResCategory(rec.RestrictionCategory) {
			invalidCategoryIssue.Add(rowIdx, rec.RestrictionCategory)
		}

		if !checks.CheckValidGovPartnerAgency(agencySet, rec.GovPartnerAgency) {
			invalidAgencyIssue.Add(rowIdx, rec.GovPartnerAgency)
		}

		if !checks.CheckValidStatus(rec.Status) {
			invalidStatusIssue.Add(rowIdx, rec.Status)
		}

		if !checks.CheckUniqueKey(rec.UniqueKey, rec.Status, scriptRef) {
			invalidUniqueKeyIssue.Add(rowIdx, rec.UniqueKey)
		}

		if (rec.Status == "U" || rec.Status == "D") && (checks.CheckDuplicateUniqueKeys(seenUniqueKeys, rec.UniqueKey)) {
			duplicateUniqueKeyIssue.Add(rowIdx, rec.UniqueKey)
		} else {
			seenUniqueKeys.Add(rec.UniqueKey)
		}

		if !checks.CheckValidEffectiveAt(rec.EffectiveAt, rec.ExpiryAt) {
			invalidEffectiveAtIssue.Add(rowIdx, rec.UniqueKey)
		}
	}

	issues := []issue.Issue{
		*invalidHSCodeIssue,
		*invalidImportingCountryIssue,
		*invalidExportingCountryIssue,
		*invalidOriginCountryIssue,
		*invalidExemptedCountryIssue,
		*invalidConventionIssue,
		*invalidCategoryIssue,
		*invalidStatusIssue,
		*invalidUniqueKeyIssue,
		*duplicateUniqueKeyIssue,
		*duplicateRowIssue,
		*invalidAgencyIssue,
		*invalidEffectiveAtIssue,
	}

	result := make([]issue.Issue, 0, len(issues))
	for _, validationIssue := range issues {
		if validationIssue.Indexes.Cardinality() > 0 {
			result = append(result, validationIssue)
		}
	}

	return result
}
