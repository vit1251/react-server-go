package react

import (
	"fmt"
	"strings"
)

type Attrs struct {
	attrs []*Attr
}

type Attr struct {
	name 	string
	value	string
}

func NewAttr(name string, value string) *Attr {
	return &Attr{
		name: name,
		value: value,
	}
}

func NewAttrs(attrs ...*Attr) *Attrs {
	result := &Attrs {
	    attrs: attrs,
	}
	return result
}

func (self Attrs) String() string {
	var rows []string
	for _, prop := range self.attrs {
		if prop != nil {
			row := fmt.Sprintf("%s=\"%s\"", prop.name, prop.value) // TODO - escape ...
			rows = append(rows, row)
		}
	}
	return strings.Join(rows, " ")
}
