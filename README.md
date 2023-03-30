# Minimal .csv preview in Terminal.

## What it does
- small footprint, and quick, you can preview into even big .csv files in the terminal. 
- only linux/unix support for now
- pass in filepath of csv file and optional index of FirstRow to display from
- uses full terminal width and height to display dataset. Colums above Terminal-width get cutt off.

![example of use](./files/example.gif)

## build the binary (golang required, or golang build dockerfile)
- clone the repository cd into the folder
- then build with golang for your target system `go build -o ./bin/csv`

## how to run
- `Binary FilepathToCsv OptionalStartRow`
- `./bin/csv ./files/grades.csv`
- `./bin/csv ./files/orga.csv 456`

## known limitations
since it uses the std-lib csv file reader trailing spaces after doublequte enclosed fields are not supported. ex: ` "asdf" ,"bbb"` will fail while `asdf , bbb` or ` "asdf",  "bbb"` are ok. (https://github.com/golang/go/issues/25131 if every in need to adjust the FileReader implementation)