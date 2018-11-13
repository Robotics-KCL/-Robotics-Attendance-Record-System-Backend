package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// for record every entry in the database
type Member struct {
	Name             string
	Gender           string
	Course           string
	Faculty          string
	AttendedWorkshop int16
	Date             int32
}

func main() {
	fmt.Println("Starting working")
	Writein(Member{Name: "Alex Bonin", Gender: "Male", Course: "Electonics", Faculty: "NMS", AttendedWorkshop: 10, Date: 13112018})

}

func Writein(newMember Member) {

	/*
	 * code for opening a json file from storage
	 */
	jsonFile, err := os.Open("members.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	//read all data in the json file through ioutil
	existingData, _ := ioutil.ReadAll(jsonFile)

	var mem []Member

	if err := json.Unmarshal([]byte(existingData), &mem); err != nil {
		log.Println(err)
	}

	mem = append(mem, newMember)

	newJson, err := json.Marshal(mem)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(newJson))

	ioutil.WriteFile("members.json", newJson, 0644)

}
