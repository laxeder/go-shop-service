package tpl

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

// Gera o template
func New(data map[string]interface{}, pathName string) (tpl string, err error) {

	var log = logger.New()

	tpl = ""
	err = nil

	tpl, err = Read(pathName)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar ler arquivo no disco. (%v)", pathName)
		return
	}

	// Preenche o template com dados fornecido
	for key, value := range data {

		variableKey := fmt.Sprintf("{{.%s}}", key)
		variabelValue := fmt.Sprintf("%s", value)

		tpl = strings.Replace(tpl, variableKey, variabelValue, 1)
	}

	return
}

func Read(pathName string) (tpl string, err error) {

	var log = logger.New()

	tpl = ""
	err = nil

	_, base, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(base))
	rootDir := filepath.Dir(dir)

	pathDir := path.Join(rootDir, pathName)
	tplBytes, err := ioutil.ReadFile(pathDir)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar ler arquivo no disco. (%v)", pathName)
		return
	}

	tpl = string(tplBytes)

	return
}
