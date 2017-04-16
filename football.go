package main

import (
	"fmt"
	"time"
)

const GAMETIME int = 60

// Positions
const QB int = 0
const RB int = 1
const WR int = 2
const OL int = 3
const DL int = 4
const LB int = 5
const DB int = 6

// Map #'s to positions
var offPos = map[int]int{
	0:  QB,
	1:  RB,
	2:  RB,
	3:  WR,
	4:  WR,
	5:  WR,
	6:  WR,
	7:  WR,
	8:  OL,
	9:  OL,
	10: OL,
}

var defPos = map[int]int{
	0:  DL,
	1:  DL,
	2:  DL,
	3:  DL,
	4:  LB,
	5:  LB,
	6:  LB,
	7:  DB,
	8:  DB,
	9:  DB,
	10: DB,
}

var playerChan = make(chan Player)

func main() {
	setupTeams()
	go handlePlayerChannel(playerChan)

	fmt.Println("Ready for kickoff!")
	fmt.Println("The", stallions.TeamName, "vs The", mustangs.TeamName, ". Should be a good one!")

	ticks := 0

	for ticks < GAMETIME {
		snapTheBall()
		time.Sleep(5 * time.Second)
		ticks++
	}

}

func handlePlayerChannel(players <-chan Player) {
	for player := range players {
		fmt.Println(player.Name, "needs something to do")
	}
}

func snapTheBall() {
	go doOffense()
	go doDefense()
}

func doOffense() {
	for _, player := range stallions.Players {
		playerChan <- player
	}
}

func doDefense() {
	for _, player := range mustangs.Players {
		playerChan <- player
	}
}

func setupTeams() {
	teams := []*Team{&stallions, &mustangs}
	for _, team := range teams {
		for i := 0; i < 11; i++ {
			player := createPlayer(i, team)
			team.Players[i] = player
		}
	}
}

func createPlayer(num int, team *Team) Player {
	player := Player{Number: num}
	player.Name = fmt.Sprintf("%v %v", team.TeamName, num)
	player.OffPos = offPos[num]
	player.DefPos = defPos[num]
	return player
}

var stallions = Team{TeamName: "Stallions", Color: "White"}
var mustangs = Team{TeamName: "Mustangs", Color: "Black"}

type Team struct {
	TeamName string
	Color    string
	Players  [11]Player
}

type Player struct {
	Name   string
	Number int
	OffPos int
	DefPos int
}
