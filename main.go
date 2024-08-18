package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	ID    int
	Dices []int
	Score int
}

// Bug: due to the rule that dice 1 is given to the next player, Player #1 has an advantage to win because he has more dice than other players
func main() {
	totalPlayer := 3
	totalDices := 5

	players := createPlayers(totalPlayer, totalDices)
	finalResult := make([]Player, len(players))
	turn(&players, &finalResult)
}

func createPlayers(totalPlayer int, totalDice int) []Player {
	var players []Player
	for i := 1; i <= totalPlayer; i++ {
		players = append(players, Player{
			ID:    i,
			Dices: make([]int, totalDice),
			Score: 0,
		})
	}
	return players
}

func turn(players *[]Player, finalResult *[]Player) {
	turnNumber := 1
	for {
		if checkPlayerDiceLeft(*players) == 1 {
			fmt.Printf("===========================================\n")
			fmt.Printf("Final Result:\n")
			winner := Player{}
			for i := range *finalResult {
				fmt.Printf("\t Player #%d (%d) : %v \n", (*finalResult)[i].ID, (*finalResult)[i].Score, (*finalResult)[i].Dices)
				if (*finalResult)[i].Score > winner.Score {
					winner = (*finalResult)[i]
				} else if (*finalResult)[i].Score == winner.Score {
					if len(winner.Dices) < len((*finalResult)[i].Dices) {
						winner = (*finalResult)[i]
					}
				}
			}
			fmt.Printf("The winner is Player #%d with score %d\n", winner.ID, winner.Score)
			break
		}
		fmt.Printf("===========================================\n")
		fmt.Printf("Turn %d:\n", turnNumber)
		for i := range *players {
			if len((*players)[i].Dices) == 0 {
				continue
			}
			for j := 0; j < len((*players)[i].Dices); j++ {
				(*players)[i].Dices[j] = rollDice()
			}

			fmt.Printf("\t Player #%d (%d) : %v \n", (*players)[i].ID, (*players)[i].Score, (*players)[i].Dices)
		}
		evaluateDice(players, finalResult)

		turnNumber++
	}
}

func checkPlayerDiceLeft(players []Player) int {
	playerHaveDiceLeft := 0
	for i := range players {
		if len(players[i].Dices) > 0 {
			playerHaveDiceLeft++
		}
	}
	return playerHaveDiceLeft
}

func rollDice() int {
	return rand.Intn(6) + 1
}

func evaluateDice(players *[]Player, finalResult *[]Player) {
	fmt.Printf("Evaluate Result:\n")
	additionalDices := make(map[int][]int)
	for i := range *players {
		player := &(*players)[i]
		diceIdx := 0
		for diceIdx < len(player.Dices) {
			dice := player.Dices[diceIdx]
			if dice == 6 {
				player.Score++
				player.Dices = removeIndex(player.Dices, diceIdx)
				// Continue to evaluate the same index after removal
				continue
			} else if dice == 1 {
				if i+1 >= len(*players) {
					additionalDices[0] = append(additionalDices[0], 1)
				} else {
					additionalDices[i] = append(additionalDices[i], 1)
				}
				player.Dices = removeIndex(player.Dices, diceIdx)
				// Continue to evaluate the same index after removal
				continue
			}
			// Move to next index if no action was taken
			diceIdx++
		}
	}

	// Additional Dice 1
	for key, dices := range additionalDices {
		for i := range *players {
			player := &(*players)[i]
			if key == i {
				player.Dices = append(player.Dices, dices...)
			}
		}
	}

	// Print Result
	for i := range *players {
		player := &(*players)[i]
		fmt.Printf("\t Player #%d (%d) : %v \n", player.ID, player.Score, player.Dices)

		(*finalResult)[i] = *player
	}
}

func removeIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}
