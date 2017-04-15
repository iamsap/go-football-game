package main

import (
	"fmt"
)

func main() {
	setupTeams()

	fmt.Println("Playing for the ", stallions.TeamName)
	for _, player := range stallions.Players {
		fmt.Println("Introducing", player.Name)
	}

	fmt.Println("Ready for kickoff!")
	fmt.Println("The", stallions.TeamName, "vs The", mustangs.TeamName, ". Should be a good one!")
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
