package htm_test

import (
	"fmt"
	"github.com/protamail/htm"
	"strconv"
	"testing"
)

type HTML = htm.HTML

var _el, _vel, attr, append = htm.Element, htm.VoidElement, htm.Attributes, htm.Append
var printf, itoa = fmt.Sprintf, strconv.Itoa
var henc, uenc, I = htm.HTMLEncode, htm.URIComponentEncode, htm.AsIs
var empty = HTML{}

func Test1(t *testing.T) {
	//var r HTML
	//
	//		<html class="heh" data-href="sdsd?sds=1">
	//			<body>
	//				<nav class="heh" data-href="sdsd?sds=1">
	//					<div>
	//						<ul>
	//							<li data-href="hj&'gjh&ha=wdfw eee"></li>
	//							<img src="j">
	//							<br>
	//							<span data-href="ddd">dsdsdsd</span>
	//							...
	//						</ul>
	//					</div>
	//				</nav>
	//			</body>
	//		</html>
	//
	for i := 0; i < 1000; i++ {
		_ = _el("html", attr("class=", "heh", "data-href=", "sdsd?sds=1"),
			_el("body", "",
				_el("nav", attr("class=", "heh", "data-href=", "sdsd?sds=1"),
					_el("div", "",
						_el("ul", "", func() HTML {
							var result HTML
							for j := 0; j < 1000; j++ {
								result = append(result,
									_el("li", attr("data-href=", uenc(`hj&"'>gjh`)+`&ha=`+uenc(`wdfw&`)+func() string {
										if true {
											return "&eee"
										}
										return ""
									}()), henc(printf("%d", j))),
									_vel("img", attr("src=", printf("img%d", j))),
									_vel("img", attr("src=", itoa(j))),
									_vel(`img`, attr("src=", printf("img%.2f", float32(j)))),
									_vel("br", ""),
									_el("span", attr("data-href", "ddd"), henc("dsdsi&dsd")),
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

/*
import (
	"bytes"
	"html/template"
	"testing"
)

func Test1(t *testing.T) {
	var tpl bytes.Buffer
	tmpl, _ := template.New("foo").Parse(`
	{{range $idx, $e := .}}
	<li><a href="/?page={{$idx}}{{$idx}}">{{$idx}}{{$idx}}</a></li>
	{{end}}
	`)
	//<li><a href="/?page={{$idx}}">{{$idx}}</a></li>
	var a = make([]struct{}, 1000)
	for i := 0; i < 1000; i++ {
		tmpl.Execute(&tpl, a)
	}
}
*/
