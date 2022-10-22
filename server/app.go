package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Matches struct {
	Id             int
	Season         int
	City           string
	Date           string
	Team1          string
	Team2          string
	Toss_winner    string
	Toss_decision  string
	Result         string
	Dl_applied     int
	Winner         string
	Win_by_runs    int
	Win_by_wickets int
	Payer_of_match string
	Venue          string
	Umpire1        string
	Umpire2        string
	Umpire3        string
}

func main() {
	jsonFile, err := os.Open("../data/matches.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var matches []Matches
	json.Unmarshal(byteValue, &matches)
	foo(matches)
	
}

func foo(matches []Matches){

	m := make(map[int]int)
	for _, each := range matches {
		if m[each.Season] != 0{
			m[each.Season] += 1
		}else{
			m[each.Season] = 1
		}
		delete(m, 0)
	}
	
	fmt.Println(m)

}
