package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct{
	Id int
	Season int
	City string
	Date string
	Team1 string
	Team2 string
	Toss_winner string
	Toss_decision string
	Result string
	Dl_applied int
	Winner string
	Win_by_runs int
	Win_by_wickets int
	Payer_of_match string
	Venue string
	Umpire1 string
	Umpire2 string
	Umpire3 string
}
func main() {
	jsonFile, err := os.Open("../data/matches.json")
	if err != nil {
		fmt.Println(err)
	}
	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// var users []Users
	
	// json.Unmarshal(byteValue, &users)

	// for _, each := range users{
	// 	fmt.Println(each.Season)
	// }

	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for{
		n, err := file.Read(buf)
		
	}
}
