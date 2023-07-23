package utils

import (
	"houx/models"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseTable(page string) []models.Table {

	document, err := goquery.NewDocumentFromReader(strings.NewReader(page))
	if err != nil {
		panic(err)
	}
	column := make([]models.Column, 0)
	row := make([]models.Row, 0)
	var columnIndexes = 0
	var rowNumber = 0
	var tableNumber = 0
	var currentTableRowsNumber = 0
	var aux = 0
	var rowAux = 0
	var isHeaderLess = false
	filteredRow := make([]models.Row, 0)
	table := make([]models.Table, 0)
	columns := make([]string, 0)
	document.Find("table").Each(func(i int, item *goquery.Selection) {
		rowNumber = item.Find("tr").Size() - item.Find("thead tr").Size()
		item.Find("th").Each(func(i int, th *goquery.Selection) {
			columns = append(columns, th.Text())
		})

		isHeaderLess = item.Find("thead").Size() == 0
		for i, currentColumn := range columns {
			column = append(column, models.Column{Name: currentColumn, Index: i})

		}
		columnIndexes = len(columns) - 1
		table = append(table, models.Table{Name: "", RowsNumber: columnIndexes * rowNumber, Index: tableNumber})
		columns = make([]string, 0)
		item.Find("tr").Each(func(i int, tr *goquery.Selection) {
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				if aux <= columnIndexes {
					row = append(row, models.Row{Index: aux, Data: td.Text()})
				} else {
					aux = 0
					row = append(row, models.Row{Index: aux, Data: td.Text()})
				}
				aux++
			})
			if isHeaderLess && (rowNumber-1) == i {
				i++
			}
			if rowNumber == i {
				sort.Slice(column, func(i2, j int) bool {
					return column[i2].Index < column[j].Index
				})

				sort.Slice(row, func(i, j int) bool {
					return row[i].Index < row[j].Index
				})
				for i, curColumn := range column {
					currentTableRowsNumber = len(row)
					for i = rowAux; i < currentTableRowsNumber; i++ {
						if curColumn.Index == row[i].Index || i == currentTableRowsNumber-1 {
							if i == currentTableRowsNumber {
								column[curColumn.Index].Data = filteredRow
								filteredRow = make([]models.Row, 0)
								rowAux = i
							} else {
								filteredRow = append(filteredRow, models.Row{Index: curColumn.Index, Data: row[i].Data})
								if i == currentTableRowsNumber-1 {
									column[curColumn.Index].Data = filteredRow
									filteredRow = make([]models.Row, 0)
									rowAux = i
								}
							}

						} else {
							column[curColumn.Index].Data = filteredRow
							filteredRow = make([]models.Row, 0)
							rowAux = i
							break
						}

					}

				}

				table[tableNumber].Data = column
				tableNumber++
				rowAux = 0
				column = make([]models.Column, 0)
				row = make([]models.Row, 0)
			}

		})

	})

	return table

}
