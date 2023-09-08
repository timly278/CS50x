package main

import "fmt"

// Max number of candidates
const MAX int = 9

// preferences[i][j] is number of voters who prefer i over j
var preferences [MAX][MAX]int

// locked[i][j] means i is locked in over j
var locked [MAX][MAX]bool

// Each pair has a winner, loser
type pair struct {
	winner int
	loser  int
	diff   int
}

var candidates [MAX]string

var pairs [MAX * (MAX - 1) / 2]pair

var pair_count int
var candidate_count int

func main() {

	// Check for invalid usage, number of candidates
	// Populate array of candidates
	if !getCandidate() { // TODO: handle the double enters
		return
	}

	// Input voters
	voter_count := 0
	fmt.Printf("Number of voters: ")
	fmt.Scanf("%d", &voter_count) // TODO: loop for to get right number

	for i := 0; i < voter_count; i++ {
		ranks := make([]int, candidate_count)

		for j := 0; j < candidate_count; j++ {
			var name string
			fmt.Printf("Rank %d: ", j+1)
			fmt.Scanf("%s", &name)

			if !vote(j, name, ranks) {
				fmt.Println("Invalid vote.")
				return
			}

		}

		record_preferences(ranks)
		fmt.Printf("\n")
	}

	// Initilize locked
	for i := 0; i < candidate_count; i++ {
		for j := 0; j < candidate_count; j++ {
			locked[i][j] = false
		}
	}

	add_pairs()
	sort_pairs()
	lock_pairs()
	print_winner()

}

// getCandidate input the candidate from keyboard
func getCandidate() bool {
	var text [MAX]string
	retValue := true
	fmt.Println("Enter the candidates please:")
	for candidate_count = 0; candidate_count < MAX+1; candidate_count++ {
		n, err := fmt.Scanf("%s", &text[candidate_count])
		if n == 0 {
			if candidate_count == 0 || candidate_count == 1 {
				fmt.Println("Error Usage: tideman [candidate ...]")
				retValue = false
			}

			break
		}
		if n > 0 {
			if n > MAX {
				fmt.Printf("Maximum number of candidates is %d\n", MAX)
				retValue = false
			}
		}
		if err != nil {
			fmt.Println("error:", err)
			retValue = false
		}
	}
	if retValue {
		fmt.Println("Num of candidate:", candidate_count)
		for i := 0; i < candidate_count; i++ {
			candidates[i] = text[i]
		}
		fmt.Println("List of candidate:", candidates)
	}

	return retValue
}

// vote
func vote(rank int, name string, ranks []int) bool {
	ret := true
	i := 0
	for i = 0; i < candidate_count; i++ {
		if name == candidates[i] {
			ranks[rank] = i
			break
		}
	}
	if i == candidate_count {
		ret = false
	}

	return ret
}

// record_preferences
func record_preferences(ranks []int) {
	for i := range ranks {
		for j := i + 1; j < candidate_count; j++ {
			preferences[ranks[i]][ranks[j]]++
		}
	}
}

// add_pairs
func add_pairs() {
	pair_count = 0
	for i := 0; i < candidate_count; i++ {
		for j := 0; j < candidate_count; j++ {
			if j > i {
				if preferences[i][j] > preferences[j][i] {
					pairs[pair_count].loser = j
					pairs[pair_count].winner = i
				} else {
					pairs[pair_count].loser = i
					pairs[pair_count].winner = j
				}
				pairs[pair_count].diff = preferences[pairs[pair_count].winner][pairs[pair_count].loser] - preferences[pairs[pair_count].loser][pairs[pair_count].winner]
				pair_count++
			}
		}
	}
}

// sort_pairs
func sort_pairs() {

	for i := 0; i < pair_count; i++ {
		for j := i + 1; j < pair_count; j++ {
			if pairs[i].diff < pairs[j].diff {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}
}

/*

I figured out the order of the pairs's element affects the final outcome
for example:
	unsort: (2, 3) (2, 4) (3, 4) (1, 2) (1, 3) (1, 4) (0, 1) (0, 4) (2, 0) (3, 0)
	sorted: (2, 3) (2, 4) (3, 4) (1, 2) (1, 3) (1, 4) `(0, 1) (0, 4) (2, 0) (3, 0)`
	and
	unsort2: {0 1} {2 0} {3 0} {4 0} {1 2} {1 3} {1 4} {2 3} {2 4} {3 4}
	sorted2: {2 3} {2 4} {3 4} {1 2} {1 3} {1 4} `{3 0} {4 0} {0 1} {2 0}`
As you can see, the last four elements of two cases have different order.
and this will comeout the different results.

=> maybe the problem is at add_pairs() function.

*/

// lock_pairs
func lock_pairs() {
	// the first two pair element would be true all the time.
	locked[pairs[0].winner][pairs[0].loser] = true
	if pair_count > 1 {
		locked[pairs[1].winner][pairs[1].loser] = true
	}

	for i := 2; i < pair_count; i++ {
		// find to avoid cycle
		if !isCycle(pairs[i].winner, pairs[i].loser) {
			locked[pairs[i].winner][pairs[i].loser] = true
		}

	}
}

// print_winner
func print_winner() {
	isNotWinner := false
	var winner int
	for winner = 0; winner < candidate_count; winner++ {
		// to find which one is source - not a loser at any time
		for winIdx := 0; winIdx < candidate_count; winIdx++ {
			if winner != winIdx {
				for loseIdx := 0; loseIdx < candidate_count; loseIdx++ {
					if (loseIdx == winner) && locked[winIdx][loseIdx] {
						isNotWinner = true
					}
				}
			}
		}
		if isNotWinner {
			isNotWinner = false
		} else {
			fmt.Printf("\nwinner is candidates[%d]: %s\n", winner, candidates[winner])
			return
		}
	}
}

// isCycle scans the locked[i][j] to figure out
// if there would be any cycle with the parameter `loser`
func isCycle(winner, loser int) bool {
	ret := false

	for i := 0; i < candidate_count; i++ {
		if locked[loser][i] == true {
			// find i as a winner that makes locked[i][x] = true
			for j := 0; j < candidate_count; j++ {
				if locked[i][j] == true && j == winner {
					ret = true
				}
			}
		}
	}

	return ret
}
