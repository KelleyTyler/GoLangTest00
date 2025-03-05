package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func A() {
	fmt.Printf("\n HELLO THIS IS A \n")
}

type User struct {
	first_name string
	last_name  string
	//date of birth?
	id_Number  int //lol
	gender     string
	user_score float64
}

func CreateUser(fname string, lname string, idNum int, gen string) User {
	temp := User{first_name: fname, last_name: lname, id_Number: idNum, gender: gen, user_score: 0.0}
	//seekvalidation;
	return temp
}
func GetUserOutputString(usr User) string {
	return fmt.Sprintf("Name: %-10s %-20s\nGender: %s\nID NUMBER: %4d\nScore:%3.2f", usr.first_name, usr.last_name, usr.gender, usr.id_Number, usr.user_score)
}

func WriterHelper(writr *bufio.Writer, str string) {
	writr.WriteString(str)
	writr.Flush()
}
func WriterHelperNL(writr *bufio.Writer, str string) {
	writr.WriteString(str + "\n")
	writr.Flush()
}
func MyBufIOStringRead(readr *bufio.Reader) (string, error) {
	rawString, err := readr.ReadString('\n')
	rawString = strings.ToLower(strings.Trim(rawString, "\r\n"))
	return rawString, err
}

func MyBufIOStringReadArray(readr *bufio.Reader) ([]string, error) {
	rawString, err := MyBufIOStringRead(readr)
	if err != nil {
		return nil, err
	}
	strngr := strings.Fields(rawString)
	return strngr, err
}
func GetANumber(readr *bufio.Reader, eRune rune) (int, error) {
	rawString, err := MyBufIOStringRead(readr)

	fmt.Printf("\n\nYOU ENTERED: %s\n\n", rawString)
	if err != nil {
		return 0, err
	}
	temp, err := strconv.Atoi(rawString)
	fmt.Printf("\n TEMP IS %d\n\n", temp)
	return temp, err
}
func GetArrayOfNumbersULim(readr *bufio.Reader) ([]int, error) {
	tempStringArr, err := MyBufIOStringReadArray(readr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("tempStringArr is %d spaces long \n", len(tempStringArr))
	var tempAr []int
	for _, s := range tempStringArr {
		tempI, err := strconv.Atoi(s)
		if err != nil {
			return tempAr, err
		}
		tempAr = append(tempAr, tempI)
	}
	return tempAr, err
}
func GetArrayOfNumbersLim(readr *bufio.Reader, lim int) ([]int, error) {
	tempStringArr, err := MyBufIOStringReadArray(readr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("tempStringArr is %d spaces long \n", len(tempStringArr))
	var tempAr []int
	for i, s := range tempStringArr {
		tempI, err := strconv.Atoi(s)
		if err != nil {
			return tempAr, err
		}
		if i < lim {
			tempAr = append(tempAr, tempI)
		}
	}
	return tempAr, err
}

func ShellStuff(readr *bufio.Reader, writr *bufio.Writer) int {
	strng, err := MyBufIOStringReadArray(readr)
	if err != nil {
		panic(err)
	}
	var ret int = 0
	if len(strng) > 0 {
		switch {
		case strng[0] == "new":
			WriterHelperNL(writr, "NEW:"+" "+strng[1])
			ret = 5
			break
		case strng[0] == "info":
			WriterHelperNL(writr, "USER INFO:"+" "+strng[1])
			ret = 4
			break
		case strng[0] == "load":
			WriterHelperNL(writr, "LOAD:"+" "+strng[1])
			ret = 3
			break
		case strng[0] == "save":
			WriterHelperNL(writr, "SAVE:"+" "+strng[1])
			ret = 2
			break
		case strng[0] == "delete":
			WriterHelperNL(writr, "DELETE:"+" "+strng[1])
			ret = 1
			break
		case strng[0] == "quit":
			WriterHelperNL(writr, "QUITTING")
			ret = -1
			break
		default:
			WriterHelperNL(writr, "DEFAULT STUFF")
			ret = 0
		}
	}
	fmt.Printf("----------\n")
	return ret
}
