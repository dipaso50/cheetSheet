package domain

type SheetRepo interface {
	Exists(name string) bool
	Save(sheet Sheet)
	Delete(sheet Sheet)
	Get(name string) []Sheet
}
