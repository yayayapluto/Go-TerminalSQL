package meta

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"terminalSQL/types"
	"terminalSQL/utils"
)

func ShowTables(db *gorm.DB) {
	var tables []string

	dbQuery := fmt.Sprintf("SHOW TABLES")
	result := db.Raw(dbQuery).Scan(&tables)
	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	if len(tables) <= 0 {
		fmt.Println("❔ There is no tables")
		return
	}

	maxLength := 0
	for _, dbname := range tables {
		if len(dbname) > maxLength {
			maxLength = len(dbname)
		}
	}
	padding := maxLength + 2

	utils.Separator("-", padding)
	utils.PrintField(&padding, "Tables")
	utils.Separator("-", padding)
	for _, dbname := range tables {
		utils.PrintField(&padding, dbname)
	}
	utils.Separator("-", padding)
}

func DescTable(db *gorm.DB, tblName string) {
	var cols []types.ColumnInfo

	result := db.Raw("DESCRIBE " + tblName).Scan(&cols)
	if result.Error != nil {
		fmt.Println("❌ Error occurred:", result.Error)
		return
	}

	fieldLen := len("Field")
	typeLen := len("Type")
	nullLen := len("Null")
	keyLen := len("Key")
	defaultLen := len("Default")
	extraLen := len("Extra")

	for _, col := range cols {
		if len(col.Field) > fieldLen {
			fieldLen = len(col.Field)
		}
		if len(col.Type) > typeLen {
			typeLen = len(col.Type)
		}
		if len(col.Null) > nullLen {
			nullLen = len(col.Null)
		}
		if len(col.Key) > keyLen {
			keyLen = len(col.Key)
		}
		if col.Default != nil && len(*col.Default) > defaultLen {
			defaultLen = len(*col.Default)
		}
		if len(col.Extra) > extraLen {
			extraLen = len(col.Extra)
		}
	}

	totalWidth := fieldLen + typeLen + nullLen + keyLen + defaultLen + extraLen + 19
	separator := strings.Repeat("-", totalWidth)

	fmt.Println(separator)
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		fieldLen, "Field", typeLen, "Type", nullLen, "Null",
		keyLen, "Key", defaultLen, "Default", extraLen, "Extra")
	fmt.Println(separator)

	for _, col := range cols {
		defaultVal := ""
		if col.Default != nil {
			defaultVal = *col.Default
		}
		fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
			fieldLen, col.Field,
			typeLen, col.Type,
			nullLen, col.Null,
			keyLen, col.Key,
			defaultLen, defaultVal,
			extraLen, col.Extra)
	}

	fmt.Println(separator)
}

func CreateTable(db *gorm.DB, tblName string) {
	queryf := fmt.Sprintf("CREATE TABLE %v (id INT PRIMARY KEY, created_at DATETIME NULL DEFAULT NOW(), updated_at DATETIME NULL DEFAULT NOW() ON UPDATE NOW())", tblName)
	result := db.Exec(queryf)

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	fmt.Printf("✅ Successfully added new table '%v'\n", tblName)
}

func RenameTable(db *gorm.DB, oldTblName string, newTblName string) {
	queryf := fmt.Sprintf("ALTER TABLE %v RENAME TO %v", oldTblName, newTblName)
	result := db.Exec(queryf)

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	fmt.Printf("✅ Successfully renames table '%v' to '%v'\n", oldTblName, newTblName)
}

func DropTable(db *gorm.DB, tblName string) {
	result := db.Exec("DROP TABLE " + tblName)

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	fmt.Printf("✅ Successfully dropped table '%v'\n", tblName)
}
