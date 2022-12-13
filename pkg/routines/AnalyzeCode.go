package routines

import (
	"os/exec"
)

func AnalyzeCode() {

	go func() {
		cmd := "go vet ."
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Error().Err(err).Msgf("Erro ao tentar analizar o código (go vet)")
		}

		log.Info().Msgf("\n[GO VET] -------------------------------------------- \n %v", string(out))
	}()

	go func() {
		cmd := "golint ."
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Error().Err(err).Msgf("Erro do Linter na estrutura do código. (golint .)")
		}

		log.Info().Msgf("\n[GOLINT] -------------------------------------------- \n %v", string(out))
	}()

	go func() {
		cmd := "goimports"
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Error().Err(err).Msgf("Erro nas importações do código. (goimports)")
		}

		log.Info().Msgf("\n[GOIMPORTS] -------------------------------------------- \n %v", string(out))
	}()

	go func() {
		cmd := "gocritic check ."
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Error().Err(err).Msgf("Erro cricito no código. (gocritic check .)")
		}

		log.Info().Msgf("\n[GOCRITIC] -------------------------------------------- \v %v", string(out))
	}()

	go func() {
		cmd := "staticcheck ."
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Error().Err(err).Msgf("Erro erro ")
		}

		log.Info().Msgf("\n[GO STATIC CHECK] -------------------------------------------- \n %v", string(out))
	}()

}
