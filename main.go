package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	term "github.com/vincepr/go_csv/terminal_size"
)

/*
*   Terminal cli to quickly print out data from a csv slightly formated +----+-----+
*   - prints exactly one terminal worth of data
*   - --from 23 to print rows from 23 onwards
*   - should not load full csv in  memory but only useful part
 */

 
 func loadArgs() (string, int){
	args := os.Args
	if len(args) <2 {
		log.Fatal("wrong arguments, try passing in the filepath")
	} 
	if len(args) ==3{
		fromNr, err := strconv.ParseUint(args[2], 10, 32)
		if err!= nil{
			log.Fatalln("2nd argument MUST be positiveINT")
		}
		return args[1], int(fromNr)


	} else if len(args) !=2{
		log.Fatalln("only allowed arguments are: pathname firstRow?")
	}

    return args[1], 0
}

func main() {
	// load Arguments, :todo --flags
	path, startRow := loadArgs()

	// read the file
	str, err := readCsvFile(path)
	if err != nil{
		panic(err)
	}

	// get our terminal width and height to decide how many rows we need
	_, height, err := term.SizeXY()
	if err!= nil{
		panic(err)
	}
	endRow 		:= startRow + height - 3

	// parse the csv row by row
	r := csv.NewReader(strings.NewReader(str))
	targetRows := make([][]string, 0)
	nthRow := 0
	for {
		nthRow ++
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// only save target rows (for now):
		
		if nthRow > endRow{
			break
		} else if nthRow > startRow {
			targetRows = append(targetRows, row)
		}
	}

	// check if no output
	if len(targetRows) <1{
		log.Fatalln("--------------no rows found--------------")
	}
	printRowsFancy(targetRows, startRow)
}

func readCsvFile(path string) (string, error){
	buf, err := os.ReadFile(path)
	if err !=nil{
		return "", err
	}
	return string(buf), nil
}

func printRowsFancy(table [][]string, offset int){
    // calculate the max length of symbols for each column
    maxLen := make([]int, len(table[0]))
    for _,row := range table{
        for i, str := range row{
            len := len(str)
            if len > maxLen[i]{
                maxLen[i]=len
            }
        }
    }

    // print the table
    for _,len := range maxLen{
        fmt.Printf("+%s", strings.Repeat("-", len))
    }
    fmt.Printf("+")
    for idx,row := range table{
        fmt.Printf("\n|")
        for i ,str := range row{
            fmt.Printf("%s%s|",str, strings.Repeat(" ", maxLen[i]-len(str) ))
        }
		fmt.Printf("  %v",idx+offset)
    }
    fmt.Printf("\n")
    for _,len := range maxLen{
        fmt.Printf("+%s", strings.Repeat("-", len))
    }
    fmt.Printf("+\n")

}