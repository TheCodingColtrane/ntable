package utils

import (
	"strings"
)

func GetTableData(requestedBody string) [][]string {
	hasMultipleTables := strings.Count(requestedBody, "table") > 0
	if hasMultipleTables {
		tables := strings.Split(requestedBody, "table")
		var incompleteExtractedTable [][]string
		var completeExtractedTable [][]string

		for i, curTable := range tables {
			if i > 0 {
				// if strings.Contains(curTable, "thead") {
				incompleteExtractedTable = CreateTable(curTable, true)
				// } else {
				// 	incompleteExtractedTable = CreateHeadLessTable(curTable, true)

				// }
				for i = 0; i < len(incompleteExtractedTable); i++ {
					completeExtractedTable = append(completeExtractedTable, incompleteExtractedTable[i])
				}
			}
		}

		return completeExtractedTable

	}

	return CreateTable(requestedBody, true)

}

func CreateTable(requestBody string, isMultipleTables bool) [][]string {
	var tableBody string = ""
	var tableHeads []string
	tableBody = CreateTableSchema(requestBody, "table", "</table>")
	tableBody = strings.Replace(tableBody, "\n", "", -1)
	tableTh := CreateTableSchema(tableBody, "thead", "</thead>")
	tableTrRow := CreateTableSchema(tableTh, "tr", "</tr>")
	thDescription := strings.Split(strings.Replace(tableTrRow, "th", "", -1), "<>")
	if len(thDescription) == 1 {
		tableHeads = make([]string, 1)

	} else {
		tableHeads = make([]string, len(thDescription)-1)

	}
	for i, description := range thDescription {
		if strings.Contains(description, "</>") {
			if i == 0 {
				tableHeads[i] = strings.Replace(description, "</>", "", -1)

			} else {
				tableHeads[i-1] = strings.Replace(description, "</>", "", -1)
			}
		}
	}
	body := CreateTableSchema(tableBody, "tbody", "</tbody>")
	rows := make([]string, strings.Count(body, "</tr>")*len(tableHeads))
	table := make([][]string, len(tableHeads))
	var rowNumber int = 0
	var columnNumber int = 0
	for i := 0; i < len(tableHeads); i++ {
		table[i] = make([]string, len(rows)/len(tableHeads))
		table[columnNumber][0] = tableHeads[i]
		columnNumber++
	}
	columnNumber = 0
	decoupledRows := strings.Split(strings.Replace(body, "tr", "", -1), "<>")
	for _, row := range decoupledRows {
		rowsData := strings.Split(string(row), "td")
		for i, data := range rowsData {
			lastCharacterNumber := strings.LastIndex(data, "</")
			if strings.Contains(data, ">") && len(data) > 1 && i > 0 && string(data[1]) != "<" && lastCharacterNumber > -1 {
				rows[rowNumber] = data[1:lastCharacterNumber]
				if rowNumber <= 2 {
					table[columnNumber][rowNumber] = data[1:lastCharacterNumber]
				} else {
					if columnNumber == 0 {
						table[columnNumber][GetEmptyArrayIndex(table[columnNumber])] = data[1:lastCharacterNumber]

					} else if columnNumber == 1 {
						table[columnNumber][GetEmptyArrayIndex(table[columnNumber])] = data[1:lastCharacterNumber]

					} else {
						table[columnNumber][GetEmptyArrayIndex(table[columnNumber])] = data[1:lastCharacterNumber]
					}
				}
				rowNumber++
				columnNumber++
			}
			if columnNumber > len(tableHeads)-1 {
				columnNumber = 0
			}
		}
	}
	return table
}

func CreateTableSchema(tableBody string, firstDelimiter string, secondDelimiter string) string {
	startStringPosition := strings.Index(tableBody, firstDelimiter)
	endStringPosition := strings.Index(tableBody, secondDelimiter)
	schema := tableBody[startStringPosition:endStringPosition]
	return schema
}

func GetEmptyArrayIndex(rows []string) int {
	var i int
	for i, row := range rows {
		if row == "" {
			return i
		}
	}
	return i
}

func CreateHeadLessTable(requestBody string, isMultipleTables bool) [][]string {
	var tableBody string = ""
	if !isMultipleTables {
		tableBody = CreateTableSchema(requestBody, "table", "</table>")
		tableBody = strings.Replace(tableBody, "\n", "", -1)

	} else {
		tableBody = strings.Replace(requestBody, "\n", "", -1)
	}
	body := CreateTableSchema(tableBody, "tbody", "</tbody>")
	rows := make([]string, strings.Count(body, "</tr>")*2)
	table := make([][]string, len(rows))
	var rowNumber int = 0
	var columnNumber int = 0
	for i := 0; i < len(rows); i++ {
		table[i] = make([]string, len(rows)/len(rows))
		table[columnNumber][0] = rows[i]
		columnNumber++
	}
	columnNumber = 0
	decoupledRows := strings.Split(strings.Replace(body, "tr", "", -1), "<>")
	for _, row := range decoupledRows {
		rowsData := strings.Split(string(row), "td")
		for i, data := range rowsData {
			lastCharacterNumber := strings.LastIndex(data, "</")
			if strings.Contains(data, ">") && len(data) > 1 && i > 0 && string(data[1]) != "<" && lastCharacterNumber > -1 {
				rows[rowNumber] = data[1:lastCharacterNumber]
				if rowNumber <= 2 {
					table[columnNumber][rowNumber] = data[1:lastCharacterNumber]
				} else {
					if columnNumber == 0 {
						table[columnNumber][GetEmptyArrayIndex(table[columnNumber])] = data[1:lastCharacterNumber]

					} else if columnNumber == 1 {
						table[columnNumber][GetEmptyArrayIndex(table[columnNumber])] = data[1:lastCharacterNumber]

					} else {
						table[columnNumber][GetEmptyArrayIndex(table[columnNumber])] = data[1:lastCharacterNumber]
					}
				}
				rowNumber++
				columnNumber++
			}
			if columnNumber > len(rows)-1 {
				columnNumber = 0
			}
		}
	}
	return table
}

// func GetAnyTable(rawTable string) {
// 	tableBody := strings.Replace(rawTable, "\n", "", -1)
// 	var thead []string
// 	if strings.Contains(rawTable, "<thead") {
// 		thead = DecoupleTheadElements(CreateTableSchema(tableBody, "<thead", "</thead>"))
// 	} else {

// 	}

// }

func DecoupleTheadElements(thead string) []string {
	hasRows := strings.Contains(thead, "<tr")
	var theadRows = ""
	var theadColumns = ""
	var tableHeads []string
	if hasRows {
		theadRows = CreateTableSchema(thead, "<tr", "</tr>")
		theadColumns = CreateTableSchema(theadRows, "<th", "</th>")
	} else {
		theadColumns = CreateTableSchema(thead, "<th", "</th>")
	}
	columns := strings.Split(strings.Replace(theadColumns, "th", "", -1), "<>")
	if len(columns) == 1 {
		tableHeads = make([]string, 1)

	} else {
		tableHeads = make([]string, len(columns)-1)

	}
	for i, theadElement := range columns {
		if strings.Contains(theadElement, "</>") {
			if i == 0 {
				tableHeads[i] = strings.Replace(theadElement, "</>", "", -1)

			} else {
				tableHeads[i-1] = strings.Replace(theadElement, "</>", "", -1)
			}
		}
	}

	return tableHeads
}

// func DecoupledRowsElements(table string){
// 	hasRows := strings.Contains(table, "<tr")
// 	var theadRows = ""
// 	var theadColumns = ""
// 	var tableHeads []string
// 	if hasRows {
// 		theadRows = CreateTableSchema(table, "<tr", "</tr>")
// 		hasRowsColumns := strings.Contains(theadRows, "<th")
// 		if hasRowsColumns {
// 			theadColumns = CreateTableSchema(table, "<th", "</th>")
// 			thDescription := strings.Split(strings.Replace(theadColumns, "th", "", -1), "<>")
// 			if len(thDescription) == 1 {
// 				tableHeads = make([]string, 1)

// 			} else {
// 				tableHeads = make([]string, len(thDescription)-1)
// 			}
// 		} else {

// 		}
// 		theadColumns = CreateTableSchema(theadRows, "<th", "</th>")
// 	} else {
// 		theadColumns = CreateTableSchema(thead, "<th", "</th>")
// 		thDescription := strings.Split(strings.Replace(tableTrRow, "th", "", -1), "<>")

// 	}
// }
