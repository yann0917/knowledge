package base

import (
	"bytes"
	"fmt"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GenPdf(buf *bytes.Buffer, fileName string) (err error) {
	pdfg, _ := wkhtmltopdf.NewPDFGenerator()

	page := wkhtmltopdf.NewPageReader(buf)
	page.Encoding.Set("utf-8")
	page.FooterFontSize.Set(10)
	page.FooterRight.Set("[page]")
	page.DisableSmartShrinking.Set(true)

	page.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(page)

	pdfg.Cover.EnableLocalFileAccess.Set(true)

	pdfg.Dpi.Set(300)

	pdfg.TOC.Include = true
	pdfg.TOC.TocHeaderText.Set("目 录")
	pdfg.TOC.HeaderFontSize.Set(18)

	pdfg.TOC.TocLevelIndentation.Set(15)
	pdfg.TOC.TocTextSizeShrink.Set(0.9)
	pdfg.TOC.DisableDottedLines.Set(false)
	pdfg.TOC.EnableTocBackLinks.Set(true)

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.MarginTop.Set(15)
	pdfg.MarginBottom.Set(15)
	pdfg.MarginLeft.Set(15)
	pdfg.MarginRight.Set(15)

	err = pdfg.Create()
	if err != nil {
		fmt.Printf("pdfg create err: %#v\n", err)
		return
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(fileName)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return
	}
	fmt.Printf("\033[32;1m%s\033[0m\n", "完成")
	// err = os.Remove(coverPath)
	return
}
