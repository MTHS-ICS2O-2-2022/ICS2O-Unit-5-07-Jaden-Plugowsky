// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: May 2023
// This file contains the JS functions for index.html

"use strict"

function buttonOneClicked() {
  // This program takes a user-given number and adds every natural number from 1 to that number together, then giving you the sum
  // Input through Textfield
  const valueOfN = parseInt(document.getElementById("value-of-n").value)
  let answer = null

  // Process
  if (valueOfN >= 1 && valueOfN % 1 == 0) {
    // Only allows positive, whole numbers above 0.
    for (let counter = 1; counter <= valueOfN; counter++) {
      answer += counter
    }
  } else {
    document.getElementById("answer").innerHTML =
      "That input is not a natural number; it is either negative, zero, a decimal, or simply not a number."
  }

  // Output
  if (valueOfN >= 1 && valueOfN % 1 == 0) {
    document.getElementById("answer").innerHTML =
      "The sum of all the numbers from 1 to " + valueOfN + " is: " + answer
  }
}
