package utils

import (
	"io/ioutil"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type ReadFile struct {
	path string
	file string
}

func (r *ReadFile) Run(path string) string {
	var log = logger.New()

	if path == "" {
		return ""
	}

	r.path = path

	readedFile, err := ioutil.ReadFile(r.path)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível ler o arquivo %v", r.path)
		return ""
	}

	r.file = string(readedFile)

	return r.file

}
