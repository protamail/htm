package htm

import (
	"fmt"
	"html"
	"net/url"
	"strings"
)

// contains well-formed HTML fragment
type HTML struct {
	pieces []string
}

type Attr string

var voidEl = map[string]bool{"area": true, "base": true, "br": true, "col": true, "command": true, "embed": true, "hr": true, "img": true, "input": true, "keygen": true, "link": true, "meta": true, "source": true, "track": true, "wbr": true}

func NewHTML(cap int) HTML {
	return HTML{make([]string, 0, cap)}
}

func Element(tag string, attr Attr, bodyEls ...HTML) HTML {
	var r, body HTML
	switch len(bodyEls) {
	case 0:
		if voidEl[tag] || voidEl[strings.ToLower(tag)] {
			return HTML{[]string{"<" + tag + string(attr) + "\n>"}}
		}
	case 1:
		body = bodyEls[0]
	default:
		body = Append(body, bodyEls...)
	}

	switch len(body.pieces) {
	case 0:
		r = HTML{make([]string, 1, 1)}
	case 1:
		if len(body.pieces[0]) < 256 {
			return HTML{[]string{"<" + tag + string(attr) + "\n>" + body.pieces[0] + "</" + tag + ">"}}
		}
		r = HTML{[]string{"", body.pieces[0], ""}}
	default:
		r = body
	}
	r.pieces[0] = "<" + tag + string(attr) + "\n>" + r.pieces[0]
	r.pieces[len(r.pieces)-1] += "</" + tag + ">"
	return r
}

var attrEscaper = strings.NewReplacer(`"`, URIComponentEncode(`"`))

func Prepend(doctype string, html HTML) HTML {
	if len(html.pieces) > 0 {
		html.pieces[0] = doctype + html.pieces[0]
		return html
	}
	return HTML{[]string{doctype}}
}

func Attributes(nv ...string) Attr {
	sar := make([]string, 0, len(nv)*5/2)
	for i := 1; i < len(nv); i += 2 {
		sar = append(sar, " ")
		k, v := nv[i-1], nv[i]
		if strings.Index(v, `"`) >= 0 {
			v = attrEscaper.Replace(v)
		}
		if k[len(k)-1] == 61 { //if already ends with =
			sar = append(sar, k, `"`, v, `"`)
		} else { //if attr key is not ending with =, output bare attrribute discarding the value
			sar = append(sar, k)
		}
	}
	return Attr(strings.Join(sar, ""))
}

func JoinAttributes(attrs ...Attr) Attr {
	var n int

	for _, attr := range attrs {
		n += len(attr)
	}

	var b strings.Builder
	b.Grow(n)

	for _, attr := range attrs {
		b.WriteString(string(attr))
	}

	return Attr(b.String())
}

func See(what ...any) string {
	return Map(what, func(i int) HTML {
		return HTML{[]string{fmt.Sprintf("%+v\n", what[i])}}
	}).String()
}

func Map[T any](a []T, f func(int) HTML) HTML {
	r := NewHTML(len(a))
	for i := range a {
		r = Append(r, f(i))
	}
	return r
}

func If[T ~string | HTML](cond bool, result T) T {
	if cond {
		return result
	}
	var r T
	return r
}

func IfCall[T ~string | HTML](cond bool, call func() T) T {
	if cond {
		return call()
	}
	var r T
	return r
}

func IfElse[T ~string | HTML](cond bool, ifR T, elseR T) T {
	if cond {
		return ifR
	}
	return elseR
}

func IfElseCall[T ~string | HTML](cond bool, ifCall func() T, elseCall func() T) T {
	if cond {
		return ifCall()
	}
	return elseCall()
}

func Append(collect HTML, frags ...HTML) HTML {
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

	for _, frag := range frags {
		collect.pieces = append(collect.pieces, frag.pieces...)
	}
	return collect
}

func (c HTML) String() string {
	return strings.Join(c.pieces, "")
}

func AsIs(a ...string) HTML {
	return HTML{a}
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
