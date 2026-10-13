package checks

import "time"

var (
	defaultExpiryAt = time.Date(2050, time.January, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
)
