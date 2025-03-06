package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello World\n")
	A()
	uIReadr := bufio.NewReader(os.Stdin)
	uIWriter := bufio.NewWriter(os.Stdout)
	var mySettings Settings
	mySettings = mySettings.ReadFromFile("init.json")
	mySettings.PrintDisplayStrings(uIWriter)
	usr0 := CreateUser("John", "Doe", 12424, "M")
	usr0.PrintOutPutBox(uIWriter)
	fmt.Printf("\n---WELCOME!---\n--BASIC SHELL--\n-----------------\n")
	for {
		ret := ShellStuff(uIReadr, uIWriter)
		if ret == -1 {
			fmt.Printf("EXITING:\n")
			break
		}
	}
	fmt.Printf("ENTER SOME NUMBERS:")
	mee, _ := GetArrayOfNumbersULim(uIReadr)
	fmt.Print(mee)
	uIWriter.Flush()
	fmt.Scanf("\n")
}
