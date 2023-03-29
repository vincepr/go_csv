package term

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

)
/*
*	linux only, returs width, height, and error of current terminal size. Uses Letters not pixels.
*/
func Size() (int, int, error){
	str, err :=getSize()
	if err != nil{
		return 0,0,err
	}
	w,h,err := parse(str)
	return w,h,err
}

func getSize() (string, error){
	cmd := exec.Command("stty", "size")
    cmd.Stdin = os.Stdin
    out, err := cmd.Output()
	return string(out), err
}

func parse(str string) (int, int, error) {
	arr := strings.Split(str, " ") 
	width, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, err
	}
	cleanStr := strings.TrimSpace(arr[1])	//cmd.Output() adds a newline we have to remove!
	height, err := strconv.Atoi(cleanStr)
	return width, height, err
}