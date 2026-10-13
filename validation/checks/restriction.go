package checks

import (
	"slices"
)

var (
	resCategories = []string{"Restricted", "Prohibited"}
	statuses      = []string{"I", "U", "D"}
	resTypes      = []string{"Declaration", "Prohibition", "Authorization", "Inspection", "Documentation"}
)

func CheckValidResConventions(
	conventions []string,
	convention string,
) bool {
	return slices.Contains(conventions, convention)
}

func CheckValidGovPartnerAgency(
	agencies []string,
	agency string,
) bool {
	return slices.Contains(agencies, agency)
}

func CheckValidResCategory(category string) bool {
	return slices.Contains(resCategories, category)
}
