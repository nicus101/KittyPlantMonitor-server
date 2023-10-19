package inmem

import (
	"fmt"
	"net/url"

	"golang.org/x/exp/slices"
)

func ListFromTable(tableName string, filters url.Values) []map[string]any {
	table, exists := tables[tableName]
	if !exists {
		return nil
	}

	var rows []map[string]any

findRow:
	for _, row := range table {
		for field, values := range filters {
			column, exists := row[field]
			if !exists {
				continue findRow
			}
			columnString := fmt.Sprint(column)
			if !slices.Contains(values, columnString) {
				continue findRow
			}
		}
		rows = append(rows, row)
	}

	return rows
}
