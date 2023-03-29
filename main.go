package main

import (
	"fmt"
	"os"

	"github.com/vincepr/go_csv/csv"
	"github.com/vincepr/go_csv/terminal_size"
)

/*
*   Terminal cli to quickly print out data from a csv slightly formated +----+-----+
*   - prints exactly one terminal worth of data
*   - --line 23 to print rows from 23 onwards
*   - --all to print all rows
*   - should not load full csv in  memory but only useful part
 */


func main() {

    // load args
    path := loadArgs()

    // for now we just display the terminal size (displaying only part of the table was on the table)
    width,height,err := term.Size()
    if err != nil{
        fmt.Printf("%s\n",err)
        return
    }
    println("running on width:",width,"height:",height, "filepath:")

    // read the file
    table, err := csv.Read(path)
    if err != nil{
        fmt.Printf("%s\n",err)
        return
    }



    // printOut the table to the terminal
    csv.Print(table)
}


func loadArgs() string{
    args := os.Args
    if len(args) !=2 {
        panic("pass in filepath ex: ./go_csv ./files/addr.csv")
    }
    return args[1]
}