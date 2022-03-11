package application

import "testing"

func TestValidatorShouldError(t *testing.T) {
	type errorTestCases struct {
		description  string
		inputA       string
		inputB       string
		errorMessage string
	}

	for _, escenarion := range []errorTestCases{
		{
			description:  "parametros vacios",
			inputA:       "",
			inputB:       "",
			errorMessage: "Expecter error != nil if parameters are invalid",
		},
		{
			description:  "directorio invalido",
			inputA:       "/home/kk",
			inputB:       "file.json",
			errorMessage: "Expecter error != nil if parameters are invalid",
		},
		{
			description:  "directorio sin permisos",
			inputA:       "/mnt/",
			inputB:       "file.json",
			errorMessage: "Expecter error != nil if parameters are invalid",
		},
	} {
		t.Run(escenarion.description, func(t *testing.T) {
			_, err := NewJsonRepo(escenarion.inputA, escenarion.inputB)

			if err == nil {
				t.Fatalf(escenarion.errorMessage, err)
			}
		})
	}
}

func TestValidatorShouldNotError(t *testing.T) {
	type errorTestCases struct {
		description  string
		inputA       string
		inputB       string
		errorMessage string
	}

	for _, escenarion := range []errorTestCases{

		{
			description:  "directorio valido",
			inputA:       "/tmp/",
			inputB:       "file.json",
			errorMessage: "Valid folder, error not expected",
		},
	} {
		t.Run(escenarion.description, func(t *testing.T) {
			_, err := NewJsonRepo(escenarion.inputA, escenarion.inputB)

			if err != nil {
				t.Fatalf(escenarion.errorMessage)
			}
		})
	}
}
