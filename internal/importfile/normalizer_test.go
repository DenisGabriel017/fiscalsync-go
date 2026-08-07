package importfile

import "testing"

func TestNormalizarDocumento(t *testing.T) {
	entrada := "045.456.813-44"
	esperado := "04545681344"

	resultado, erro := NormalizarDocumento(entrada)

	if erro != nil {
		t.Fatalf("não esperava erro: %v", erro)
	}

	if resultado != esperado {
		t.Errorf("esperado %q, recebido %q", esperado, resultado)
	}

}

func TestNormalizarDocumento_EntradaVazia(t *testing.T) {
	entrada := ""

	_, erro := NormalizarDocumento(entrada)

	if erro == nil {
		t.Fatalf("nesperava um erro para entrada vazia")
	}

}

func TestNormalizarDocumento_RemoverLetrasESimbolos(t *testing.T) {
	entrada := "abc123-45"
	esperado := "12345"

	resultado, err := NormalizarDocumento(entrada)

	if err != nil {
		t.Fatalf("erro na execução: %v", err)
	}

	if resultado != esperado {
		t.Errorf("Resultado esperado %q, Resultado recebido %q", esperado, resultado)
	}
}
