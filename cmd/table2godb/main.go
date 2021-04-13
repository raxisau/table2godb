package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" // Standard way of importng mysql driver
)

var (
	dataSourceName = flag.String("ds", "", "Data source for the database username:password@tcp(127.0.0.1:3306)/dbname")
	tableName      = flag.String("t", "", "Name of the table to code up")
)

func main() {
	flag.Parse()

	driverName := "mysql"

	tblField := sql.NullString{}
	tblType := sql.NullString{}
	tblNull := sql.NullString{}
	tblKey := sql.NullString{}
	tblDefault := sql.NullString{}
	tblExtra := sql.NullString{}

	db, err := sql.Open(driverName, *dataSourceName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	results, err := db.Query("DESCRIBE " + *tableName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tName := colToGo(*tableName)

	fmt.Println("//", tName, "A row in the table", *tableName)
	fmt.Println("type", tName, "struct {")
	for results.Next() {

		err = results.Scan(&tblField, &tblType, &tblNull, &tblKey, &tblDefault, &tblExtra)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("    ", colToGo(tblField.String), "    ", typeToGo(tblType.String), fmt.Sprintf("`db:\"%s%s%s\"`", tblField.String, keyToGo(tblKey.String), extraToGo(tblExtra.String)))

	}
	fmt.Println("}")
	fmt.Println("// TableName name of the table")
	fmt.Printf("func (*%s) TableName() string {\n", tName)
	fmt.Printf("	return \"%s\"\n", *tableName)
	fmt.Println("}")

}

func keyToGo(columnKey string) string {
	if columnKey == "PRI" {
		return ",key"
	}
	return ""
}

func extraToGo(columnExtra string) string {
	if columnExtra == "auto_increment" {
		return ",auto"
	}
	return ""
}

func colToGo(columnName string) string {
	if columnName == "id" {
		return "ID"
	}

	components := strings.Split(columnName, "_")
	for i, element := range components {
		if element == "id" {
			components[i] = "ID"
		} else {
			components[i] = strings.ToUpper(element[0:1]) + element[1:]
		}
	}
	return strings.TrimPrefix(strings.Join(components, ""), "Fld")
}

func typeToGo(tblType string) string {
	if strings.Contains(tblType, "int") {
		return "int64"
	} else if strings.Contains(tblType, "time") || strings.Contains(tblType, "date ") {
		return "time.Time"
	} else if strings.Contains(tblType, "double") || strings.Contains(tblType, "decimal ") {
		return "float64"
	}
	return "string"
}
