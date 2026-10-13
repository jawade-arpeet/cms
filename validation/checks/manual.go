package checks

import (
	"slices"
	"strings"
)

func CheckValidStatus(status string) bool {
	return slices.Contains(statuses, status)
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
