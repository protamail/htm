package htm_test

import (
	//	"fmt"
	"github.com/protamail/htm"
	"strconv"
	"testing"
)

type HTML = htm.HTML

var E, V, T, A = htm.Element, htm.VoidElement, htm.Attributes, htm.Append
var H, U, I = htm.HTMLEncode, htm.URIComponentEncode, htm.AsIs
var empty = HTML{}

func Test1(t *testing.T) {
	//	var r HTML
	for i := 0; i < 1000; i++ {
		_ = E(`html`, T(`class`, "heh", `data-href`, "sdsd?sds=1"),
			E(`body`, "",
				E(`nav`, T(`class`, "heh", `data-href`, "sdsd?sds=1"),
					E(`div`, ``,
						E(`ul`, ``, func() HTML {
							var result HTML
							for j := 0; j < 1000; j++ {
								result = A(result,
									E("li", T(`data-href`, `hj&"'>gjh`+`&ha=`+U(`wdfw&`)+func() string {
										if true {
											return " eee"
										}
										return ""
									}()), empty),
									V(`img`, T(`src`, `img`+strconv.Itoa(j))),
									V(`br`, ""),
									E(`span`, T(`data-href`, `ddd`), H(`dsdsdsd`)),
								)
							}
							return result
						}()),
					),
				),
			),
		)
		//		fmt.Println(r.String())
	}
}
