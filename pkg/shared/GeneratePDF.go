package shared

import (
	"github.com/laxeder/go-shop-service/pkg/modules/archive"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/tpl"
)

func GeneratePDF(data map[string]interface{}, pathNameTpl string) (pdf string, err error) {

	var log = logger.New()

	pdf = ""
	err = nil

	template, err := tpl.New(data, pathNameTpl)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível carregar o template no path %v.", pathNameTpl)
		return
	}

	pdf, err = archive.NewPDF(template)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível gerar uma PDf com o template (%v).", template)
		return
	}

	return

}
