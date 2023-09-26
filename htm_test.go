package htm_test

import (
	"github.com/protamail/htm"
	"strconv"
	"testing"
)

type Safe = htm.Safe

var E, V, A, J = htm.Element, htm.VoidElement, htm.Attributes, htm.Join
var H, U, I = htm.HTMLEncode, htm.URIComponentEncode, htm.AsIs
var empty = Safe{}

func Test1(t *testing.T) {
//	var r Safe
	for i := 0; i < 1000; i++ {
		var collect Safe
		for j := 0; j < 1000; j++ {
			collect = J(collect,
				E("li", A(`data-href`, U(`hj&"'>gjh`)+`&ha=`+U(`wdfw&`)+func() string {
					if true {
						return " eee"
					}
					return ""
				}()), empty),
				V("img", A(`src`, `img`+strconv.Itoa(j))),
				V("br", ""),
				E("span", A("data-href", "ddd"), H("dsdsdsd")),
				E("span", A("data-href", "ddd"), H("dsdsdsd")),
				E("span", A("data-href", "ddd"), H("dsdsdsd")),
				E("span", A("data-href", "ddd"), H("dsdsdsd")),
				E("span", A("data-href", "ddd"), H("dsdsdsd")),
				E("span", A("data-href", "ddd"), H("dsdsdsd")),
				V("br", ""),
			)
		}
		_ = E("html", A(`class`, `heh`, `data-href`, "sdsd?sds=1"),
			E("body", "",
				E("nav", A(`class`, "heh", `data-href`, "sdsd?sds=1"),
					E("div", "",
						E("ul", "", collect),
					),
				),
			),
		)
//		fmt.Println(r.String())
	}
}
