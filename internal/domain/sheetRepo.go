package domain

type SheetRepo interface {
	ExistsCommand(name string) bool
	ExistsID(name string) bool
	Save(sheet Sheet)
	Delete(ID string)
	Get(name string) []Sheet
	GetAll() []Sheet
}
