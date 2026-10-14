package checks

import set "github.com/deckarep/golang-set/v3"

var (
	addCvdCategories = set.NewSet("Anti-Dumping duties", "Countervailing duties")
)

func CheckValidAddCvdCategories(category string) bool {
	return addCvdCategories.Contains(category)
}

func CheckValidControllingAuthority(
	authorities set.Set[string],
	authority string,
) bool {
	return authorities.Contains(authority)
}
