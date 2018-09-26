package run

import (
	"fmt"

	"github.com/BabichMikhail/holdem/ai"
	"github.com/BabichMikhail/holdem/game"
)

type Result struct {
}

const GOROUTINE_COUNT = 8
const DISTRIBUTION_LIMIT = 100

func playOneGame(aiPlayers []ai.ArtificialIntelligence) {
	gamePlayers := game.Players{}
	for i := 0; i < len(aiPlayers); i++ {
		gamePlayers.NewPlayer(i, fmt.Sprintf("AI #%d", i))
	}
	g := game.NewGame(game.GAME_TYPE_USUAL, 10, 20, gamePlayers, []int{})
	for !g.IsGameOver() {
		g.ApplyAction(aiPlayers[g.Distributions.Last().Round.Current].GetAction(g.Distributions.Last()))
	}
}

func RunGames(count int, aiTypes []int) {
	runningNowCount := 0
	doneChannel := make(chan struct{})
	aiPlayers := []ai.ArtificialIntelligence{}
	for i := 0; i < len(aiTypes); i++ {
		aiPlayers = append(aiPlayers, ai.NewAI(aiTypes[i]))
	}
	runFunc := func(runningCount *int) {
		for *runningCount < count {
			//num := *runningCount
			*runningCount++
			playOneGame(aiPlayers)
		}

		doneChannel <- struct{}{}
	}

	for i := 0; i < GOROUTINE_COUNT; i++ {
		go runFunc(&runningNowCount)
	}
	for i := 0; i < GOROUTINE_COUNT; i++ {
		<-doneChannel
	}
}
