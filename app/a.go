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

func GetUserOutputStringBox(usr User) string {
	temp0 := fmt.Sprintf("╔" + "══════════" + "══════════" + "══════════" + "══════════" + "══════════" + "══════" + "╗\n")
	temp0 += fmt.Sprintf("║%15s %5s%-15s %20s\n", " ", " ", "-USER DATA-", "║")
	temp0 += fmt.Sprintf("║%15s %5s%-15s %20s\n", " ", " ", " ", "║")
	temp0 += fmt.Sprintf("║%15s%10s, %-20s%10s\n", "Name:", usr.first_name, usr.last_name, "║")
	temp0 += fmt.Sprintf("║%15s %10s %30s\n", "Gender:", usr.gender, "║")
	temp0 += fmt.Sprintf("║%15s %10d %30s\n", "ID NUMBER:", usr.id_Number, "║")
	temp0 += fmt.Sprintf("║%15s %10.2f %30s\n", "Score:", usr.user_score, "║")
	temp0 += fmt.Sprintf("║%15s %5s%-15s %20s\n", " ", " ", " ", "║")
	temp0 += fmt.Sprintf("╚" + "══════════" + "══════════" + "══════════" + "══════════" + "══════════" + "══════" + "╝")
	return temp0
}
func GetUserOutputStringBox2(usr User) string {
	temp0 := fmt.Sprintf("╔" + "══════════" + "══════════" + "══════════" + "══════════" + "═════" + "╗\n")
	temp0 += fmt.Sprintf("║%15s %-15s%15s\n", " ", "-USER DATA-", "║")
	temp0 += fmt.Sprintf("║%10s %-15s%20s\n", " ", " ", "║")
	temp0 += fmt.Sprintf("║%10s:%10s, %-20s%3s\n", "Name", usr.first_name, usr.last_name, "║")
	temp0 += fmt.Sprintf("║%10s:%10s%25s\n", "Gender", usr.gender, "║")
	temp0 += fmt.Sprintf("║%10s:%10d%25s\n", "ID NUMBER", usr.id_Number, "║")
	temp0 += fmt.Sprintf("║%10s:%10.2f%25s\n", "Score", usr.user_score, "║")
	temp0 += fmt.Sprintf("║%10s %-15s%20s\n", " ", " ", "║")
	temp0 += fmt.Sprintf("╚" + "══════════" + "══════════" + "══════════" + "══════════" + "═════" + "╝")
	return temp0
}

/*This makes a text box; it's the long version of doing it however; very inefficient
 */
func MakeATextBoxLong(strng []string) {
	var longest int = 0
	for i := 0; i < len(strng); i++ {
		if len(strng[i]) > longest {
			longest = len(strng[i])
		}
	}
	temps := "╔"
	for j := 0; j < longest; j++ {
		temps += "═"
	}
	temps += "╗\n"
	for i := 0; i < len(strng); i++ {

	}
}

// type UserData struct {
// 	userProf User
// }

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

/*Returns a true or false statement; taking into account the
 */
func GetYN(readr *bufio.Reader, writr *bufio.Writer) (bool, error) {

	var output = false
	strng, err := MyBufIOStringRead(readr)
	for {
		if strng == "y" {
			output = true
			break
		} else if strng == "n" {
			output = false
			break
		} else {
			WriterHelperNL(writr, "PLEASE HAVE CORRECT OUTPUT, Y/N?")
			strng, err = MyBufIOStringRead(readr)
		}

	}
	return output, err
}

func ShellStuff(readr *bufio.Reader, writr *bufio.Writer) int {
	strng, err := MyBufIOStringReadArray(readr)
	if err != nil {
		panic(err)
	}
	var ret int = 0
	if len(strng) > 0 {
		if len(strng) < 2 {
			strng = append(strng, "null")
		}
		switch {
		case strng[0] == "list":
			WriterHelperNL(writr, "LIST:"+" "+strng[1]+" "+"Y/N?")

			ret = 6
			break
		case strng[0] == "new":
			WriterHelperNL(writr, "NEW:"+" "+strng[1]+" "+"Y/N?")

			ret = 5
			break
		case strng[0] == "load":
			WriterHelperNL(writr, "LOAD:"+" "+strng[1]+" "+"Y/N?")

			ret = 4
			break
		case strng[0] == "save":
			WriterHelperNL(writr, "SAVE:"+" "+strng[1]+" "+"Y/N?")

			ret = 3
			break
		case strng[0] == "info":
			WriterHelperNL(writr, "USER INFO:"+" "+strng[1]+" "+"Y/N?")

			ret = 2
			break
		case strng[0] == "delete":
			WriterHelperNL(writr, "DELETE:"+" "+strng[1]+" "+"Y/N?")

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
	fmt.Printf("\n---o---\n")
	return ret
}
