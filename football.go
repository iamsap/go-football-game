package main

import (
	"fmt"
	"time"
)

const GAMETIME int = 60

func main() {
	setupTeams()
	fmt.Println("Ready for kickoff!")
	fmt.Println("The", stallions.TeamName, "vs The", mustangs.TeamName, ". Should be a good one!")

	ticks := 0

	for ticks < GAMETIME {
		snapTheBall()
		time.Sleep(1 * time.Second)
		ticks++
	}

}

func snapTheBall() {
	go doOffense()
	go doDefense()
}

func doOffense() {
	fmt.Println("Go long!")
}

func doDefense() {
	fmt.Println("Stop them!")
}

func setupTeams() {

	// stallions first
	for i := 0; i < 11; i++ {
		player := Player{Number: i}
		player.Name = fmt.Sprintf("%v %v", stallions.TeamName, i)
		stallions.Players[i] = player
	}

	// mustangs next
	for i := 0; i < 11; i++ {
		player := Player{Number: i}
		player.Name = fmt.Sprintf("%v %v", mustangs.TeamName, i)
		mustangs.Players[i] = player
	}
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
}
