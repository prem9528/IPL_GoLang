package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)
type Matches struct{
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
	foo()
	bar()

}
func bar(){
	csvFile, err := os.Open("../data/matches.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Matches
	var allRecords []Matches

	for _, each := range csvData {

		oneRecord.Id, _ = strconv.Atoi(each[0])
		oneRecord.Season, _ = strconv.Atoi(each[1])
		oneRecord.City=each[2]
		oneRecord.Date=each[3]
		oneRecord.Team1=each[4]
		oneRecord.Team2=each[5]
		oneRecord.Toss_winner=each[6]
		oneRecord.Toss_decision=each[7]
		oneRecord.Result=each[8]
		oneRecord.Dl_applied, _ = strconv.Atoi(each[9])
		oneRecord.Winner=each[10]
		oneRecord.Win_by_runs, _ = strconv.Atoi(each[11])
		oneRecord.Win_by_wickets, _ = strconv.Atoi(each[12])
		oneRecord.Payer_of_match=each[13]
		oneRecord.Venue=each[14]
		oneRecord.Umpire1=each[15]
		oneRecord.Umpire2=each[16]
		oneRecord.Umpire3=each[17]
		allRecords = append(allRecords, oneRecord)

	}

	jsondata, err := json.Marshal(allRecords)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create("../data/matches.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}

func foo(){
	csvFile, err := os.Open("../data/deliveries.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Deliveries
	var allRecords []Deliveries

	for _, each := range csvData {

		oneRecord.Match_id, _ = strconv.Atoi(each[0])
		oneRecord.Inning, _ = strconv.Atoi(each[1])
		oneRecord.Batting_team = each[2]
		oneRecord.Bowling_team = each[3]
		oneRecord.Over, _ = strconv.Atoi(each[4])
		oneRecord.Ball, _ = strconv.Atoi(each[5])
		oneRecord.Batsman = each[6]
		oneRecord.Non_striker = each[7]
		oneRecord.Bowler = each[8]
		oneRecord.Is_super_over, _ = strconv.Atoi(each[9])
		oneRecord.Wide_runs, _ = strconv.Atoi(each[10])
		oneRecord.Bye_runs, _ = strconv.Atoi(each[11])
		oneRecord.Legbye_runs, _ = strconv.Atoi(each[12])
		oneRecord.Noball_runs, _ = strconv.Atoi(each[13])
		oneRecord.Penalty_runs, _ = strconv.Atoi(each[14])
		oneRecord.Batsman_runs, _ = strconv.Atoi(each[15])
		oneRecord.Extra_runs, _ = strconv.Atoi(each[16])
		oneRecord.Total_runs, _ = strconv.Atoi(each[17])
		oneRecord.Player_dismissed = each[18]
		oneRecord.Dismissal_kind = each[19]
		oneRecord.Fielder = each[20]
		allRecords = append(allRecords, oneRecord)

	}

	jsondata, err := json.Marshal(allRecords)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create("../data/deliveries.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsondata)
	jsonFile.Close()
}