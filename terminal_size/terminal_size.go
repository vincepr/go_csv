package term

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

)
/*
*	width, height, err of Terminal Size in Chars not pixels.
*/
func SizeXY() (int, int, error){
	str, err :=getSize()
	if err != nil{
		return 0,0,err
	}
	width , height, err := parse(str)
	return width, height, err
}

func getSize() (string, error){
	cmd := exec.Command("stty", "size")
    cmd.Stdin = os.Stdin
    out, err := cmd.Output()
	return string(out), err
}

func parse(str string) (int, int, error) {
	arr := strings.Split(str, " ") 
	height, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, err
	}
	cleanStr := strings.TrimSpace(arr[1])	//cmd.Output() adds a newline we have to remove!
	width, err := strconv.Atoi(cleanStr)
	return width, height, err
}