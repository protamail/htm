package htm

import (
	"html"
	"net/url"
	"strings"
)

type Safe struct {
	frag []string
}

func Element(tag string, attr string, body Safe) Safe {
	ss := make([]string, 0, len(body.frag)+2)
	if len(attr) > 0 {
		ss = append(ss, strings.Join([]string{"<", tag, " ", attr, "\n>"}, ""))
	} else {
		ss = append(ss, strings.Join([]string{"<", tag, ">"}, ""))
	}
	ss = append(ss, body.frag...)
	ss = append(ss, strings.Join([]string{"</", tag, ">"}, ""))
	return Safe{ss}
}

func VoidElement(tag string, attr ...string) Safe {
	return Safe{[]string{strings.Join([]string{"<", tag, " ", strings.Join(attr, ""), "\n>"}, "")}}
}

func Append(dst Safe, src ...Safe) Safe {
	for _, s := range src {
		if len(s.frag) > 0 && len(s.frag[0]) > 256 {
			dst.frag = append(dst.frag, s.frag...)
		} else {
			dst.frag = append(dst.frag, strings.Join(s.frag, ""))
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

func URIComponentEncode(a string) Safe {
	return Safe{[]string{url.QueryEscape(a)}}
}

var JSStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) Safe {
	return Safe{[]string{JSStringEscaper.Replace(a)}}
}
