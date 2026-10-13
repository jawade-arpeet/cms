package checks

import "slices"

var (
	addCvdCategories = []string{"Anti-Dumping duties", "Countervailing duties"}
)

func CheckValidAddCvdCategories(category string) bool {
	return slices.Contains(addCvdCategories, category)
}

func CheckValidControllingAuthority(
	authorities []string,
	authority string,
) bool {
	return slices.Contains(authorities, authority)
}
