package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"terminalSQL/service"
	"terminalSQL/service/meta"
)

func StartCLI() {
	reader := bufio.NewReader(os.Stdin)
	var userInput string

	conn := service.NewConnService()
	service.DefaultConnection(conn)

	var activeDb *string
	for {

		inputf := "terminalSQL> "
		if activeDb != nil && strings.TrimSpace(*activeDb) != "" {
			inputf = fmt.Sprintf("terminalSQL | %v> ", *activeDb)
		}

		fmt.Printf(inputf)
		userInput, _ = reader.ReadString('\n')
		userQuery := strings.TrimSpace(userInput)

		switch {
		case strings.HasPrefix(userQuery, "help"):
			fmt.Println(`
Database Commands:
  dbs                        	→ show all databases
  mkdb   <name>              	→ create new database
  rmdb   <name>              	→ drop/delete database
  rnmdb  <old> <new>         	→ rename database
  use    <name>              	→ use/switch to database

Table Commands (requires active database):
  tbls                       	→ show all tables
  desc   <table>             	→ describe table structure
  mktbl  <name>              	→ create table
  rmtbl  <name>              	→ drop/delete table
  rnmtbl <old> <new>         	→ rename table
  addcol <table> <col> <type>	→ add column
  rnmcol <table> <old> <new> 	→ rename column
  rmcol  <table> <column>    	→ remove column

Other:
  help                       	→ show this help menu
  exit                       	→ close the program
`)

		case strings.HasPrefix(userQuery, "dbs"):
			meta.ShowDatabases(conn.DB)
		case strings.HasPrefix(userQuery, "mkdb"):
			meta.CreateDatabase(conn.DB, strings.TrimSpace(strings.TrimPrefix(userQuery, "mkdb")))
		case strings.HasPrefix(userQuery, "rmdb"):
			meta.DropDatabase(conn.DB, strings.TrimSpace(strings.TrimPrefix(userQuery, "rmdb")))
		case strings.HasPrefix(userQuery, "rnmdb"):
			inputf := strings.Fields(strings.TrimSpace(strings.TrimPrefix(userQuery, "rnmdb")))
			if len(inputf) >= 2 {
				meta.RenameDatabase(conn.DB, inputf[0], inputf[1])
			} else {
				fmt.Println("❌ Usage: rnmdb <old> <new>")
			}
		case strings.HasPrefix(userQuery, "use"):
			dbName := strings.TrimSpace(strings.TrimPrefix(userQuery, "use"))
			if strings.TrimSpace(dbName) == "" {
				continue
			}
			activeDb = meta.UseDatabase(conn.DB, dbName)
			if activeDb == nil {
				continue
			}
		case strings.HasPrefix(userQuery, "tbls"):
			if activeDb == nil {
				fmt.Println("❌ There is no active db")
				continue
			}
			meta.ShowTables(conn.DB)
		case strings.HasPrefix(userQuery, "desc"):
			if activeDb == nil {
				fmt.Println("❌ There is no active db")
				continue
			}
			meta.DescTable(conn.DB, strings.TrimSpace(strings.TrimPrefix(userQuery, "desc")))
		case strings.HasPrefix(userQuery, "mktbl"):
			if activeDb == nil {
				fmt.Println("❌ There is no active db")
				continue
			}
			meta.CreateTable(conn.DB, strings.TrimSpace(strings.TrimPrefix(userQuery, "mktbl")))
		case strings.HasPrefix(userQuery, "rnmtbl"):
			if activeDb == nil {
				fmt.Println("❌ There is no active db")
				continue
			}
			inputf := strings.Fields(strings.TrimSpace(strings.TrimPrefix(userQuery, "rnmtbl")))
			if len(inputf) >= 2 {
				meta.RenameTable(conn.DB, inputf[0], inputf[1])
			} else {
				fmt.Println("❌ Usage: rnmtbl <old> <new>")
			}
		case strings.HasPrefix(userQuery, "rmtbl"):
			if activeDb == nil {
				fmt.Println("❌ There is no active db")
				continue
			}
			meta.DropTable(conn.DB, strings.TrimSpace(strings.TrimPrefix(userQuery, "rmtbl")))
		case strings.HasPrefix(userQuery, "addcol"):
			if activeDb == nil {
				fmt.Println("❌ There is no active db")
				continue
			}
			inputf := strings.Fields(strings.TrimSpace(strings.TrimPrefix(userQuery, "addcol")))
			if len(inputf) >= 3 {
				tblName := inputf[0]
				colName := inputf[1]
				colType := inputf[2]
				meta.AddColumn(conn.DB, tblName, colName, colType)
			} else {
				fmt.Println("❌ Usage: addcol <table> <column> <type>")
			}
		case strings.HasPrefix(userQuery, "rnmcol"):
			if activeDb == nil {
				fmt.Println("❌ There is no active db")
				continue
			}
			inputf := strings.Fields(strings.TrimSpace(strings.TrimPrefix(userQuery, "rnmcol")))
			if len(inputf) >= 3 {
				tblName := inputf[0]
				oldColName := inputf[1]
				newColName := inputf[2]
				meta.RenameColumn(conn.DB, tblName, oldColName, newColName)
			} else {
				fmt.Println("❌ Usage: rnmcol <table> <old> <new>")
			}
		case strings.HasPrefix(userQuery, "rmcol"):
			if activeDb == nil {
				fmt.Println("❌ There is no active db")
				continue
			}
			inputf := strings.Fields(strings.TrimSpace(strings.TrimPrefix(userQuery, "rmcol")))
			if len(inputf) >= 2 {
				tblName := inputf[0]
				colName := inputf[1]
				meta.DropColumn(conn.DB, tblName, colName)
			} else {
				fmt.Println("❌ Usage: rmcol <table> <column>")
			}
		case strings.HasPrefix(userQuery, "exit"):
			fmt.Println("ugh ok bye.")
			return
		}
	}
}
