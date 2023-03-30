# Minimal .csv preview in Terminal.

## What it does
- small footprint, and quick, you can preview into even big .csv files in the terminal. (only linux support atm)
- pass in filepath of csv file and optional index of FirstRow to display from
- uses full terminal width and height to display dataset. Colums above Terminal-width get cutt off.

![example of use](./files/example.gif)

## build the binary (golang required, or golang build dockerfile)
- `go build -o ./bin/csv`

## how to run
- `Binary FilepathToCsv OptionalStartRow`
- `./bin/csv ./files/grades.csv`
- `./bin/csv ./files/orga.csv 456`