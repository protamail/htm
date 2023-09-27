package htm

import (
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragment
type HTML struct {
	pieces	[]string
}

func NewHTML(body string) HTML {
	return HTML{[]string{"", body, ""}}
}

type Attr string

func Element(tag string, attr Attr, body HTML) HTML {
	h := NewHTML("")
	if len(attr) > 0 {
		h.pieces[0] = "<" + tag + " " + string(attr) + "\n>" + body.pieces[0]
	} else {
		h.pieces[0] = "<" + tag + ">" + body.pieces[0]
	}
	h.pieces[1] = body.pieces[1]
	h.pieces[2] = body.pieces[2] + "</" + tag + ">"
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

join ->creates flattened pieces array
String() -> concats
func Join(frags ...HTML) HTML {
	switch len(frags) {
	case 0:
		return NewHTML("")
	case 1:
		return frags[0]
	}

	var n int
	for _, frag := range frags {
		n += len(frag.before) + len(frag.body) + len(frag.after)
	}

	var b strings.Builder
	b.Grow(n)
	for _, frag := range frags {
		if len(frag.before) > 0 {
			b.WriteString(frag.before)
		}
		if len(frag.body) > 0 {
			b.WriteString(frag.body)
		}
		if len(frag.after) > 0 {
			b.WriteString(frag.after)
		}
	}
	return NewHTML(b.String())
}

func (c HTML) String() string {
	return c.before + c.body + c.after
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
