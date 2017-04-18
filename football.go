package main

import (
	"fmt"
	"math/rand"
	"sync"
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
	2:  WR,
	3:  WR,
	4:  WR,
	5:  WR,
	6:  OL,
	7:  OL,
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

var offPlayerChan = make(chan Player)
var defPlayerChan = make(chan Player)
var isHiked = false

func main() {
	setupTeams()
	go handleOffPlayerChannel(offPlayerChan)
	go handleDefPlayerChannel(defPlayerChan)

	fmt.Println("Ready for kickoff!")
	fmt.Println("The", stallions.TeamName, "vs The", mustangs.TeamName)

	ticks := 0

	for ticks < GAMETIME {
		snapTheBall()
		time.Sleep(5 * time.Second)
		ticks++
	}

}

func printOffAction(player *Player) {
	if isHiked {
		print("!")
	} else {
		print(".")
	}
}
func printDefAction(player *Player) {
	if isHiked {
		print("$")
	} else {
		print("-")
	}
}

func handleOffPlayerChannel(players <-chan Player) {
	for player := range players {
		duration, _ := time.ParseDuration("200ms")
		time.Sleep(duration)

		switch player.OffPos {
		case QB:
			print("Hike!")
			isHiked = true
		case RB:
			printOffAction(&player)
		case WR:
			printOffAction(&player)
		case OL:
			printOffAction(&player)
		default:
			panic("I don't know what position this is!")

		}
	}

}

func handleDefPlayerChannel(players <-chan Player) {
	for player := range players {
		go doPlayerDefense(&player)
	}
}

func doPlayerDefense(player *Player) {
	m := sync.Mutex{}
	m.Lock()

	for !isHiked {
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	m.Unlock()

	duration, _ := time.ParseDuration(fmt.Sprintf("%vms", rand.Intn(100)))
	time.Sleep(duration)

	switch player.DefPos {

	case DL:
		printDefAction(player)
	case LB:
		printDefAction(player)
	case DB:
		printDefAction(player)
	default:
		panic("I don't know what position this is!")

	}
}

func snapTheBall() {
	go doOffense()
	go doDefense()
}

func doOffense() {
	for _, player := range stallions.Players {
		offPlayerChan <- player
	}
}

func doDefense() {
	for _, player := range mustangs.Players {
		defPlayerChan <- player
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
