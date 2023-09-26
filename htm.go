package htm

import (
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragment
type HTML struct {
	before string
	body   string
	after  string
}

type Attr string

func Element(tag string, attr Attr, body HTML) HTML {
	h := HTML{}
	if len(attr) > 0 {
		h.before = "<" + tag + " " + string(attr) + "\n>" + body.before
	} else {
		h.before = "<" + tag + ">" + body.before
	}
	h.body = body.body
	h.after = body.after + "</" + tag + ">"
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
	return HTML{"", "<" + tag + " " + string(attr) + "\n>", ""}
}

func Join(frags ...HTML) HTML {
	switch len(frags) {
	case 0:
		return HTML{}
	case 1:
		return frags[0]
	}

	first := frags[0]
	n := len(first.body) + len(first.after)
	for _, frag := range frags[1:] {
		n += len(frag.before) + len(frag.body) + len(frag.after)
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(first.body)
	b.WriteString(first.after)
	for _, frag := range frags[1:] {
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
	first.body = b.String()
	first.after = ""
	return first
}

func (c HTML) String() string {
	return strings.Join([]string{c.before, c.body, c.after}, "")
}

func AsIs(a ...string) HTML {
	return HTML{"", strings.Join(a, ""), ""}
}

// Used to output HTML text, escaping HTML reserved characters <>&"
func HTMLEncode(a string) HTML {
	return HTML{"", html.EscapeString(a), ""}
}

var URIComponentEncode = url.QueryEscape

var jsStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) HTML {
	return HTML{"", jsStringEscaper.Replace(a), ""}
}
