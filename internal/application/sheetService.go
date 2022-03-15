package application

import (
	"dipaso/cs/internal/domain"
	"fmt"
)

type SheetService struct {
	idGenerator domain.IDGenerator
	sheetRepo   domain.SheetRepo
}

func NewSheetService(idgenerator domain.IDGenerator, sheetRepo domain.SheetRepo) SheetService {
	return SheetService{idGenerator: idgenerator, sheetRepo: sheetRepo}
}

func (service SheetService) CreateSheet(command, description string) string {

	newID := service.idGenerator.GenerateID()

	newSheet := domain.Sheet{
		ID:          newID,
		Command:     command,
		Description: description,
	}

	service.sheetRepo.Save(newSheet)

	return newID
}

func (service SheetService) DeleteSheet(ID string) error {
	if !service.sheetRepo.ExistsID(ID) {
		return fmt.Errorf("%s not found", ID)
	}

	service.sheetRepo.Delete(ID)
	return nil
}

func (service SheetService) List(command string) []domain.Sheet {
	return service.sheetRepo.Get(command)
}

func (service SheetService) ListAll() []domain.Sheet {
	return service.sheetRepo.GetAll()
}
