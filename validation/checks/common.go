package checks

import (
	"time"

	set "github.com/deckarep/golang-set/v3"
)

func CheckValidTaxableHSCode(
	taxableHS set.Set[string],
	hsCode string,
) bool {
	return taxableHS.Contains(hsCode)
}

func CheckValidEffectiveAt(
	effAt string,
	expAt string,
) bool {
	effTime, err := time.Parse("2006-01-02", effAt)
	if err != nil {
		return false
	}

	expTime, err := time.Parse("2006-01-02", expAt)
	if err != nil {
		return false
	}

	return effTime.Before(expTime)
}

func CheckDuplicateRows(seenRowKeys set.Set[string], row string) bool {
	return seenRowKeys.ContainsOne(row)
}
