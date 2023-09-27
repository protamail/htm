package htm

import (
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragment
type HTML struct {
	pieces []string
}

type Attr string

func Element(tag string, attr Attr, body HTML) HTML {
	var r HTML
	switch len(body.pieces) {
	case 0:
		r = HTML{make([]string, 2, 2)}
	case 1:
		r = HTML{[]string{"", body.pieces[0], ""}}
	case 2:
		r = HTML{[]string{body.pieces[0], body.pieces[1]}}
	default:
		r = body
	}
	if len(attr) > 0 {
		r.pieces[0] = "<" + tag + " " + string(attr) + "\n>" + r.pieces[0]
	} else {
		r.pieces[0] = "<" + tag + ">" + r.pieces[0]
	}
	r.pieces[len(r.pieces)-1] += "</" + tag + ">"
	return r
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
	return HTML{[]string{"<" + tag + " " + string(attr) + "\n>"}}
}

func Join(collect HTML, frags ...HTML) HTML {
	var n int
	for _, frag := range frags {
		n += len(frag.pieces)
	}
	if cap(collect.pieces) < len(collect.pieces)+n {
		var newPieces []string
		if len(collect.pieces) > n {
			newPieces = make([]string, 0, len(collect.pieces)*2)
		} else {
			newPieces = make([]string, 0, len(collect.pieces)+n)
		}
		collect.pieces = append(newPieces, collect.pieces...)
	}

	np := make([]string, 0, n)
	for _, frag := range frags {
		np = append(np, frag.pieces...)
	}
	return HTML{np}
}

func (c HTML) String() string {
	return strings.Join(c.pieces, "")
}

func AsIs(a ...string) HTML {
	return HTML{[]string{strings.Join(a, "")}}
}

// Used to output HTML text, escaping HTML reserved characters <>&"
func HTMLEncode(a string) HTML {
	return HTML{[]string{html.EscapeString(a)}}
}

var URIComponentEncode = url.QueryEscape

var jsStringEscaper = strings.NewReplacer(
	`"`, `\"`,
	`'`, `\'`,
	"`", "\\`",
	`\`, `\\`,
)

func JSStringEscape(a string) HTML {
	return HTML{[]string{jsStringEscaper.Replace(a)}}
}
