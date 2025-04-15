package meta

import (
	"fmt"
	"gorm.io/gorm"
	"terminalSQL/utils"
)

func ShowDatabases(db *gorm.DB) {
	var databases []string

	result := db.Raw("SHOW DATABASES").Scan(&databases)
	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	maxLength := 0
	for _, dbname := range databases {
		if len(dbname) > maxLength {
			maxLength = len(dbname)
		}
	}
	padding := maxLength + 2

	utils.Separator("-", padding)
	utils.PrintField(&padding, "Databases")
	utils.Separator("-", padding)
	for _, dbname := range databases {
		utils.PrintField(&padding, dbname)
	}
	utils.Separator("-", padding)
}

func CreateDatabase(db *gorm.DB, dbname string) {
	queryf := fmt.Sprintf("CREATE DATABASE %v", dbname)
	result := db.Exec(queryf)

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	fmt.Printf("✅ Successfully added new database '%v'\n", dbname)
}

func DropDatabase(db *gorm.DB, dbname string) {
	queryf := fmt.Sprintf("DROP DATABASE %v", dbname)
	result := db.Exec(queryf)

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return
	}

	fmt.Printf("✅ Successfully dropped database '%v'\n", dbname)
}

func RenameDatabase(db *gorm.DB, oldDbname string, newDbName string) {
	step1 := db.Exec("CREATE DATABASE " + newDbName)
	if step1.Error != nil {
		fmt.Println("❌ Error occurred: ", step1.Error, "| step 1")
		return
	}

	step2 := db.Exec("USE " + oldDbname)
	if step2.Error != nil {
		fmt.Println("❌ Error occurred: ", step2.Error, "| step 2")
		return
	}

	var oldDbTables *[]string
	step3 := db.Raw("SHOW TABLES").Scan(&oldDbTables)
	if step3.Error != nil {
		fmt.Println("❌ Error occurred: ", step3.Error, "| step 3")
		return
	}
	if oldDbTables != nil {
		for i, tblName := range *oldDbTables {
			queryf := fmt.Sprintf("CREATE TABLE %v.%v AS SELECT * FROM %v.%v", newDbName, tblName, oldDbname, tblName)
			result := db.Exec(queryf)
			if result.Error != nil {
				fmt.Println("❌ Error occurred: ", result.Error, "| step 3.", i)
				return
			}
		}
	}

	step4 := db.Exec("DROP DATABASE " + oldDbname)
	if step4.Error != nil {
		fmt.Println("❌ Error occurred: ", step4.Error, "| step 4")
		return
	}

	fmt.Printf("✅ Successfully renamed database from '%v' to '%v'\n", oldDbname, newDbName)
}

func UseDatabase(db *gorm.DB, dbname string) *string {
	queryf := fmt.Sprintf("USE %v", dbname)
	result := db.Exec(queryf)

	if result.Error != nil {
		fmt.Println("❌ Error occurred: ", result.Error)
		return nil
	}

	fmt.Printf("✅ Successfully used database '%v'\n", dbname)
	return &dbname
}
