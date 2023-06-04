package base

func GenStartHtml() (result string) {
	result = `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<style>
		@font-face { font-family: "FZFangSong-Z02"; src:local("FZFangSong-Z02"), url("https://imgcdn.umiwi.com/ttf/fangzhengfangsong_gbk.ttf"); }
		@font-face { font-family: "FZKai-Z03"; src:local("FZKai-Z03"), url("https://imgcdn.umiwi.com/ttf/fangzhengkaiti_gbk.ttf"); }
		@font-face { font-family: "PingFang SC"; src:local("PingFang SC"); }
		@font-face { font-family: "DeDaoJinKai"; src:local("DeDaoJinKai"), url("https://imgcdn.umiwi.com/ttf/dedaojinkaiw03.ttf");}
		@font-face { font-family: "Source Code Pro"; src:local("Source Code Pro"), url("https://imgcdn.umiwi.com/ttf/0315911806889993935644188722660020367983.ttf"); }
		table, tr, td, th, tbody, thead, tfoot {page-break-inside: avoid !important;}
		img { page-break-inside: avoid; max-width: 100% !important;}
		* { font-family: Helvetica, Arial, sans-serif; font-size: 100%; line-height:1.4; }

	</style>
</head>
<body>
`
	return
}

func GenEndHtml() (result string) {
	result += `
</body>
</html>`
	return
}

func GenPageBreak() (result string) {
	result = `<p style="page-break-after: always;">`
	return
}

func GenHLevelHtml(level int, startTag bool) (result string) {
	sTag := map[int]string{0: `<h1>`, 1: `<h2>`, 2: `<h3>`, 3: `<h4>`, 4: `<h5>`, 5: `<h6>`}
	eTag := map[int]string{0: `</h1>`, 1: `</h2>`, 2: `</h3>`, 3: `</h4>`, 4: `</h5>`, 5: `</h6>`}
	if startTag {
		if tag, ok := sTag[level]; ok {
			result = tag
		}
	} else {
		if tag, ok := eTag[level]; ok {
			result = tag
		}
	}
	return
}
