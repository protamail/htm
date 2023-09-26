package htm

import (
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragments
type Safe struct {
	html string
}

type Attr string

func Element(tag string, attr Attr, body ...Safe) Safe {
	ss := make([]string, 0, 9)
	if len(attr) > 0 {
		ss = append(ss, "<", tag, " ", string(attr), "\n>")
	} else {
		ss = append(ss, "<", tag, ">")
	}
	ss = append(ss, Join(body...).html)
	ss = append(ss, "</", tag, ">")
	return Safe{strings.Join(ss, "")}
}

var attrEscaper = strings.NewReplacer(`"`, `&quot;`)

func Attributes(kv ...string) Attr {
	ss := make([]string, 0, len(kv)*5/2)
	for i := 1; i < len(kv); i += 2 {
		if i > 1 {
			ss = append(ss, ` `)
		}
		v := kv[i]
		if strings.Index(v, `"`) >= 0 {
			v = attrEscaper.Replace(v)
		}
		ss = append(ss, kv[i-1], `="`, v, `"`)
	}
	return Attr(strings.Join(ss, ""))
}

// create HTML tag with no closing, e.g. <input type="text">
func VoidElement(tag string, attr Attr) Safe {
	return Safe{"<" + tag + " " + string(attr) + "\n>"}
}

/*func Join(ss ...Safe) Safe {
	r := make([]string, 0, len(ss))
	for _, s := range ss {
		r = append(r, s.html)
	}
	return Safe{strings.Join(r, "")}
}*/

func Join(frags ...Safe) Safe {
	switch len(frags) {
	case 0:
		return Safe{}
	case 1:
		return Safe{frags[0].html}
	}

	var n int
	for _, frag := range frags {
		n += len(frag.html)
		if n < len(frag.html) {
			panic("htm: Join output length overflow")
		}
	}

	var b strings.Builder
	b.Grow(n)
	for _, s := range frags {
		b.WriteString(s.html)
	}
	return Safe{b.String()}
}

func (c Safe) String() string {
	return c.html
}

func AsIs(a ...string) Safe {
	return Safe{strings.Join(a, "")}
}

// Used to output HTML text, escaping HTML reserved characters <>&"
func HTMLEncode(a string) Safe {
	return Safe{html.EscapeString(a)}
}

var URIComponentEncode = url.QueryEscape

var jsStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) Safe {
	return Safe{jsStringEscaper.Replace(a)}
}
