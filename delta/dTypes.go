package delta

import (
	"encoding/xml"
	"strings"
)

type StrList []string

func (s StrList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(s) == 0 {
		return e.EncodeElement("[]", start)
	}

	var b strings.Builder
	b.Grow(2 + len(s)*8) // rough pre-allocation

	b.WriteByte('[')

	for i, v := range s {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteByte('\'')
		b.WriteString(v)
		b.WriteByte('\'')
	}

	b.WriteByte(']')

	return e.EncodeElement(b.String(), start)
}
