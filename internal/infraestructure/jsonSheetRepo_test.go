package infraestructure

import (
	"dipaso/cs/internal/domain"
	"testing"
)

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

func TestWriteJson(t *testing.T) {
	type errorTestCases struct {
		sheet        domain.Sheet
		errorMessage string
		fileLocation string
		fileName     string
		description  string
	}

	for _, escenarion := range []errorTestCases{

		{
			sheet:        domain.Sheet{Command: "docker", Description: "kakak", ID: "lala"},
			errorMessage: "Commando must exist into the file",
			description:  "Save sheet into file test",
			fileLocation: "/tmp",
			fileName:     "test.json",
		},
	} {
		t.Run(escenarion.description, func(t *testing.T) {
			repo, _ := NewJsonRepo(escenarion.fileLocation, escenarion.fileName)

			repo.Save(escenarion.sheet)
			exist := repo.ExistsCommand(escenarion.sheet.Command)

			if !exist {
				t.Fatalf(escenarion.errorMessage)
			}

			repo.Delete(escenarion.sheet.ID)
		})
	}
}

func TestRetreiveSheet(t *testing.T) {
	type errorTestCases struct {
		sheet        domain.Sheet
		errorMessage string
		fileLocation string
		fileName     string
		description  string
	}

	for _, escenarion := range []errorTestCases{

		{
			sheet:        domain.Sheet{Command: "docker", Description: "kakak", ID: "lala"},
			errorMessage: "Commando must exist into the file",
			description:  "Save sheet into file test",
			fileLocation: "/tmp",
			fileName:     "test.json",
		},
	} {
		t.Run(escenarion.description, func(t *testing.T) {
			repo, _ := NewJsonRepo(escenarion.fileLocation, escenarion.fileName)

			repo.Save(escenarion.sheet)
			exist := repo.Get(escenarion.sheet.Command)

			if len(exist) == 0 {
				t.Fatalf(escenarion.errorMessage)
			}

			repo.Delete(escenarion.sheet.ID)
		})
	}
}
