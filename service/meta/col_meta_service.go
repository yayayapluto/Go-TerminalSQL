package meta

import (
	"fmt"
	"gorm.io/gorm"
	"terminalSQL/types"
)

func AddColumn(db *gorm.DB, tblName string, colName string, datatype string) {
	query := fmt.Sprintf("ALTER TABLE `%v` ADD COLUMN `%v` %v;", tblName, colName, datatype)
	result := db.Debug().Exec(query)

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	fmt.Printf("✅ Successfully added new column '%v' on table '%v'\n", colName, tblName)
}

func RenameColumn(db *gorm.DB, tblName string, oldColName string, newColName string) {
	var query string

	var colsInfo []types.ColumnInfo
	query = fmt.Sprintf("SHOW COLUMNS FROM %v", tblName)
	step1 := db.Raw(query).Scan(&colsInfo)
	if step1.Error != nil {
		fmt.Println("❌ Error occurred: ", step1.Error)
		return
	}

	var currentColInfo types.ColumnInfo
	var currentColIndex int
	for index, col := range colsInfo {
		if col.Field == oldColName {
			currentColIndex = index
			currentColInfo = col
			break
		}
	}
	if (currentColInfo == types.ColumnInfo{}) {
		fmt.Println("❌ Column not found", oldColName)
		return
	}

	var colBefore string
	colBefore = colsInfo[currentColIndex-1].Field

	nullability := "NOT NULL"
	if currentColInfo.Null == "YES" {
		nullability = "NULL"
	}

	defaultValue := ""
	if currentColInfo.Default != nil {
		defaultValue = "DEFAULT " + *currentColInfo.Default
	}

	query = fmt.Sprintf("ALTER TABLE `%v` CHANGE COLUMN `%v` `%v` %v %v %v %v AFTER %v", tblName, oldColName, newColName, currentColInfo.Type, nullability, currentColInfo.Key, defaultValue, colBefore)
	result := db.Debug().Exec(query)

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	fmt.Printf("✅ Successfully renamed column '%v' to '%v' from table '%v'\n", oldColName, newColName, tblName)
}

func DropColumn(db *gorm.DB, tblName string, colName string) {
	result := db.Exec(fmt.Sprintf("ALTER TABLE %v DROP COLUMN %v", tblName, colName))

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	fmt.Printf("✅ Successfully dropped column '%v' from table '%v'\n", colName, tblName)
}
