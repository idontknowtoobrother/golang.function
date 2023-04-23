package main

import "fmt"

func main() {
	displayPlayerData("John", 100)

	const a, b = 5, 10
	fmt.Println("a + b = ", plus(a, b))

	fmt.Printf("%d + %d = %d", a, b, plus(a, b))
}

func displayPlayerData(playerName string, playerMoney int) {
	fmt.Println("Player Name: ", playerName, "Player Money: ", playerMoney)
}

func plus(n1 int, n2 int) int {
	return n1 + n2
}
