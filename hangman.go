package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Main game part
// This is the main function
func main() {
A:
	clearScreen()

	fmt.Println("***************************")
	fmt.Println("*                         *")
	fmt.Println("*      H A N G M A N      *")
	fmt.Println("*                         *")
	fmt.Println("***************************")

	fmt.Println("Welcome to Hangman Game!!!")
	fmt.Println("Press 'c' to continue to game.")
	var val string
	fmt.Scanln(&val)

	switch {
	case val == "c":
		clearScreen()
		playHangman()
	default:
		fmt.Println("Invalid Input")
		time.Sleep(2 * time.Second)
		goto A
	}
}

//Clearing the terminal

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// to replay the game or exit the game
func conti() {
B:
	fmt.Println("Press 'c' to replay game or 'e' to exit the game.")
	var va string
	fmt.Scanln(&va)

	switch {
	case va == "c":
		clearScreen()
		playHangman()
	case va == "e":
		fmt.Println("Thanks for playing the game !! Have a Nice Day !!")
		time.Sleep(2 * time.Second)
		os.Exit(5)
	default:
		fmt.Println("Invalid Input")
		time.Sleep(2 * time.Second)
		goto B
	}
}

// main game logic is Heres
func playHangman() {
	word := []string{
		"head", "wonder", "burger", "switch", "alternative", "passion", "music", "dramatic", "republic", "concatenation", "literacy", "underground", "heaven",
		"certificate", "dinosaurs", "lion", "poker", "scopophobia", "nymphomaniac", "diffcult", "colonel", "enamoured", "ironic", "regardeless", "anasthesia",
	}
	rand.Seed(time.Now().UnixNano())
	var noOfWords int = len(word)
	var noOfLetters = rand.Intn(noOfWords - 1)
	wordIs := word[noOfLetters]
	var number int = len(wordIs)
	str := wordIs

	fmt.Printf("\n")
	var arr = make([]string, number)
	for i := 0; i < number; i++ {
		arr[i] = "_ "
	}
	fmt.Printf("\n\n\n")
	var arr1 = make([]int, number)
	for i := 0; i < number; i++ {
		arr1[i] = int(str[i])
	}
	var count = 0
	var wrong = 0
	var counte = 0
	chances := 8
CALL:
	fmt.Println("********************************")
	fmt.Println(" *   GUESS THE CORRECT WORD   *")
	fmt.Println("********************************")
	fmt.Printf("chances : %d / 8 \n", chances)
	fmt.Println("The length of word is ", number)
	var ab string
	fmt.Println(arr)
	fmt.Println("Input the  guessed Letter:")
	fmt.Scanln(&ab)
	var b int
	var abc int
	var bcd int
	var cde int
	b = int(ab[0])
	for i := 0; i < number; i++ {
		abc = arr1[i]
		bcd = b

		if abc == bcd {
			count += 1
			cde = bcd
			char := string(cde)
			arr[i] = char
			fmt.Println(arr)
			time.Sleep(2 * time.Second)

		}
	}
	for i := 0; i < number; i++ {
		if arr[i] == "_ " {
			counte += 1
		}
	}
	// Winning game Screen
	if counte == 0 {
		fmt.Println("*****************************************************")
		fmt.Println(" *                                                 *")
		fmt.Println("  *  C O N G R A T U L A T I O N  Y O U  W I N !  *")
		fmt.Println(" *                                                 *")
		fmt.Println("*****************************************************")
		conti()
	}
	if count == 0 {
		chances = chances - 1
		fmt.Println("The letter you have entered is incorrect")
		time.Sleep(2 * time.Second)

		wrong += 1
	}
	//Displaying the output
	if wrong == 0 {
		fmt.Println("______________")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		time.Sleep(2 * time.Second)
	} else if wrong == 1 {
		fmt.Println("______________")
		fmt.Println("      |       |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		time.Sleep(2 * time.Second)
	} else if wrong == 2 {
		fmt.Println("______________")
		fmt.Println("      |       |")
		fmt.Println("      O       |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		time.Sleep(2 * time.Second)
	} else if wrong == 3 {
		fmt.Println("______________")
		fmt.Println("      |       |")
		fmt.Println("      O       |")
		fmt.Println("      |       |")
		fmt.Println("      |       |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		time.Sleep(2 * time.Second)
	} else if wrong == 4 {
		fmt.Println("______________")
		fmt.Println("      |       |")
		fmt.Println("      O       |")
		fmt.Println("   <--|       |")
		fmt.Println("      |       |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		time.Sleep(2 * time.Second)
	} else if wrong == 5 {
		fmt.Println("______________")
		fmt.Println("      |       |")
		fmt.Println("      O       |")
		fmt.Println("   <--|-->    |")
		fmt.Println("      |       |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		fmt.Println("              |")
		time.Sleep(2 * time.Second)
	} else if wrong == 6 {
		fmt.Println("______________")
		fmt.Println("      |       |")
		fmt.Println("      O       |")
		fmt.Println("   <--|-->    |")
		fmt.Println("      |       |")
		fmt.Println("       (      |")
		fmt.Println("        |     |")
		fmt.Println("              |")
		fmt.Println("              |")
		time.Sleep(2 * time.Second)
	} else if wrong == 7 {
		fmt.Println("______________")
		fmt.Println("      |       |")
		fmt.Println("      O       |")
		fmt.Println("   <--|-->    |")
		fmt.Println("      |       |")
		fmt.Println("    )   (     |")
		fmt.Println("   |     |    |")
		fmt.Println("              |")
		fmt.Println("              |")
		time.Sleep(2 * time.Second)
	}
	if wrong == 8 {
		fmt.Println("********************")
		fmt.Println(" * Game Over!!!!! *")
		fmt.Println("********************")
		fmt.Println("______________")
		fmt.Println("      |       |")
		fmt.Println("      O       |")
		fmt.Println("   <--|-->    |")
		fmt.Println("      |       |")
		fmt.Println("    )   (     |")
		fmt.Println("   |     |    |")
		fmt.Println("              |")
		fmt.Println("               |")
		fmt.Printf("/n/n")
		fmt.Println("The word is : ", str)
		conti()
	}
	counte = 0
	count = 0
	time.Sleep(2 * time.Second)
	clearScreen()
	goto CALL
}
