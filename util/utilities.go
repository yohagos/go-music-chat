package util

import (
	"bufio"
	"os"
	"time"
)

func CreateTimestamp() string {
	currentTIme := time.Now()
	return currentTIme.Format("2006-01-02 15:04:05")
}

func ReadInDatabaseTables() []string {
	f, err := os.Open("./database/database.tables.txt")
	CheckErr(err)

	var result []string

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		result = append(result, sc.Text())
	}
	return result
}