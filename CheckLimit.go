package main

import (
	"math"
)

func CheckLimit(UserId string, Amount float64, uL map[string]float64) bool {
	if Amount < 0 && uL[UserId] >= math.Abs(Amount) {
		uL[UserId] += Amount
		return true
	} else {
		return false
	}
}
