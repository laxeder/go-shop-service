package archive

import (
	"fmt"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

func NewPDF(template string) (pdf string, err error) {

	var log = logger.New()

	pdf = ""
	err = nil

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível construir uma contexto para o arquivo. (%v)", template)
		return
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(template))
	page.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(page)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	err = pdfg.Create()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível gerar uma arquivo PDF. (%v)", template)
		return
	}

	pdf = string(pdfg.Bytes())

	file := fmt.Sprintf("./temp-pdf/%s.pdf", date.NowUTC())
	err = Write(pdfg, file)

	return

}

func Write(pdfg *wkhtmltopdf.PDFGenerator, pathName string) (err error) {

	var log = logger.New()

	err = nil

	err = pdfg.WriteFile(pathName)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível salvar em disco um arquivo PDF. (%v)", pathName)
		return
	}

	return

}

func maxSize1MB(base64 string) bool {
	maxBytes := 1024 ^ 2
	bytes := size(base64)
	return bytes > maxBytes

}

func size(base64 string) int {
	return len(base64) * (4 / 3)
}
