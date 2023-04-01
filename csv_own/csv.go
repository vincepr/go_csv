package csv_own

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

// File Stream Reader to Read csv files row by row
type Reader struct{
	reader *bufio.Reader
	
	// Our Character that seperates rows, usually comma: ',' for comma-seperated-values:
	Delimiter rune	
	

}

// constructor for Reader to read csv files row by row from r
func NewReader(r io.Reader) *Reader{
	return &Reader{
		Delimiter: ',',
		reader:	bufio.NewReader(r),
	}
}

// tries to read the whole csv file till EOF
// if it fails returns an error
func (r *Reader) ReadAll() (rows [][]string, err error){
	for {
		row, err := r.readRow()
		if err != nil{
			if err == io.EOF{
				fmt.Printf("ReadAll() EOF reached! \n")		// TODO: remove when done
				return rows, err
			}
			return nil, err
		}
		rows = append(rows, row)
	}
}

// Read one Row worth of csv table and returns it as []string.
// err if EOF. err if some parsing error
func (r *Reader) Read() (row []string, err error){
	row, err = r.readRow()
	return row, err
}

// main func that does most the parsing
// notice that multiple lines can make up a row if enclosed in " "
func (r *Reader) readRow() ([]string, error){
	var line []byte
	var err error

	/* read one Line */
	for err == nil{
		line, err = r.readLine()
		
		// TODO: add skiping leading whitespace
		// skip empty lines:
		if len(line)== 0 {
			continue 					
		}
		break
	}
	if err != nil {
		return nil, err		// if EOF reached no need to parse anymore
	}

	/* try to parse the line */
	fmt.Printf("full line: %s \n", line)
	delimiterLen 	:= utf8.RuneLen(r.Delimiter)
	quoteLen 		:= utf8.RuneLen('"')
	quoteByte 		:= []byte{'"'}
	for {
		if line[0] != '"'{
			// easy-case not "" enclosed field:
			field 	:= line
			idx 	:=  bytes.IndexRune(line, r.Delimiter)
			if bytes.ContainsRune(field, '"'){
				panic(`Parse Error: field contains " in the middle`)
			}
			if idx > -1 {
				field 	= field[:idx]
				line 	= line[(idx+delimiterLen) :]
			}
			fmt.Printf("field: '%s'\n", field)
			if idx == -1{
				// TODO: check if same length as previous-rows if exist already
				break	// reached line-end
			}
		} else{
			// quotes-case, field starts with a " and thus is enclosed:
			line 	= line[quoteLen:]	//remove first '"'
			fmt.Printf("line without first'': '%s'\n", line)
			// var fieldBuffer []byte
			idx 	:= bytes.IndexRune(line, r.Delimiter)
			count	:= bytes.Count(line, quoteByte)
			if count %2 ==1{	// make for loop of this
				fmt.Println("count:", count, "idx:", idx)
			}
			field	:=line
			if idx > -1 {
				field 	= field[:idx]
				line 	= line[(idx+delimiterLen) :]

			}

			
			
			//bytes.ReplaceAll("" with ")
			//buf = append(buf, []byte("121323")... )



		}
	}



	return nil, nil


}

// helper func to read till newline or EOF
func (r *Reader) readLine() ([]byte, error){
	line, _, err := r.reader.ReadLine()
	if err == bufio.ErrBufferFull {
		// TODO: check if this can ErrBufferFull and then handle it with looping and appending slices together.
		println("ERROR: in readLine() ErrBufferFull with bufio.reader.ReadLine()")
		panic(err)
	}
	return line, err
}


