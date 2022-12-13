package routines

import (
	"os/exec"
)

func DocSwagger() {

	cmd := "swag init"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar analizar o código")
	}

	log.Info().Msgf("\n[SWAGGER] ------------------- \n	%v", string(out))

}
