package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	//// calling functions
	// total_matches(matches)
	// matches_per_year(matches)
	// Extra_runs_2016(matches, deliveries)
	economical_bowlers_2015(matches, deliveries)

}

func total_matches(matches []Matches) {

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

func matches_per_year(matches []Matches) {
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

func Extra_runs_2016(matches []Matches, deliveries []Deliveries) {
	var match2016 []int
	for _, each := range matches {
		if each.Season == 2016 {
			match2016 = append(match2016, each.Id)
		}
	}

	m := make(map[string]int)
	for _, each := range deliveries {
		for _, val := range match2016 {
			if val == each.Match_id {
				if m[each.Bowling_team] != 0 {
					m[each.Bowling_team] += each.Extra_runs
				} else {
					m[each.Bowling_team] = each.Extra_runs
				}
			}
		}
	}
	fmt.Println(m)
	jsondata, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create("../data/extra_runs_2016.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}

func economical_bowlers_2015(matches []Matches, deliveries []Deliveries) {
	var match2015 []int
	for _, each := range matches {
		if each.Season == 2015 {
			match2015 = append(match2015, each.Id)
		}
	}
	m := make(map[string]int)
	for _, each := range deliveries {
		for _, val := range match2015 {
			if val == each.Match_id {
				if m[each.Bowler] != 0 {
					m[each.Bowler] += each.Total_runs
				} else {
					m[each.Bowler] = each.Total_runs
				}
			}
		}
	}
	var overs = func(bowler string) int {
		var balls int
		for _, each := range deliveries {
			for _, val := range match2015 {
				if val == each.Match_id {
					if each.Bowler == bowler {
						balls += 1
					}
				}
			}
		}
		return balls / 6
	}

	var average = func(m map[string]int) map[string]int {
		for key, each := range m {
			m[key] = (each / overs(key))
		}
		return m
	}
	average(m)
	for key, value := range m {
		if value > 7 {
			delete(m, key)
		}
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] < m[keys[j]] })

	mp := map[string]int{}
	for i := 0; i < 10; i++ {
		mp[keys[i]] = m[keys[i]]

	}
	jsondata, err := json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create("../data/economic_bowlers.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}
