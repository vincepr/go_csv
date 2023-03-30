package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	util "github.com/vincepr/go_csv/csv_util"
	term "github.com/vincepr/go_csv/terminal_size"
)

/*
*   Terminal cli to quickly print out data from a csv slightly formated +----+-----+
*	- quick (enough for me), even for files of a few 100Mb.
*   - prints exactly one terminal-heigh worth of data
*   - example for first few rows: ./gocsv ./files/travel.csv		
*   - example for first few rows: ./gocsv ./files/travel.csv 5		for lines 5 and onwards
 */

 
func main() {
	// load Arguments
	path, startRow := loadArgs()

	// read the file
	str, err := util.ReadCsvFile(path)
	if err != nil{
		panic(err)
	}

	// get our terminal width and height to decide how many rows we need
	width, height, err := term.SizeXY()
	if err!= nil{
		panic(err)
	}
	endRow 		:= startRow + height - 3

	// parse the csv row by row
	csvReader := csv.NewReader(strings.NewReader(str))
	csvReader.TrimLeadingSpace = true
	targetRows := make([][]string, 0)
	nthRow := 0
	for {
		nthRow ++
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		
		// only save targeted rows:
		if nthRow > endRow{
			break
		} else if nthRow > startRow {
			targetRows = append(targetRows, row)
		}
	}

	// error message if empty
	if len(targetRows) <1{
		log.Fatalln("--------------no rows found--------------")
	}
	util.PrintRowsFancy(targetRows, startRow, width)
}

func ReadCsvFile(path string) {
	panic("unimplemented")
}

func readCsvFile(path string) {
	panic("unimplemented")
}

// loads terminal arguments ( PATHNAME FIRSTROW ) and error checks them
func loadArgs() (string, int){
	args := os.Args
	if len(args) <2 {
		log.Fatal("MUST pass in the filepath")
	} 

	if len(args) ==3{
		fromNr, err := strconv.ParseUint(args[2], 10, 32)
		if err!= nil{
			log.Fatalln("2nd argument MUST be positiveINT")
		}
		return args[1], int(fromNr)

	} else if len(args) !=2{
		log.Fatalln("only allowed arguments are: PATHNAME optionalSTARTROW")
	}
    return args[1], 0
}
