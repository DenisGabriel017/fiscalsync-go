package importfile

import "errors"

func NormalizarDocumento(entrada string) (string, error) {

	if entrada == "" {
		return "", errors.New("documento obrigatório")
	}

	resultado := ""

	for _, caractere := range entrada {
		if caractere >= '0' && caractere <= '9' {
			resultado += string(caractere)

		}
	}

	if resultado == "" {
		return "", errors.New("documento sem numero")
	}
	return resultado, nil
}
