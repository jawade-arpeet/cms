package checks

import (
	"slices"
	"time"
)

func CheckValidTaxableHSCode(
	taxableHS []string,
	hsCode string,
) bool {
	return slices.Contains(taxableHS, hsCode)
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
