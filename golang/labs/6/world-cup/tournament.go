package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

const N int = 1000000

func main() {
	var filename string

	fmt.Printf("Enter file path: ")
	fmt.Scan(&filename)

	start := time.Now()

	fileDescriptor, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't open file!", err)
		return
	}
	defer fileDescriptor.Close()

	fileReader := csv.NewReader(fileDescriptor)

	teams, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println("something went wrong when reading out teams!", err)
		return
	}

	// Print each team's chances of winning, according to simulation
	counts := make(map[string]int)
	for i := 0; i < N; i++ {
		counts[simulateTournament(teams)[0]] += 1
	}

	printResult(counts)

	fmt.Println("Time Executed:", time.Since(start).Seconds())
	// fmt.Printf("Time Executed: %dm%.3fs\n", elapse/60,)
}

// simulateGame return true if team1 wins and vice versal
func simulateGame(team1, team2 []string) bool {

	rating1, _ := strconv.Atoi(team1[1])
	rating2, _ := strconv.Atoi(team2[1])

	probability := 1 / (1 + math.Pow(10, float64(((rating2-rating1)/600))))
	return rand.Float64() < float64(probability)
}

// simulateRound
func simulateRound(teams [][]string) [][]string {
	winners := [][]string{
		{"team", "rating"},
	}

	//simulate games for all pairs of teams
	for i := 1; i < len(teams); i += 2 {
		if simulateGame(teams[i], teams[i+1]) {
			winners = append(winners, teams[i])
		} else {
			winners = append(winners, teams[i+1])
		}
	}

	return winners
}

// simulateTournament return the winning team
func simulateTournament(teams [][]string) []string {

	for len(teams) > 2 {
		teams = simulateRound(teams)
	}

	return teams[1] //
}

// printResult sorts and print out the percentage chance of winning of teams
func printResult(counts map[string]int) {
	keys := make([]string, 0, len(counts))
	for team := range counts {
		keys = append(keys, team)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return counts[keys[i]] > counts[keys[j]]
	})

	for _, team := range keys {
		fmt.Printf(" %d   ", counts[team])
	}
	fmt.Println()
	for _, team := range keys {
		fmt.Printf("%s: %.1f%% chance of winning\n", team, float64(counts[team])*100.0/float64(N))
	}
}
