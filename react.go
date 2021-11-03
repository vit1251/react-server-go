package react

import (
//	"log"
	"fmt"
	"strings"
)

type Element struct {
	component	string
	content		string
	props		*Attrs
	children	[]Element
}

func (self Element) String() string {

	var out strings.Builder

	/* Open */
	if self.component != "" {
		if self.props == nil {
			fmt.Fprintf(&out, "<%s>", self.component)
		} else {
			props := self.props.String()
			fmt.Fprintf(&out, "<%s %s>", self.component, props)
		}
	}

	/* Body */
	if self.content != "" {
		fmt.Fprintf(&out, "%s", self.content)
	}

	/* Children */
	for _, e := range self.children {
		fmt.Fprintf(&out, "%s", e)
	}

	/* Close */
	if self.component != "" {
		fmt.Fprintf(&out, "</%s>", self.component)
	}

	return out.String()
}

func CreateText(content string) Element {
	return Element{
		content: content,
	}
}

/**
 * CreateElement(component, props, ...children)
 */
func CreateElement(component string, props *Attrs, children ...Element) Element {
	return Element{
		component: component,
		props: props,
		children: children,
	}
}
