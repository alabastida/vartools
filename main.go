package main

import (
    _ "fmt"
	_ "time"
	_ "strconv"  
    "github.com/tealeg/xlsx"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
)

const SIZE = 1000000

func main() {

	a := 1
	b := "test"
	
    var file *xlsx.File
    var sheet *xlsx.Sheet
    var err error
	
	use(a, file, sheet, err)
}

func use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}