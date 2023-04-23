package main

import "fmt"

func main() {
	displayPlayerData("John", 100)
}

func displayPlayerData(playerName string, playerMoney int) {
	fmt.Println("Player Name: ", playerName, "Player Money: ", playerMoney)
}
