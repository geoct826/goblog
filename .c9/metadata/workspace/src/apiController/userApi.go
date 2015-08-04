{"filter":false,"title":"userApi.go","tooltip":"/src/apiController/userApi.go","undoManager":{"mark":100,"position":100,"stack":[[{"start":{"row":118,"column":0},"end":{"row":119,"column":0},"action":"remove","lines":["",""],"id":5101}],[{"start":{"row":122,"column":3},"end":{"row":123,"column":2},"action":"remove","lines":["","\t\t"],"id":5102}],[{"start":{"row":127,"column":3},"end":{"row":128,"column":0},"action":"remove","lines":["",""],"id":5103}],[{"start":{"row":219,"column":17},"end":{"row":219,"column":18},"action":"insert","lines":["G"],"id":5104}],[{"start":{"row":219,"column":18},"end":{"row":219,"column":19},"action":"insert","lines":["E"],"id":5105}],[{"start":{"row":219,"column":19},"end":{"row":219,"column":20},"action":"insert","lines":["T"],"id":5106}],[{"start":{"row":219,"column":20},"end":{"row":219,"column":21},"action":"insert","lines":[" "],"id":5107}],[{"start":{"row":316,"column":0},"end":{"row":317,"column":0},"action":"insert","lines":["",""],"id":5108}],[{"start":{"row":317,"column":0},"end":{"row":322,"column":0},"action":"insert","lines":["\t\t\tif user.Email == \"\" {","\t\t\t\tlog.Println(\"GET /api/users/userReqID: user not found\", userReqID)","\t\t\t\tnotFound(w, r)","\t\t\t\treturn","\t\t\t}",""],"id":5109}],[{"start":{"row":317,"column":0},"end":{"row":317,"column":1},"action":"remove","lines":["\t"],"id":5110},{"start":{"row":318,"column":0},"end":{"row":318,"column":1},"action":"remove","lines":["\t"]},{"start":{"row":319,"column":0},"end":{"row":319,"column":1},"action":"remove","lines":["\t"]},{"start":{"row":320,"column":0},"end":{"row":320,"column":1},"action":"remove","lines":["\t"]},{"start":{"row":321,"column":0},"end":{"row":321,"column":1},"action":"remove","lines":["\t"]}],[{"start":{"row":317,"column":0},"end":{"row":317,"column":1},"action":"remove","lines":["\t"],"id":5111},{"start":{"row":318,"column":0},"end":{"row":318,"column":1},"action":"remove","lines":["\t"]},{"start":{"row":319,"column":0},"end":{"row":319,"column":1},"action":"remove","lines":["\t"]},{"start":{"row":320,"column":0},"end":{"row":320,"column":1},"action":"remove","lines":["\t"]},{"start":{"row":321,"column":0},"end":{"row":321,"column":1},"action":"remove","lines":["\t"]}],[{"start":{"row":317,"column":8},"end":{"row":317,"column":9},"action":"insert","lines":["E"],"id":5112}],[{"start":{"row":317,"column":9},"end":{"row":317,"column":10},"action":"insert","lines":["d"],"id":5113}],[{"start":{"row":317,"column":10},"end":{"row":317,"column":11},"action":"insert","lines":["i"],"id":5114}],[{"start":{"row":317,"column":11},"end":{"row":317,"column":12},"action":"insert","lines":["t"],"id":5115}],[{"start":{"row":317,"column":12},"end":{"row":317,"column":13},"action":"insert","lines":["e"],"id":5116}],[{"start":{"row":317,"column":13},"end":{"row":317,"column":14},"action":"insert","lines":["d"],"id":5117}],[{"start":{"row":318,"column":17},"end":{"row":318,"column":18},"action":"remove","lines":["T"],"id":5118}],[{"start":{"row":318,"column":16},"end":{"row":318,"column":17},"action":"remove","lines":["E"],"id":5119}],[{"start":{"row":318,"column":15},"end":{"row":318,"column":16},"action":"remove","lines":["G"],"id":5120}],[{"start":{"row":318,"column":15},"end":{"row":318,"column":16},"action":"insert","lines":["P"],"id":5121}],[{"start":{"row":318,"column":16},"end":{"row":318,"column":17},"action":"insert","lines":["O"],"id":5122}],[{"start":{"row":318,"column":17},"end":{"row":318,"column":18},"action":"insert","lines":["S"],"id":5123}],[{"start":{"row":318,"column":18},"end":{"row":318,"column":19},"action":"insert","lines":["T"],"id":5124}],[{"start":{"row":318,"column":30},"end":{"row":318,"column":40},"action":"remove","lines":["/userReqID"],"id":5125}],[{"start":{"row":318,"column":57},"end":{"row":318,"column":58},"action":"remove","lines":["D"],"id":5126}],[{"start":{"row":318,"column":56},"end":{"row":318,"column":57},"action":"remove","lines":["I"],"id":5127}],[{"start":{"row":318,"column":55},"end":{"row":318,"column":56},"action":"remove","lines":["q"],"id":5128}],[{"start":{"row":318,"column":54},"end":{"row":318,"column":55},"action":"remove","lines":["e"],"id":5129}],[{"start":{"row":318,"column":53},"end":{"row":318,"column":54},"action":"remove","lines":["R"],"id":5130}],[{"start":{"row":318,"column":52},"end":{"row":318,"column":53},"action":"remove","lines":["r"],"id":5131}],[{"start":{"row":318,"column":51},"end":{"row":318,"column":52},"action":"remove","lines":["e"],"id":5132}],[{"start":{"row":318,"column":50},"end":{"row":318,"column":51},"action":"remove","lines":["s"],"id":5133}],[{"start":{"row":318,"column":49},"end":{"row":318,"column":50},"action":"remove","lines":["u"],"id":5134}],[{"start":{"row":318,"column":48},"end":{"row":318,"column":49},"action":"remove","lines":[" "],"id":5135}],[{"start":{"row":318,"column":47},"end":{"row":318,"column":48},"action":"remove","lines":[","],"id":5136}],[{"start":{"row":317,"column":0},"end":{"row":322,"column":0},"action":"remove","lines":["\tif userEdited.Email == \"\" {","\t\tlog.Println(\"POST /api/users: user not found\")","\t\tnotFound(w, r)","\t\treturn","\t}",""],"id":5137}],[{"start":{"row":308,"column":0},"end":{"row":313,"column":0},"action":"remove","lines":["\t\tif userEdited.Email != userPost.Email {","\t\t\tlog.Println(\"Unauthorized User Post\")","\t\t\tunauthorized(w,r)","\t\t\treturn","\t\t}",""],"id":5138},{"start":{"row":308,"column":0},"end":{"row":313,"column":0},"action":"insert","lines":["\tif userEdited.Email == \"\" {","\t\tlog.Println(\"POST /api/users: user not found\")","\t\tnotFound(w, r)","\t\treturn","\t}",""]}],[{"start":{"row":308,"column":0},"end":{"row":308,"column":1},"action":"insert","lines":["\t"],"id":5139},{"start":{"row":309,"column":0},"end":{"row":309,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":310,"column":0},"end":{"row":310,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":311,"column":0},"end":{"row":311,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":312,"column":0},"end":{"row":312,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":309,"column":47},"end":{"row":309,"column":48},"action":"insert","lines":[","],"id":5140}],[{"start":{"row":309,"column":48},"end":{"row":309,"column":49},"action":"insert","lines":[" "],"id":5141}],[{"start":{"row":309,"column":49},"end":{"row":309,"column":50},"action":"insert","lines":["u"],"id":5142}],[{"start":{"row":309,"column":50},"end":{"row":309,"column":51},"action":"insert","lines":["s"],"id":5143}],[{"start":{"row":309,"column":51},"end":{"row":309,"column":52},"action":"insert","lines":["e"],"id":5144}],[{"start":{"row":309,"column":52},"end":{"row":309,"column":53},"action":"insert","lines":["r"],"id":5145}],[{"start":{"row":309,"column":53},"end":{"row":309,"column":54},"action":"insert","lines":["P"],"id":5146}],[{"start":{"row":309,"column":53},"end":{"row":309,"column":54},"action":"remove","lines":["P"],"id":5147}],[{"start":{"row":309,"column":52},"end":{"row":309,"column":53},"action":"remove","lines":["r"],"id":5148}],[{"start":{"row":309,"column":51},"end":{"row":309,"column":52},"action":"remove","lines":["e"],"id":5149}],[{"start":{"row":309,"column":50},"end":{"row":309,"column":51},"action":"remove","lines":["s"],"id":5150}],[{"start":{"row":309,"column":49},"end":{"row":309,"column":50},"action":"remove","lines":["u"],"id":5151}],[{"start":{"row":309,"column":48},"end":{"row":309,"column":49},"action":"remove","lines":[" "],"id":5152}],[{"start":{"row":309,"column":47},"end":{"row":309,"column":48},"action":"remove","lines":[","],"id":5153}],[{"start":{"row":309,"column":48},"end":{"row":309,"column":49},"action":"insert","lines":[","],"id":5154}],[{"start":{"row":309,"column":49},"end":{"row":309,"column":50},"action":"insert","lines":[" "],"id":5155}],[{"start":{"row":309,"column":50},"end":{"row":309,"column":51},"action":"insert","lines":["u"],"id":5156}],[{"start":{"row":309,"column":51},"end":{"row":309,"column":52},"action":"insert","lines":["p"],"id":5157}],[{"start":{"row":309,"column":52},"end":{"row":309,"column":53},"action":"insert","lines":["s"],"id":5158}],[{"start":{"row":309,"column":53},"end":{"row":309,"column":54},"action":"insert","lines":["e"],"id":5159}],[{"start":{"row":309,"column":54},"end":{"row":309,"column":55},"action":"insert","lines":["r"],"id":5160}],[{"start":{"row":309,"column":55},"end":{"row":309,"column":56},"action":"insert","lines":["P"],"id":5161}],[{"start":{"row":309,"column":55},"end":{"row":309,"column":56},"action":"remove","lines":["P"],"id":5162}],[{"start":{"row":309,"column":54},"end":{"row":309,"column":55},"action":"remove","lines":["r"],"id":5163}],[{"start":{"row":309,"column":53},"end":{"row":309,"column":54},"action":"remove","lines":["e"],"id":5164}],[{"start":{"row":309,"column":52},"end":{"row":309,"column":53},"action":"remove","lines":["s"],"id":5165}],[{"start":{"row":309,"column":51},"end":{"row":309,"column":52},"action":"remove","lines":["p"],"id":5166}],[{"start":{"row":309,"column":51},"end":{"row":309,"column":52},"action":"insert","lines":["s"],"id":5167}],[{"start":{"row":309,"column":52},"end":{"row":309,"column":53},"action":"insert","lines":["e"],"id":5168}],[{"start":{"row":309,"column":53},"end":{"row":309,"column":54},"action":"insert","lines":["r"],"id":5169}],[{"start":{"row":309,"column":54},"end":{"row":309,"column":55},"action":"insert","lines":["P"],"id":5170}],[{"start":{"row":309,"column":55},"end":{"row":309,"column":56},"action":"insert","lines":["o"],"id":5171}],[{"start":{"row":309,"column":56},"end":{"row":309,"column":57},"action":"insert","lines":["s"],"id":5172}],[{"start":{"row":309,"column":57},"end":{"row":309,"column":58},"action":"insert","lines":["t"],"id":5173}],[{"start":{"row":309,"column":58},"end":{"row":309,"column":59},"action":"insert","lines":["."],"id":5174}],[{"start":{"row":309,"column":59},"end":{"row":309,"column":60},"action":"insert","lines":["E"],"id":5175}],[{"start":{"row":309,"column":60},"end":{"row":309,"column":61},"action":"insert","lines":["m"],"id":5176}],[{"start":{"row":309,"column":61},"end":{"row":309,"column":62},"action":"insert","lines":["a"],"id":5177}],[{"start":{"row":309,"column":62},"end":{"row":309,"column":63},"action":"insert","lines":["i"],"id":5178}],[{"start":{"row":309,"column":63},"end":{"row":309,"column":64},"action":"insert","lines":["l"],"id":5179}],[{"start":{"row":313,"column":0},"end":{"row":315,"column":0},"action":"remove","lines":["\t\t","\t\tlog.Println(userOld)",""],"id":5180}],[{"start":{"row":314,"column":0},"end":{"row":315,"column":0},"action":"remove","lines":["",""],"id":5181}],[{"start":{"row":313,"column":2},"end":{"row":314,"column":0},"action":"remove","lines":["",""],"id":5182}],[{"start":{"row":313,"column":2},"end":{"row":314,"column":0},"action":"insert","lines":["",""],"id":5183},{"start":{"row":314,"column":0},"end":{"row":314,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":315,"column":11},"end":{"row":315,"column":12},"action":"remove","lines":["d"],"id":5184}],[{"start":{"row":315,"column":10},"end":{"row":315,"column":11},"action":"remove","lines":["l"],"id":5185}],[{"start":{"row":315,"column":9},"end":{"row":315,"column":10},"action":"remove","lines":["O"],"id":5186}],[{"start":{"row":315,"column":9},"end":{"row":315,"column":10},"action":"insert","lines":["E"],"id":5187}],[{"start":{"row":315,"column":10},"end":{"row":315,"column":11},"action":"insert","lines":["d"],"id":5188}],[{"start":{"row":315,"column":11},"end":{"row":315,"column":12},"action":"insert","lines":["i"],"id":5189}],[{"start":{"row":315,"column":12},"end":{"row":315,"column":13},"action":"insert","lines":["t"],"id":5190}],[{"start":{"row":315,"column":13},"end":{"row":315,"column":14},"action":"insert","lines":["e"],"id":5191}],[{"start":{"row":315,"column":14},"end":{"row":315,"column":15},"action":"insert","lines":["d"],"id":5192}],[{"start":{"row":321,"column":11},"end":{"row":321,"column":12},"action":"remove","lines":["d"],"id":5193}],[{"start":{"row":321,"column":10},"end":{"row":321,"column":11},"action":"remove","lines":["l"],"id":5194}],[{"start":{"row":321,"column":9},"end":{"row":321,"column":10},"action":"remove","lines":["O"],"id":5195}],[{"start":{"row":321,"column":9},"end":{"row":321,"column":10},"action":"insert","lines":["E"],"id":5196}],[{"start":{"row":321,"column":10},"end":{"row":321,"column":11},"action":"insert","lines":["d"],"id":5197}],[{"start":{"row":321,"column":11},"end":{"row":321,"column":12},"action":"insert","lines":["i"],"id":5198}],[{"start":{"row":321,"column":12},"end":{"row":321,"column":13},"action":"insert","lines":["t"],"id":5199}],[{"start":{"row":321,"column":13},"end":{"row":321,"column":14},"action":"insert","lines":["e"],"id":5200}],[{"start":{"row":321,"column":14},"end":{"row":321,"column":15},"action":"insert","lines":["d"],"id":5201}]]},"ace":{"folds":[{"start":{"row":26,"column":41},"end":{"row":40,"column":0},"placeholder":"..."},{"start":{"row":57,"column":57},"end":{"row":107,"column":0},"placeholder":"..."},{"start":{"row":109,"column":67},"end":{"row":132,"column":0},"placeholder":"..."},{"start":{"row":134,"column":55},"end":{"row":142,"column":0},"placeholder":"..."},{"start":{"row":144,"column":55},"end":{"row":154,"column":0},"placeholder":"..."},{"start":{"row":156,"column":84},"end":{"row":166,"column":0},"placeholder":"..."},{"start":{"row":225,"column":72},"end":{"row":267,"column":0},"placeholder":"..."},{"start":{"row":368,"column":60},"end":{"row":399,"column":0},"placeholder":"..."}],"scrolltop":523,"scrollleft":0,"selection":{"start":{"row":328,"column":33},"end":{"row":328,"column":33},"isBackwards":false},"options":{"guessTabSize":true,"useWrapMode":false,"wrapToView":true},"firstLineState":{"row":50,"state":"start","mode":"ace/mode/golang"}},"timestamp":1438134812000,"hash":"f1e83d3ac2efd50b6c36a1573f7da3dabd590d95"}