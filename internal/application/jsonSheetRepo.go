package application

import (
	"fmt"
	"os"
	"path/filepath"
)

type JsonSheetRepo struct {
	filelocation string
	filename     string
	file         *os.File
}

func NewJsonRepo(fileLocation string, filename string) (*JsonSheetRepo, error) {

	locationExist, err := exists(fileLocation)

	if err != nil {
		return nil, err
	}

	if !locationExist || len(filename) == 0 {
		return nil, fmt.Errorf("parameters can't be empty nor nil (%s,%s)", fileLocation, filename)
	}

	completeFilename := filepath.Join(fileLocation, filename)

	f, err := os.OpenFile(completeFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	return &JsonSheetRepo{filelocation: fileLocation, filename: filename, file: f}, nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**func (jsonRepo *JsonSheetRepo) Exists(name string) bool {

}

func (jsonRepo *JsonSheetRepo) Save(sheet domain.Sheet) {

}

func (jsonRepo *JsonSheetRepo) Delete(sheet domain.Sheet) {

}
func (jsonRepo *JsonSheetRepo) Get(name string) []domain.Sheet {

}**/
