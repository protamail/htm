package htm_test

import (
	"fmt"
	"github.com/protamail/htm"
	"strconv"
	"testing"
)

type HTML = htm.HTML

var _el, attr, add = htm.Element, htm.Attributes, htm.Append
var printf, itoa = fmt.Sprintf, strconv.Itoa
var henc, uenc, id = htm.HTMLEncode, htm.URIComponentEncode, htm.AsIs

func Test1(t *testing.T) {
	type B struct {
		a string
		B int
	}
	var b = B{"heh", 2}
	fmt.Printf(htm.See(1, b))
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
	a := make([]string, 1000, 1000)
	for i := 0; i < 1000; i++ {
		r :=
			_el("html", attr("class=", "heh", "data-href=", "sdsd?sds=1"),
				_el("body", "",
					_el("nav", attr("class=", "heh", "data-href=", "sdsd?sds=1"),
						_el("div", "",
							/*_el("ul", "", func() HTML {
								var result = htm.NewHTML(1000)
								for j := 0; j < 1000; j++ {
									result = add(result,
										_el("li", attr("data-href=", uenc(`hj&"'>gjh`)+`&ha=`+uenc(`wdfw&`)+func() string {
											if true {
												return "&eee"
											}
											return ""
										}()),
											henc(printf("%d", j)),
											_el("img", attr("src=", printf("img%d", j))),
											_el("img", attr("src=", itoa(j))),
											_el(`img`, attr("src=", printf("img%.2f", float32(j)))),
											_el("br", ""),
											_el("div", "", henc("heh"), id("da"), henc("boom")),
											_el("span", attr("data-href", "ddd"), henc("dsdsi&dsd")),
										),
									)
								}
								return result
							}()),*/
							_el("ul", "", htm.Map(a, func(j int) HTML {
								return _el("li", attr("data-href=", uenc(`hj&"'>gjh`)+`&ha=`+uenc(`wdfw&`)+func() string {
									if true {
										return "&eee"
									}
									return ""
								}()),
									henc(printf("%d", j)),
									_el("img", attr("src=", printf("img%d", j))),
									_el("img", attr("src=", itoa(j))),
									_el(`img`, attr("src=", printf("img%.2f", float32(j)))),
									_el("br", ""),
									_el("div", "", henc("heh"), id("da"), henc("boom")),
									_el("span", attr("data-href", "ddd"), henc("dsdsi&dsd"), henc(a[j])),
								)
							})),
						),
					),
				),
			)
		_ = r
		//		fmt.Println(r.String())
	}
}

func aTest2(t *testing.T) {
	var buckets = []map[string]string{
		{"bucket": "WLGCRU", "bucketName": "Wireline Growth & CRU"},
		{"bucket": "TOTAL", "bucketName": "Total"},
	}
	var listHeader = func() HTML {
		result :=
			_el("tr", attr("class=", "tr-hdr trb-t trb-s trb-b narrow-font"),
				_el("td", attr("class=", "tdb-l"), _el("br", "")),
				_el("td", "", id("PID")),
				_el("td", "", id("RVP")),
				_el("td", "", id("Sales Center")),
				func() HTML {
					var result HTML
					for _, b := range buckets {
						result = add(result, _el("td", "", henc(b["bucketName"])))
					}
					return result
				}())
		return result
	}
	fmt.Println(listHeader().String())
	/*
	   var listGroup = (arr, lastRegGroup) => !arr? false :

	   	arr.map((i, idx) => Htm`<tr
	   	    class="trb-s ${idx == arr.length-1 && ((i.rmtType == '2' || lastRegGroup) && 'trb-b' || 'trb-b-dot')} ${i.rmtType == '2' && ' tr-rollup'}">
	   	    <td class="tdb-l">${format(rowNum++, '#')}</td>
	   	    <td>${i.rmtid.length == 4 && i.rmtid}</td>
	   	    <td>${i.parentName}</td>
	   	    <td>${i.name}</td>
	   	    ${t.bucket.map(b => {
	   	        let q = t.quota[i.rmtid] && t.quota[i.rmtid][b.bucket];
	   	        return Htm`<td ${i.rmtType != '2' && Htm`class="ptr ${q && q.quota < 0 && 'red' || 'blue'}" data-href="ajaxDoUpdate?rmtid=${encodeURIComponent(i.rmtid)}&bucket=${b.bucket}&field=quota&v=${q && q.quota || 0}#ShowInput"`}>
	   	            ${q && format(q.quota, '0.000;;-')}</td>`;
	   	    })}
	   	</tr>`);

	   return Htm`<table class="quota-input" style="table-layout: fixed">

	   	<tr>
	   	    <td class="width2"></td>
	   	    <td class="width5"></td>
	   	    <td class="width12"></td>
	   	    <td class="width20"></td>
	   	    ${t.bucket.map(b => Htm`<td class="width7"></td>`)}
	   	</tr>
	   	<tr class="tr-hdr-name"><td colspan="100">${t.currentOrg.name}<span class="note">(Enter in millions of dollars)</span></td></tr>
	   	${listHeader()}
	   	${!t.orgRollup.length && listGroup(t.orgRegMap[Object.keys(t.orgRegMap)[0]], true)}
	   	${t.orgRollup.concat.apply([], t.orgRollup)
	   	    .map((i, idx, arr) => listGroup(t.orgRegMap[`${i.idPath}${i.rmtid}|`], idx == arr.length - 1))}
	   	${t.orgRollup.map(i => listGroup(i))}
	   	${t.orgRollup.length > 0 && Htm`
	   	    <tr class="trb-b"><td colspan="100"><br></td></tr>
	   	    ${!rowNum++}
	   	    ${listGroup([{rmtid: t.currentOrg.rmtid+'=', parentName: '', name: 'Control Total', rmtType: '1'}], true)}
	   	    ${listGroup([{rmtid: 'check', parentName: '', name: 'Check', rmtType: '2'}])}
	   	`}

	   </table>`;
	*/
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
