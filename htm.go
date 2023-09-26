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
	//smaller fragments are concatenated as soon as available, larger ones are deferred
	ss := make([]string, 0, len(body.frag)+2)
	if len(attr) > 0 {
		ss = append(ss, "<"+tag+" "+string(attr)+"\n>")
	} else {
		ss = append(ss, "<"+tag+">")
	}
	ss = append(ss, body.frag...)
	ss = append(ss, "</"+tag+">")
	return Safe{ss}
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

func Join(dst Safe, src ...Safe) Safe {
	for _, s := range src {
		if len(s.frag) > 1 {
			dst.frag = append(dst.frag, strings.Join(s.frag, ""))
		} else {
			dst.frag = append(dst.frag, s.frag...)
		}
	}
	return dst
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
