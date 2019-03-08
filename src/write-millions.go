package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	SqlDatetimeFormat = "2006-01-02 15:04:05"
	// 90M rows should take < 3 min on i7 CPU
	// Insertion into MySQL for 10M rows may take 20 min with 3 indexes
	// 90M rows w/3 indexes takes 3 to 4 hours
	NumRowsToGenerate = 10 * 1000 * 1000
	Delimiter         = "\t"
	RowFormat         = "%v" + Delimiter + "%v" + Delimiter + "%v" + Delimiter + "%v\n"
	InitialPrimaryId  = 1 // + (10 * 1000 * 1000)
)

func main() {
	startTime := time.Now()
	log.Printf("Start: %v", startTime)
	file, e := os.Create("output_table_content.csv")
	if e != nil {
		log.Fatal(e)
	}
	defer func() {
		e := file.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()
	writer := bufio.NewWriter(file)
	// No header row allowed for LOAD DATA INFILE
	for i := InitialPrimaryId; i <= InitialPrimaryId+NumRowsToGenerate-1; i++ {
		// % 10,000 to prevent going out of range of 1000-01-01 MySQL limits
		startDate := time.Now().AddDate(0, -1*(i-1)%10000, 0).Format(SqlDatetimeFormat)
		endDate := time.Now().AddDate(0, -1*i%10000, 0).Format(SqlDatetimeFormat)
		_, e := fmt.Fprintf(writer, RowFormat, i, startDate, endDate, i*10)
		if e != nil {
			log.Fatal(e)
		}
	}
	e = writer.Flush()
	if e != nil {
		log.Fatal(e)
	}
	log.Printf("Duration: %v\n", time.Now().Sub(startTime))
}
