package csv_custom

import (
	"fmt"
	"os"
	"strings"
)

/*
*   .csv files as defined in RFC-4180      https://www.rfc-editor.org/rfc/rfc4180
- each record is located on a separate line, delimted by a line break
- the last record in the file may or may not have an ending line break
(- there may be an optional header)
- within the header and each record there may be one or more fields separated by commas. each line should contain the same number of fields.
- spaces are part of a field and should not be ignored.
- the last record must not be followed by a comma. ex: aa,bb,ccc
- Each field my or may not be enclosed in double-quotes. If the field is not enclosed in double quotes, then there may not appear quotes inside the field.
- fields containing line breaks (CRLF), double quotes and commas should be enclosed in double-quotes:
ex:
    aaa, "b CRLF
    bb", "ccc" CRLF
    zzz,yyy,xxx
- if double-quotes are used to enclose fields, then a double-quote appearing inside a field must be escaped by preceding it with another double quote.
ex:
    "aaa", "b "bb"", "ccc"
*/


type Table [][]string

/*
*   Read in the csv and squeeze it into array
*/
func Read(path string) (Table, error) {
    dat, err := os.ReadFile(path)
	if err != nil{
		return nil, err
	}


    // split into rows on both "\n" and "\r" (windows)
    raw_rows := strings.FieldsFunc(string(dat), func(c rune) bool {
        return c == '\n' || c== '\r'
    })
    
    // remove repalce inner "" with '
    for i,row := range raw_rows{
        raw_rows[i] = strings.ReplaceAll(row,"\"\"","'")
    }

    // split into [rows][columns] , 
    rows := make([][]string,len(raw_rows))
    for i, row := range raw_rows{
        toManyCols:= strings.Split(row,",")
        var cols []string

        ubertrag := ""
        for _, part := range toManyCols{
            // if uneven count of " then we split on a internal , so we join them again till even count
            if strings.Count((ubertrag+part), "\"") %2 !=0{
                if ubertrag == ""{
                    ubertrag = part+","
                } else{
                    ubertrag = ubertrag+part+","
                }
            } else {
                cols = append(cols, (ubertrag+part))
                ubertrag=""
            }
        }
        rows[i] = cols
    }



	// making sure were not empty
	if len(rows)<1{
		return nil, fmt.Errorf("error, csv.Read(): rows <1")
	}
	// making sure every row is of same length
	length := -1
	for i,row := range rows{
		if length ==-1{
			length = len(row)
		} else if length != len(row){
			return nil, fmt.Errorf("error csv.Read(): NOT every row is of SAME LENGTH! have %v want %v on index: %v | value: %v", len(row), length, i, row)
		}
	}

	return rows, nil
}


// print out the csv to terminal
func Print(csv Table){
    // calculate the max length of symbols for each column
    maxLen := make([]int, len(csv[0]))
    for _,row := range csv{
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
    for _,row := range csv{
        fmt.Printf("\n|")
        for i ,str := range row{
            fmt.Printf("%s%s|",str, strings.Repeat(" ", maxLen[i]-len(str) ))
        }
    }
    fmt.Printf("\n")
    for _,len := range maxLen{
        fmt.Printf("+%s", strings.Repeat("-", len))
    }
    fmt.Printf("+\n")
}