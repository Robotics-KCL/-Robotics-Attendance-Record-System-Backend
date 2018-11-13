package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// for record every entry in the database
type Member struct {
	Name             string
	Gender           string
	Course           string
	faculty          string
	AttendedWorkshop int16
	Data             int16
}

func main() {

}

func Writein(jsonString []byte) {

	var jsonText = []Member

	if err:= json.Unmarshal([]byte(jsonString),&jsonText); err !=nil{
		log.Println(err)
	}
}
