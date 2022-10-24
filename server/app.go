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

type Deliveries struct {
	Match_id         int
	Inning           int
	Batting_team     string
	Bowling_team     string
	Over             int
	Ball             int
	Batsman          string
	Non_striker      string
	Bowler           string
	Is_super_over    int
	Wide_runs        int
	Bye_runs         int
	Legbye_runs      int
	Noball_runs      int
	Penalty_runs     int
	Batsman_runs     int
	Extra_runs       int
	Total_runs       int
	Player_dismissed string
	Dismissal_kind   string
	Fielder          string
}

func main() {
	jsonFile1, err := os.Open("../data/deliveries.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue1, _ := ioutil.ReadAll(jsonFile1)
	var deliveries []Deliveries
	json.Unmarshal(byteValue1, &deliveries)
	jsonFile, err := os.Open("../data/matches.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var matches []Matches
	json.Unmarshal(byteValue, &matches)

	// calling functions
	bar(matches)
	foo(matches)

}

func foo(matches []Matches) {

	m := make(map[int]int)
	for _, each := range matches {
		if m[each.Season] != 0 {
			m[each.Season] += 1
		} else {
			m[each.Season] = 1
		}
		delete(m, 0)
	}

	fmt.Println(m)
	jsondata, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create("../data/matches_per_year.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()

}

func bar(matches []Matches) {
	m := make(map[int]map[string]int)
	for _, each := range matches {
		if m[each.Season] != nil {
			continue
		} else {
			m[each.Season] = match(each.Season)
		}
	}
	delete(m, 0)
	fmt.Println(m)
	jsondata, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create("../data/matches_won_per_year.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()

}

func match(years int) map[string]int {
	jsonFile, err := os.Open("../data/matches.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var matches []Matches
	json.Unmarshal(byteValue, &matches)

	mp := make(map[string]int)
	for _, each := range matches {
		if each.Season == years {
			if mp[each.Winner] != 0 {
				mp[each.Winner] += 1
			} else {
				mp[each.Winner] = 1
			}
		}
	}
	return mp

}
