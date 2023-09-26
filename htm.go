package htm

import (
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragments
type Safe struct {
	frag []string
}

type Attr string

func Element(tag string, attr Attr, body Safe) Safe {
	ss := make([]string, 0, len(body.frag)+8)
	if len(attr) > 0 {
		ss = append(ss, "<", tag, " ", string(attr), "\n>")
	} else {
		ss = append(ss, "<", tag, ">")
	}
	ss = append(ss, body.frag...)
	ss = append(ss, "</", tag, ">")
	return Safe{[]string{strings.Join(ss, "")}}
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
	return Safe{[]string{"<" + tag + " " + string(attr) + "\n>"}}
}

func Join(ss ...Safe) Safe {
	r := make([]string, 0, len(ss))
	for _, s := range ss {
		r = append(r, strings.Join(s.frag, ""))
	}
	//Join mostly used to accumulate in cycle, so don't join result yet
	return Safe{r}
}

func (c Safe) String() string {
	return strings.Join(c.frag, "")
}

func AsIs(a ...string) Safe {
	return Safe{a}
}

// Used to output HTML text, escaping HTML reserved characters <>&"
func HTMLEncode(a string) Safe {
	return Safe{[]string{html.EscapeString(a)}}
}

var URIComponentEncode = url.QueryEscape

var jsStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) Safe {
	return Safe{[]string{jsStringEscaper.Replace(a)}}
}
