package checks

import set "github.com/deckarep/golang-set/v3"

var (
	resCategories = set.NewSet("Restricted", "Prohibited")
	statuses      = set.NewSet("I", "U", "D")
	resTypes      = set.NewSet("Declaration", "Prohibition", "Authorization", "Inspection", "Documentation")
)

func CheckValidResConventions(
	conventions set.Set[string],
	convention string,
) bool {
	return conventions.ContainsOne(convention)
}

func CheckValidGovPartnerAgency(
	agencies set.Set[string],
	agency string,
) bool {
	return agencies.ContainsOne(agency)
}

func CheckValidResCategory(category string) bool {
	return resCategories.ContainsOne(category)
}
