package main

import (
	"encoding/json"
	"fmt"
	// "strconv"
	"io/ioutil"
	"os"
)

type Users struct {
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
	var users []Users
	json.Unmarshal(byteValue, &users)
	m := make(map[int]int)
	for _, each := range users {
		if m[each.Season] != 0{
			m[each.Season] += 1
		}else{
			m[each.Season] = 1
		}
		delete(m, 0)
	}
	
	fmt.Println(m)
	// jsondata, err := json.Marshal(m)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	
	// jsonFile2, err := os.Create("../data/matches_per_year.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer jsonFile2.Close()

	// jsonFile.Write(jsondata)
	// jsonFile.Close()

}
