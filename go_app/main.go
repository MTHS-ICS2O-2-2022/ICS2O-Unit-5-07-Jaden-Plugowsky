// Created by: Jaden Plugowsky
// Created on: May 2023
//
// This program takes a user-given number and adds every natural number from 1 to that number together, then giving you the sum

package main

import (
	"fmt"
	"math"
)

var valueOfN float64
var counter float64
var answer float64
var moduloValueOfN float64

func main() {
	// This function takes a user-given number and adds every natural number from 1 to that number together, then giving you the sum
	// Input
	fmt.Println("This program takes a user-given number and adds every natural number from 1 to that number together.")
	fmt.Println("\nNote: This does not work with zero, negative numbers, or decimals, as they are not considered natural numbers.")
	fmt.Println("\nEnter the value of 'n': ")
	fmt.Scanln(&valueOfN)

	// Process
	moduloValueOfN = math.Mod(valueOfN, 1) // Used to help check if the numbers are decimals or not. If this number is different than zero, the number is a decimal.

	if valueOfN >= 1 && moduloValueOfN == 0 {
		// Only allows positive, whole numbers above 0.
		for counter <= valueOfN {
			answer += counter
			counter++
		}
	} else {
		fmt.Println("\nThat input is not a natural number; it is either negative, zero, a decimal, or simply not a number.")
	}

	// Output
	if valueOfN >= 1 && moduloValueOfN == 0 {
		fmt.Println("\nThe sum of all the numbers from 1 to", valueOfN, "is:", answer)
	}

	fmt.Println("\nDone.")
}
