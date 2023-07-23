package models

type Table struct {
	Name       string
	Data       []Column
	RowsNumber int
	Index      int
}
