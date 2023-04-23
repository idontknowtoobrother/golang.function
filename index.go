package main

import "fmt"

func main() {
	displayPlayerData("John", 100)

	const a, b = 5, 10
	fmt.Println("a + b = ", plus(a, b))

	fmt.Printf("%d + %d = %d", a, b, plus(a, b))

	var c, d = 10, 20
	fmt.Printf("\nc = %d, d = %d", c, d)

	c, d = swap(c, d)
	fmt.Printf("\nafter swap\nc = %d, d = %d\n", c, d)

	c, d = swapWithVariablesExplaination(c, d)
	fmt.Printf("\nafter swap with variables explaination\nc = %d, d = %d\n", c, d)
}

func displayPlayerData(playerName string, playerMoney int) {
	fmt.Println("Player Name: ", playerName, "Player Money: ", playerMoney)
}

func plus(n1 int, n2 int) int {
	return n1 + n2
}

func swap(n1 int, n2 int) (int, int) {
	return n2, n1
}

func swapWithVariablesExplaination(n1, n2 int) (swapToN2, swapToN1 int) {
	swapToN2 = n2
	swapToN1 = n1
	return swapToN2, swapToN1
}
