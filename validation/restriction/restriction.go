package restriction

import (
	"cms/validation/checks"
	"cms/validation/issue"
	"fmt"
)

func Validate(
	countryISO2 string,
	contentType string,
	taxableHSCodes []string,
	agencies []string,
	conventions []string,
	records []ImportResExcelRow,
) []issue.Issue {
	scriptRef := fmt.Sprintf("%s_%s", countryISO2, contentType)

	invalidHSCode := issue.New("import_hscode", issue.SevertityError)
	invalidConvention := issue.New("conventions", issue.SevertityError)
	invalidCategory := issue.New("restriction_category", issue.SevertityError)
	invalidStatus := issue.New("status", issue.SevertityError)
	invalidUniqueKey := issue.New("unique_key", issue.SevertityError)
	invalidAgency := issue.New("other_or_partner_governing_agency", issue.SevertityError)
	invalidEffectiveAt := issue.New("effective_at", issue.SevertityError)

	for idx, rec := range records {
		rowIdx := idx + 2

		if !checks.CheckValidTaxableHSCode(taxableHSCodes, rec.ImportHSCode) {
			invalidHSCode.Add(rowIdx, rec.ImportHSCode)
		}

		if !checks.CheckValidResConventions(conventions, rec.Conventions) {
			invalidConvention.Add(rowIdx, rec.Conventions)
		}

		if !checks.CheckValidResCategory(rec.RestrictionCategory) {
			invalidCategory.Add(rowIdx, rec.RestrictionCategory)
		}

		if !checks.CheckValidGovPartnerAgency(agencies, rec.GovPartnerAgency) {
			invalidAgency.Add(rowIdx, rec.GovPartnerAgency)
		}

		if !checks.CheckValidStatus(rec.Status) {
			invalidStatus.Add(rowIdx, rec.Status)
		}

		if !checks.CheckUniqueKey(rec.UniqueKey, rec.Status, scriptRef) {
			invalidUniqueKey.Add(rowIdx, rec.UniqueKey)
		}

		if !checks.CheckValidEffectiveAt(rec.EffectiveAt, rec.ExpiryAt) {
			invalidEffectiveAt.Add(rowIdx, rec.UniqueKey)
		}
	}

	issues := []issue.Issue{
		*invalidHSCode,
		*invalidConvention,
		*invalidCategory,
		*invalidStatus,
		*invalidUniqueKey,
		*invalidAgency,
		*invalidEffectiveAt,
	}

	result := make([]issue.Issue, 0, len(issues))
	for _, validationIssue := range issues {
		if len(validationIssue.Indexes) > 0 {
			result = append(result, validationIssue)
		}
	}

	return result
}
