package issue

import (
	"fmt"
	"strings"

	set "github.com/deckarep/golang-set/v3"
)

type Severity string

const (
	SevertityError   Severity = "error"
	SevertityWarning Severity = "warning"
)

type Issue struct {
	Key      string
	Message  string
	Values   set.Set[string]
	Indexes  set.Set[int]
	Severity Severity
}

func New(severity Severity, key string, msg string) *Issue {
	return &Issue{
		Key:      key,
		Values:   set.NewSet[string](),
		Indexes:  set.NewSet[int](),
		Severity: severity,
		Message:  msg,
	}
}

func (i *Issue) Error() string {
	details := fmt.Sprintf(
		"found %d invalid %s values %s at indexes %v",
		len(i.Indexes.ToSlice()),
		i.Key,
		strings.Join(i.Values.ToSlice(), ", "),
		i.Indexes,
	)

	if i.Message != "" {
		return fmt.Sprintf("%s: %s", details, i.Message)
	}

	return details
}

func (i *Issue) Add(idx int, val string) {
	i.Indexes.Add(idx)
	i.Values.Add(val)
}
