package csv_own

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
	"unicode/utf8"
)

// File Stream Reader to Read csv files row by row
type Reader struct{
	reader *bufio.Reader
	columnLength int
	

	// Our Character that seperates rows, usually comma: ',' for comma-seperated-values:
	Delimiter rune	

	// set these to ignore whitespace at the begin or end of fields
	TrimLeadingSpace 			bool
	
	// ignore space after a quotation enclosed field: ex: "a" ,"b" , c -> [a,b,c] instead of [a ,b ,c]
	TrimTrailingSpaceQuotes 	bool		
	
	// ignore space after not quotation enclosed fields aswell. ex: 12  ,1432 , 3 -> [12,1432, 3]
	TrimTrailingSpaceDefault 	bool
	
	// in a default .csv all rows must have the same ammount of columns
	AllowDifferentColumnLength	bool
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
				//fmt.Printf("ReadAll() EOF reached! \n")		// TODO: remove when done
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

	delimiterLen 	:= utf8.RuneLen(r.Delimiter)
	quoteLen 		:= utf8.RuneLen('"')
	quoteByte 		:= []byte{'"'}
	row 			:= []string{}
	OuterLoop:
	for {
		if true{
			line = bytes.TrimLeftFunc(line, unicode.IsSpace)
		}

		if line[0] != '"'{
			// easy-case not "" enclosed field:
			field 	:= line
			idx 	:=  bytes.IndexRune(line, r.Delimiter)
			if idx > -1 {
				field 	= field[:idx]
				line 	= line[(idx+delimiterLen) :]
			}
			if r.TrimTrailingSpaceQuotes {
				field = bytes.TrimRightFunc(field, unicode.IsSpace)
			}
			if bytes.ContainsRune(field, '"'){

				panic(`Parse Error: field contains " in the middle`)
			}

			row = append(row, string(field))
			if idx == -1{
				break	// reached line-end
			}
		} else{
			buf := []byte{}
			carryOverToNextLine := false
			// quotes-case, field starts with a " and thus is enclosed:
			line 	= line[quoteLen:]	//remove first '"'

			for{
				idx 	:= bytes.IndexRune(line, r.Delimiter)
				field	:= line
				if idx > -1 {
					// found next delim
					field 	= field[:idx+delimiterLen]
					line 	= line[(idx+delimiterLen) :]
				
				} else {
					// field goes to newline
					carryOverToNextLine = true
				}
				
				count := bytes.Count(field, quoteByte)
				if count %2 ==1{
					// "" count adds up -> finished the field:
					// TODO: possbile bug? can this loop run forever if wront format?
					if field[len(field)-delimiterLen] == byte('"') {
						buf = append(buf, field...)
					} else{
						buf = append(buf, field[:len(field)-delimiterLen]...)
					}
					
					if r.TrimTrailingSpaceQuotes {
						buf = bytes.TrimRightFunc(buf, unicode.IsSpace)
					}
					if buf[len(buf)-1] == byte('"') {
						buf = buf[:len(buf)-1]
					} else {
						panic("ERROR parsing csv in readRow() : fields that begin with a quote must end with one")
					}

					// found end
					row = append(row, string(buf))
					if idx == -1{
						break OuterLoop
					}
					break
				} else{
					// carry over to next field without newline
					

					if carryOverToNextLine{
						field := append([]byte("\n"), field...)
						buf = append(buf, field...)
						// carry over to next field with newline
						for err == nil{
							line, err = r.readLine()
							
							// skip empty lines:
							if len(line)== 0 {
								continue 					
							}
							break
						}
						if err != nil {
							return nil, err		// if EOF reached no need to parse anymore
						}
					} else{
						buf = append(buf, field...)
					}
				}
			}
			//bytes.ReplaceAll("" with ")
			//buf = append(buf, []byte("121323")... )
		}
	}

	//fmt.Printf("END: row: '%s'\n", row)
	// check if same column size across file:
	if !r.AllowDifferentColumnLength && r.columnLength != len(row){
		if r.columnLength == 0{
			r.columnLength = len(row)
		} else {
			panic("ERROR: Cant Parse. All rows MUST have same column count")
		}
	}

	return row, nil
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


