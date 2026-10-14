package checks

import (
	"strings"

	set "github.com/deckarep/golang-set/v3"
)

func CheckValidStatus(status string) bool {
	return statuses.Contains(status)
}

func CheckUniqueKey(
	key string,
	status string,
	scriptRef string,
) bool {
	switch status {
	case "I":
		return key == ""
	case "U", "D":
		return key != "" && strings.HasPrefix(key, scriptRef)
	default:
		return true
	}
}

func CheckDuplicateUniqueKeys(seenUniqueKeys set.Set[string], key string) bool {
	return seenUniqueKeys.ContainsOne(key)
}
