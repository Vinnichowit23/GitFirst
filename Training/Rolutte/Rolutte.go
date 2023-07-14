package Rolutte

import (
	"fmt"
)

func R1() {
	var bankroll float32
	fmt.Println("Hello Welcome to Casino Las Vegas: $ Enter the amount you want to Bnakroll $")
	fmt.Scanln(&bankroll)

	for bankroll > 0 {
		fmt.Println("\n--- Roulette Bets ---")
		fmt.Println("1. Bet on Red or Black")
		fmt.Println("2. Bet on high (19 -- 36) or low (1 -- 18)")
		fmt.Println("3. Bet on a specific number")
		fmt.Println("4. Quit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			bankroll = BetonRedORBlack(bankroll)
		case 2:
			bankroll = BetonHighORLow(bankroll)
		case 3:
			bankroll = BetonSpecificNumber(bankroll)
		case 4:
			fmt.Println("Thank you for Playing")
			return
		default:
			fmt.Println("Inavlid Choice!Please Try Again.")
		}

		if bankroll <= 0 {
			fmt.Println("You ran out of Money!!!")
		}

	}
}

func BetonRedORBlack(bankroll float32) float32 {
	var bet, winnings float32
	var color string
	fmt.Println("--- Bet on Red or Black ---")
	fmt.Print("Enter your bet: ")
	fmt.Scanln(&bet)
	fmt.Print("Choose red or black: ")
	fmt.Scanln(&color)

	if color == "red" || color == "black" {
		winnings = bet
		fmt.Printf("Congratulations! You won $%f\n", winnings)
	} else {
		winnings = 0
		fmt.Println("Sorry, you lost.")
	}

	bankroll += winnings
	fmt.Printf("New bankroll: $%f\n", bankroll)
	return bankroll

}

func BetonHighORLow(bankroll float32) float32 {
	var bet, winnings float32
	var choice string

	fmt.Println("--- Bet on High or Low ---")
	fmt.Print("Enter your bet: ")
	fmt.Scanln(&bet)
	fmt.Print("Choose high or low: ")
	fmt.Scanln(&choice)

	if (choice == "low" && bet >= 1 && bet <= 18) || (choice == "high" && bet >= 19 && bet <= 36) {
		winnings = bet
		fmt.Printf("Congratulations! You won $%f\n", winnings)
	} else {
		winnings = 0
		fmt.Println("Sorry, you lost.")
	}

	bankroll += winnings
	fmt.Printf("New bankroll: $%f\n", bankroll)
	return bankroll
}

func BetonSpecificNumber(bankroll float32) float32 {
	var bet, winnings float32
	var number int

	fmt.Println("--- Bet on a Specific Number ---")
	fmt.Print("Enter your bet: ")
	fmt.Scanln(&bet)
	fmt.Print("Enter the number you want to bet on: ")
	fmt.Scanln(&number)

	if number >= 1 && number <= 37 {
		winnings = 35 * bet
		fmt.Printf("Congratulations! You won $%f\n", winnings)
	} else {
		winnings = 0
		fmt.Println("Sorry, you lost.")
	}

	bankroll += winnings
	fmt.Printf("New bankroll: $%f\n", bankroll)
	return bankroll
}
