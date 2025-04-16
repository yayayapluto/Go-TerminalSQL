package query

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func ExecRaw(db *gorm.DB, query string) {
	query = strings.TrimSpace(strings.ToUpper(query))
	if strings.Contains(query, "SELECT") {
		var results []map[string]interface{}
		result := db.Raw(query).Scan(&results)
		if result.Error != nil {
			fmt.Println("❌ Error occurred:", result.Error)
			return
		}

		if len(results) == 0 {
			fmt.Println("✅ Query success but no result found.")
			return
		}

		var cols []string
		colWidths := make(map[string]int)

		for col := range results[0] {
			cols = append(cols, col)
			colWidths[col] = len(col)
		}

		for _, row := range results {
			for _, col := range cols {
				val := fmt.Sprintf("%v", row[col])
				if len(val) > colWidths[col] {
					colWidths[col] = len(val)
				}
			}
		}

		totalWidth := 1
		for _, col := range cols {
			totalWidth += colWidths[col] + 3
		}
		separator := strings.Repeat("-", totalWidth)

		fmt.Println(separator)
		fmt.Print("|")
		for _, col := range cols {
			fmt.Printf(" %-*s |", colWidths[col], col)
		}
		fmt.Println()
		fmt.Println(separator)

		for _, row := range results {
			fmt.Print("|")
			for _, col := range cols {
				fmt.Printf(" %-*v |", colWidths[col], row[col])
			}
			fmt.Println()
		}
		fmt.Println(separator)

	} else {
		result := db.Debug().Exec(query)
		if result.Error != nil {
			fmt.Println("❌ Error occurred:", result.Error)
			return
		}
		fmt.Println("✅ Query executed successfully.")
	}
}
