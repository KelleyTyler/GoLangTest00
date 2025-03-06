package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	//"strconv"
	//"strings"
)

func fileTest() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	// open output file
	fo, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
}

/*
INTERFACE EXPLORATION:
*/

/*This Interface is meant to be used for structs;
 */
type MyStdData interface {
	GetDisplayStrings()
	PrintDisplayStrings()
	SaveToPlace()
	ReadFromPlace()
	DeleteFromPlace()
}

type SettingsJson struct {
	IsDev     bool   `json:"isDev,omitempty"`
	VersionID string `json:"versionID,omitempty"`
	DevName   string `json:"devName,omitempty"`
}

type Settings struct {
	isDev     bool   //`json:"isDev,omitempty"`
	versionID string //`json:"versionID,omitempty"`
	devName   string //`json:"devName,omitempty"`
}

func (stng Settings) GetDisplayStrings() [3]string {
	var temp [3]string
	temp[0] = fmt.Sprintf("is Dev: %t", stng.isDev)

	temp[1] = fmt.Sprintf("Version: %s", stng.versionID)
	temp[2] = fmt.Sprintf("Dev: %s", stng.devName)
	return temp
}

func (stng Settings) PrintDisplayStrings(writr *bufio.Writer) {
	mee := stng.GetDisplayStrings()
	for i, s := range mee {
		WriterHelperNL(writr, fmt.Sprintf("%d %s", i, s))
	}
}

func (stng *Settings) MarshalJSON() ([]byte, error) {
	return json.Marshal(SettingsJson{
		stng.isDev,
		stng.versionID,
		stng.devName,
	})
}
func (stng *Settings) UnmarshalJSON(b []byte) error {
	temp := &SettingsJson{}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	stng.devName = temp.DevName
	stng.isDev = temp.IsDev
	stng.versionID = temp.VersionID
	return nil
}

// func (stng Settings) MarshalText()

/*
this needs to read and parse an init file;
this file will be encoded in JSON format;
*/
func (stng Settings) ReadFromFile(fp string) Settings {
	var temp Settings = Settings{true, "444", "444"}
	jSonFile, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := jSonFile.Close(); err != nil {
			panic(err)
		}
	}()
	rdr := bufio.NewReader(jSonFile)
	//rdal, err2 := io.ReadAll(rdr)
	rdal, err2 := io.ReadAll(rdr)
	if err2 != nil {
		panic(err2)
	}
	// rTemp2, err2 := io.ReadAll(rdr)
	// if err2 != nil {
	// 	panic(err2)

	// }

	// for _, re := range rdal {
	// 	fmt.Printf("%c", rune(re))
	// }
	err3 := json.Unmarshal(rdal, &temp)
	if err3 != nil {
		panic(err3)
	}
	// fmt.Print("\n \n TEMP: ")
	// fmt.Print(temp)
	// fmt.Print("\n \n")
	return temp
}

/*

 */
// func InitReader() {

// }
