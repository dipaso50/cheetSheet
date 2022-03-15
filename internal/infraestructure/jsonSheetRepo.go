package infraestructure

import (
	"dipaso/cs/internal/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CheetSheets struct {
	CheetSheets []Sheet `json:"Sheets"`
}

type Sheet struct {
	ID          string `json:"ID"`
	Command     string `json:"Command"`
	Description string `json:"Description"`
}

type JsonSheetRepo struct {
	filelocation     string
	filename         string
	completeFilename string
	file             *os.File
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

	return &JsonSheetRepo{filelocation: fileLocation, filename: filename, file: f, completeFilename: completeFilename}, nil
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

func (jsonRepo *JsonSheetRepo) readJson() []Sheet {

	// Open our jsonFile
	jsonFile, err := os.Open(jsonRepo.completeFilename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var allSheets CheetSheets

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &allSheets)

	return allSheets.CheetSheets
}

func (jsonRepo JsonSheetRepo) writeJson(allSheets []Sheet) {

	toWrite := CheetSheets{
		CheetSheets: allSheets,
	}

	file, _ := json.MarshalIndent(toWrite, "", " ")

	if err := os.Truncate(jsonRepo.completeFilename, 0); err != nil {
		fmt.Printf("Failed to truncate: %v", err)
		return
	}

	_ = ioutil.WriteFile(jsonRepo.completeFilename, file, 0644)
}

func (jsonRepo *JsonSheetRepo) ExistsCommand(command string) bool {

	for _, sheet := range jsonRepo.readJson() {
		if sheet.Command == command {
			return true
		}
	}

	return false
}

func (jsonRepo *JsonSheetRepo) ExistsID(id string) bool {

	for _, sheet := range jsonRepo.readJson() {
		if sheet.ID == id {
			return true
		}
	}

	return false
}

func (jsonRepo *JsonSheetRepo) Save(sheet domain.Sheet) {
	allSheets := jsonRepo.readJson()

	newSheet := Sheet{
		ID:          sheet.ID,
		Command:     sheet.Command,
		Description: sheet.Description,
	}

	allSheets = append(allSheets, newSheet)

	jsonRepo.writeJson(allSheets)
}

func RemoveIndex(s []Sheet, index int) []Sheet {
	return append(s[:index], s[index+1:]...)
}

func (jsonRepo *JsonSheetRepo) Delete(ID string) {

	allsheets := jsonRepo.readJson()

	var indexToRemove = 0

	for index, sheet := range allsheets {
		if sheet.ID == ID {
			indexToRemove = index
			break
		}
	}

	allsheets = RemoveIndex(allsheets, indexToRemove)

	jsonRepo.writeJson(allsheets)
}

func (jsonRepo *JsonSheetRepo) Get(command string) []domain.Sheet {
	allsheets := jsonRepo.readJson()

	var results []domain.Sheet

	for _, sheet := range allsheets {
		if sheet.Command == command {
			results = append(results, domain.Sheet{
				ID:          sheet.ID,
				Command:     sheet.Command,
				Description: sheet.Description,
			})
		}
	}

	return results
}

func (jsonRepo *JsonSheetRepo) GetAll() []domain.Sheet {
	allsheets := jsonRepo.readJson()

	var results []domain.Sheet

	for _, sheet := range allsheets {
		results = append(results, domain.Sheet{
			ID:          sheet.ID,
			Command:     sheet.Command,
			Description: sheet.Description,
		})
	}

	return results
}
