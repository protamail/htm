package htm

import (
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragment
type HTML struct {
	html string
}

type Attr string

func Element(tag string, attr Attr, body ...HTML) HTML {
	sar := make([]string, 0, 9)
	if len(attr) > 0 {
		sar = append(sar, "<", tag, " ", string(attr), "\n>")
	} else {
		sar = append(sar, "<", tag, ">")
	}
	sar = append(sar, Join(body...).html)
	sar = append(sar, "</", tag, ">")
	return HTML{strings.Join(sar, "")}
}

var attrEscaper = strings.NewReplacer(`"`, `&quot;`)

func Attributes(kv ...string) Attr {
	sar := make([]string, 0, len(kv)*5/2)
	for i := 1; i < len(kv); i += 2 {
		if i > 1 {
			sar = append(sar, ` `)
		}
		v := kv[i]
		if strings.Index(v, `"`) >= 0 {
			v = attrEscaper.Replace(v)
		}
		sar = append(sar, kv[i-1], `="`, v, `"`)
	}
	return Attr(strings.Join(sar, ""))
}

// create HTML tag with no closing, e.g. <input type="text">
func VoidElement(tag string, attr Attr) HTML {
	return HTML{"<" + tag + " " + string(attr) + "\n>"}
}

func Join(frags ...HTML) HTML {
	switch len(frags) {
	case 1:
		return frags[0]
	case 0:
		return HTML{}
	}

	var n int
	for _, frag := range frags {
		n += len(frag.html)
	}

	var b strings.Builder
	b.Grow(n)
	for _, s := range frags {
		b.WriteString(s.html)
	}
	return HTML{b.String()}
}

func (c HTML) String() string {
	return c.html
}

func AsIs(a ...string) HTML {
	return HTML{strings.Join(a, "")}
}

// Used to output HTML text, escaping HTML reserved characters <>&"
func HTMLEncode(a string) HTML {
	return HTML{html.EscapeString(a)}
}

var URIComponentEncode = url.QueryEscape

var jsStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) HTML {
	return HTML{jsStringEscaper.Replace(a)}
}
