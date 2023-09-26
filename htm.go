package htm

import (
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragment
type HTML struct {
	html *[3]string
}

type Attr string

func NewHTML(body string) HTML {
	return HTML{&[3]string{"", body, ""}}
}

func Element(tag string, attr Attr, body HTML) HTML {
	h := NewHTML("")
	if len(attr) > 0 {
		h.html[0] = "<" + tag + " " + string(attr) + "\n>" + body.html[0]
	} else {
		h.html[0] = "<" + tag + ">" + body.html[0]
	}
	h.html[1] = body.html[1]
	h.html[2] = body.html[2] + "</" + tag + ">"
	return h
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
	return NewHTML("<" + tag + " " + string(attr) + "\n>")
}

func Join(frags ...HTML) HTML {
	switch len(frags) {
	case 0:
		return NewHTML("")
	case 1:
		return frags[0]
	}

	first := frags[0]
	n := len(first.html[1]) + len(first.html[2])
	for _, frag := range frags[1:] {
		n += len(frag.html[0]) + len(frag.html[1]) + len(frag.html[2])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(first.html[1])
	b.WriteString(first.html[2])
	for _, frag := range frags[1:] {
		if len(frag.html[0]) > 0 {
			b.WriteString(frag.html[0])
		}
		if len(frag.html[1]) > 0 {
			b.WriteString(frag.html[1])
		}
		if len(frag.html[2]) > 0 {
			b.WriteString(frag.html[2])
		}
	}
	first.html[1] = b.String()
	first.html[2] = ""
	return first
}

func (c HTML) String() string {
	return strings.Join(c.html[:], "")
}

func AsIs(a ...string) HTML {
	return NewHTML(strings.Join(a, ""))
}

// Used to output HTML text, escaping HTML reserved characters <>&"
func HTMLEncode(a string) HTML {
	return NewHTML(html.EscapeString(a))
}

var URIComponentEncode = url.QueryEscape

var jsStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) HTML {
	return NewHTML(jsStringEscaper.Replace(a))
}
