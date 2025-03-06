package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello World\n")
	A()
	//fileTest()
	uIReadr := bufio.NewReader(os.Stdin)
	uIWriter := bufio.NewWriter(os.Stdout)
	//var uInStr string = ""
	//---------------------

	var mySettings Settings
	mySettings = mySettings.ReadFromFile("init.json")
	mySettings.PrintDisplayStrings(uIWriter)
	usr0 := CreateUser("John", "Doe", 12424, "M")
	//fmt.Print(GetUserOutputStringBox2(usr0))
	usr0.PrintOutPutBox(uIWriter)
	fmt.Printf("\n---WELCOME!---\n--BASIC SHELL--\n-----------------\n")
	for {
		//fmt.Printf(">")
		// uInStr, _ = MyBufIOStringRead(uIReadr)
		// uInStr = strings.ToLower(strings.Trim(uInStr, "\r\n"))
		//uInStr = strings.Replace(uInStr, "\n", "", -1)
		//fmt.Printf()
		// temp := fmt.Sprintf("\n\nSTRING--->%s\n", uInStr)
		// uIWriter.WriteString(temp)
		// uIWriter.Flush()
		ret := ShellStuff(uIReadr, uIWriter)
		// if strings.Contains(uInStr, "quit") {
		// 	break
		// } else {
		// 	// ShellStuff(uInStr)
		// }
		if ret == -1 {
			fmt.Printf("EXITING:\n")
			break
		}
	}
	fmt.Printf("ENTER SOME NUMBERS:")
	mee, _ := GetArrayOfNumbersULim(uIReadr)
	fmt.Print(mee)
	//fmt.Printf("\n\nGET A NUMBER: %d\n", mee)
	uIWriter.Flush()
	fmt.Scanf("\n")
}
