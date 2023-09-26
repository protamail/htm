package htm_test

import (
	_ "fmt"
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
		_ = E("html", A(`class`, `heh`, `data-href`, "sdsd?sds=1"),
			E("body", "",
				E("nav", A(`class`, "heh", `data-href`, "sdsd?sds=1"),
					E("div", "",
						E("ul", "", func() Safe {
							l := 1000
							result := make([]Safe, 0, l)
							for j := 0; j < l; j++ {
								result = append(result,
									E("li", A(`data-href`, U(`hj&"'>gjh`)+`&ha=`+U(`wdfw&`)+func() string {
										if true {
											return " eee"
										}
										return ""
									}()), empty),
									V("img", A(`src`, `img`+strconv.Itoa(j))),
									V("br", ""),
									E("span", A("data-href", "ddd"), H("dsdsdsd")),
								)
							}
							return J(result)
						}()),
					),
				),
			),
		)
		//		fmt.Println(r.String())
	}
}
