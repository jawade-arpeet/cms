package issue

import (
	"fmt"
	"slices"
	"strings"
)

type Severity string

const (
	SevertityError   Severity = "error"
	SevertityWarning Severity = "warning"
)

type Issue struct {
	Key      string
	Values   []string
	Indexes  []int
	Severity Severity
}

func New(key string, severity Severity) *Issue {
	return &Issue{
		Key:      key,
		Values:   []string{},
		Indexes:  []int{},
		Severity: severity,
	}
}

func (i *Issue) Error() string {
	return fmt.Sprintf(
		"found %d invalid %s values %s at indexes %v",
		len(i.Indexes),
		i.Key,
		strings.Join(i.Values, ", "),
		i.Indexes,
	)
}

func (i *Issue) Add(idx int, val string) {
	if !slices.Contains(i.Indexes, idx) {
		i.Indexes = append(i.Indexes, idx)
	}

	if !slices.Contains(i.Values, val) {
		i.Values = append(i.Values, val)
	}
}
